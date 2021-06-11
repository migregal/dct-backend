package main

import (
	"finnflare.com/dct_backend/net"
)

func main() {
	cfg, logger, err := setUp()

	if err != nil {
		panic(err)
	}

	if cfg == nil || logger == nil {
		return
	}

	net.StartServer(
		cfg.Daemon,
		//cfg.Daemon.Port.String(),
		//"https://api.finnflare.com:48054",
		//cfg.Daemon.AccessToken,
		//cfg.Daemon.RedirectToken,
		//cfg.Daemon.WorkersPullSize,
		logger,
	)
}
