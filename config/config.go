package config

import (
	"encoding/json"
	"errors"
	"finnflare.com/favourites/validator"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"path/filepath"
)

type Daemon struct {
	LogPath        string      `json:"log_path" yaml:"log_path" env:"LOG_PATH" env-default:"." env-description:"Log files storing path"`
	Port           json.Number `json:"port,int" yaml:"port" env:"PORT" env-description:"Daemon port"`
	Token          string      `json:"token" yaml:"token" env:"TOKEN" env-description:"Daemon access token"`
	WorkersPulSize int         `json:"workers_pul_size" yaml:"workers_pul_size" env:"WORKERS_PUL_SIZE" env-default:"16" env-description:"Max number of workers"`
}

type Config struct {
	Daemon Daemon `json:"daemon" yml:"daemon"`
}

func (cfg *Config) loadConfig(fileName string) error {
	return cleanenv.ReadConfig(fileName, cfg)
}

func (cfg *Config) loadEnvConfig() error {
	return cleanenv.ReadEnv(cfg)
}

func (d Daemon) checkDaemonConfig() bool {
	if path, err := filepath.Abs(d.LogPath); err == nil {
		d.LogPath = path
	} else {
		fmt.Println(d.LogPath)
		return false
	}

	if !validator.IsPort(d.Port.String()) {
		return false
	}

	if d.WorkersPulSize <= 0 {
		return false
	}

	return true
}

func (cfg *Config) GetDescription() (string, error) {
	return cleanenv.GetDescription(cfg, nil)
}

func (cfg *Config) GetConfig(fileName string, env bool) error {
	if err := cfg.loadConfig(fileName); err != nil {
		if !env {
			return err
		}

		if err = cfg.loadEnvConfig(); err != nil {
			return err
		}
	}

	if !cfg.Daemon.checkDaemonConfig() {
		return errors.New("incorrect daemon configuration")
	}

	if !cfg.Daemon.checkDaemonConfig() {
		return errors.New("incorrect database configuration")
	}

	return nil
}
