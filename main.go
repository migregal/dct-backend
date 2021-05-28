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

	net.StartServer(cfg.Daemon.Port.String(), logger)
}
