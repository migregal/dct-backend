package net

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
)

//-------------------------------

var filesHandler = fasthttp.FSHandler("./static", 0)

func requestHandler(ctx *fasthttp.RequestCtx) {
	filesHandler(ctx)
}

func StartServer(addr string, logger *logrus.Logger) {
	go func() {
		if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
			logger.Error(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
}
