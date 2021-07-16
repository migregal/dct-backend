package main

import (
	"finnflare.com/dct_backend/net"
	"fmt"
	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"os"
)

const usage = "Usage: dct_backend install | remove | start | stop | status"
const serviceName = "FFDCT service, version 1.0.0"
const serviceDescription = "Finn Flare DCT service"

type program struct {
	server net.Server
	logger *logrus.Logger
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	cfg, logger, err := setUp()

	if err != nil {
		panic(err)
	}

	if cfg == nil || logger == nil {
		return
	}

	p.server = net.NewServer(cfg.Daemon, logger)
	p.logger = logger
	p.server.Start(
		cfg.Daemon,
		logger,
	)
}

func (p *program) Stop(s service.Service) error {
	p.server.Stop(p.logger)
	return nil
}

func Manage(s service.Service) (string, error) {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return "installation", s.Install()
		case "remove":
			return "uninstallation", s.Uninstall()
		case "start":
			return "start", s.Start()
		case "stop":
			return "stop", s.Stop()
		case "restart":
			return "restart", s.Restart()
		case "status":
			status, err := s.Status()
			return string(status), err
		case "usage":
			return usage, nil
		default:
			if err := s.Run(); err != nil {
				panic(err)
			}
			return "finished", nil
		}
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
	return "finished", nil
}

func main() {
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}

	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		panic(err)
	}

	status, err := Manage(s)
	if err != nil {
		fmt.Println("\n", status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
