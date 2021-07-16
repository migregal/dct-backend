package main

import (
	"finnflare.com/dct_backend/config"
	"finnflare.com/dct_backend/logger"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func setUp() (*config.Config, *logrus.Logger, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	help, verbose, confFile := parseFlags()

	if *help {
		displayHelp()
		return nil, nil, nil
	}

	if *confFile == "" {
		*confFile = "config.json"
	}

	var cfg config.Config
	if err := cfg.GetConfig(exPath+string(os.PathSeparator)+*confFile, true); err != nil {
		return nil, nil, err
	}
	cfg.CurPass = exPath
	cfg.Daemon.LogPath = cfg.CurPass + string(os.PathSeparator) + cfg.Daemon.LogPath

	log, err := logger.ConfigureLogger(cfg.Daemon.LogPath)

	if err != nil {
		return nil, nil, err
	}

	if *verbose {
		fmt.Printf("Daemon: %+v\n", cfg.Daemon)
	}

	return &cfg, log, nil
}

func parseFlags() (*bool, *bool, *string) {
	help := flag.Bool("h", false, "Display config description")
	verbose := flag.Bool("v", false, "Display config result")
	confFile := flag.String("config", "", "Config file name")
	flag.Parse()

	return help, verbose, confFile
}

func displayHelp() {
	var cfg config.Config
	str, err := cfg.GetDescription()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(str)
}
