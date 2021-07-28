package main

import (
	"finnflare.com/dct_backend/config"
	"finnflare.com/dct_backend/net"
	"fmt"
	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"os"
)

const usage = "Usage: dct_backend install | remove | start | stop | restart | status"
const serviceName = "FFDCT service, version 1.0.0"
const serviceDescription = "Finn Flare DCT service"

type Conf struct {
	conf   *config.Config
	logger *logrus.Logger
}

var conf = Conf{}

type program struct {
	server net.Server
	logger *logrus.Logger
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	if conf.conf == nil {
		cfg, logger, err := setUp()

		if err != nil {
			panic(err)
		}

		if cfg == nil || logger == nil {
			panic("")
		}

		conf.conf = cfg
		conf.logger = logger
	}

	p.server = net.NewServer(*conf.conf, conf.logger)
	p.logger = conf.logger
	p.server.Start(
		conf.conf.Daemon,
		conf.logger,
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
			return "installed", s.Install()
		case "remove":
			return "removed", s.Uninstall()
		case "start":
			return "started", s.Start()
		case "stop":
			return "stopped", s.Stop()
		case "restart":
			return "restarted", s.Restart()
		case "status":
			status, err := s.Status()
			if err != nil {
				return "error", err
			}
			switch status {
			case service.StatusRunning:
				return "Status: running", nil
			case service.StatusStopped:
				return "Status: stopped", nil
			default:
				return "Status: unknown", nil
			}
		case "usage":
			return usage, nil
		}
	}

	if service.Interactive() {
		cfg, logger, err := setUp()

		if err != nil {
			panic(err)
		}

		if cfg == nil || logger == nil {
			return "", nil
		}

		conf.conf = cfg
		conf.logger = logger
	}

	err := s.Run()
	if err != nil {
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
		fmt.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	if status != "" {
		fmt.Println(status)
	}
}
