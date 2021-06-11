package net

import (
	"encoding/base64"
	"finnflare.com/dct_backend/config"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"time"
)

const (
	login = "TSDExchange.login"
	tasks = "TSDExchange.get.tasklist"
)

var (
	POST = []byte(fasthttp.MethodPost)
)

type Handler struct {
	Log           *logrus.Logger
	Url           string
	AccessToken   string
	RedirectToken string
	Auth          string
	JobQueue      chan struct{}
}

func NewHandler(maxWorkers int, url string, accessToken string, redirectToken string) *Handler {
	pool := make(chan struct{}, maxWorkers)
	return &Handler{
		Url:           url,
		AccessToken:   accessToken,
		RedirectToken: redirectToken,
		Auth:          "Basic " + basicAuth("WMS", ""),
		JobQueue:      pool,
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (h *Handler) ServeHTTP(ctx *fasthttp.RequestCtx) {
	h.JobQueue <- struct{}{}
	defer func() { <-h.JobQueue }()

	if ctx.IsGet() {
		requestHandler(ctx)
		return
	}

	if !ctx.IsPost() {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	go h.Log.Info(ctx.Method, string(ctx.Request.Body()))

	ctx.SetContentType("application/json")
	ctx.SetBody(h.processRequest(ctx.PostBody()))
	ctx.SetStatusCode(fasthttp.StatusOK)
}

//-------------------------------

var filesHandler fasthttp.RequestHandler = nil

func requestHandler(ctx *fasthttp.RequestCtx) {
	filesHandler(ctx)
}

func StartServer(cfg config.Daemon, logger *logrus.Logger) {
	if filesHandler == nil {
		filesHandler = fasthttp.FSHandler("./static", 0)
	}

	rootHandler := NewHandler(
		cfg.WorkersPullSize,
		cfg.RedirectUrl,
		cfg.AccessToken,
		cfg.RedirectToken,
	)
	rootHandler.Log = logger

	server := fasthttp.Server{
		Logger:      logger,
		Handler:     rootHandler.ServeHTTP,
		ReadTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(":" + cfg.Port.String()); err != nil {
			logger.Error(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	if err := server.Shutdown(); err != nil {
		logger.Error(err)
	}
}
