package net

import (
	"encoding/base64"
	"finnflare.com/dct_backend/config"
	"github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
	"time"
)

const (
	login                 = "TSDExchange.login"
	operations            = "TSDExchange.get.operations"
	tasks                 = "TSDExchange.get.tasklist"
	cases                 = "TSDExchange.get.caselist"
	addCase               = "TSDExchange.set.newcase"
	caseNotFound          = "TSDExchange.set.casenotfound"
	taskStopped           = "TSDExchange.set.taskstoped"
	skuCaseList           = "TSDExchange.get.skucaselist"
	skuGuid               = "TSDExchange.get.skuguid"
	skuCaseMove           = "TSDExchange.set.skucasemove"
	skuCaseNotfound       = "TSDExchange.set.skucasenotfound"
	locForCase            = "TSDExchange.get.locforcaseid"
	skuCaseDiff           = "TSDExchange.get.skulistfromcasediff"
	locForLoc             = "TSDExchange.get.locforlocid"
	skuDiffMove           = "TSDExchange.set.skudiffmove"
	cancelDefragmentation = "TSDExchange.set.canсeldiffforcaseid"
	endDefragmentation    = "TSDExchange.set.enddiffforcaseid"
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

func NewHandler(maxWorkers int, url string, accessToken string, redirectToken string, login string, pwd string) *Handler {
	pool := make(chan struct{}, maxWorkers)
	return &Handler{
		Url:           url,
		AccessToken:   accessToken,
		RedirectToken: redirectToken,
		Auth:          "Basic " + basicAuth(login, pwd),
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

type Server struct {
	srv fasthttp.Server
}

func NewServer(cfg config.Config, logger *logrus.Logger) Server {
	if filesHandler == nil {
		filesHandler = fasthttp.FSHandler(cfg.CurPass+string(os.PathSeparator)+"static", 0)
	}

	rootHandler := NewHandler(
		cfg.Daemon.WorkersPullSize,
		cfg.Daemon.RedirectUrl,
		cfg.Daemon.AccessToken,
		cfg.Daemon.RedirectToken,
		cfg.Auth.Login,
		cfg.Auth.Password,
	)
	rootHandler.Log = logger

	withCors := fasthttpcors.NewCorsHandler(fasthttpcors.Options{
		AllowedMethods:   []string{"GET", "POST"}, // only allow get or post to resource
		AllowCredentials: false,                   // resource doesn't support credentials
		AllowMaxAge:      int(10 * time.Second),   // cache the preflight result
	})
	return Server{
		fasthttp.Server{
			Logger:      logger,
			Handler:     withCors.CorsMiddleware(rootHandler.ServeHTTP),
			ReadTimeout: 5 * time.Second,
		},
	}
}

func (srv *Server) Start(cfg config.Daemon, logger *logrus.Logger) {
	go func() {
		if err := srv.srv.ListenAndServe(":" + cfg.Port.String()); err != nil {
			logger.Error(err)
		}
	}()
}

func (srv *Server) Stop(logger *logrus.Logger) {
	if err := srv.srv.Shutdown(); err != nil {
		logger.Error(err)
	}
}
