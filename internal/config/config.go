package config

import (
	"github.com/didslm/env"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"log/slog"
	"os"
)

type Config struct {
	DB     Database `obj:"db"`
	Server Server   `obj:"server"`
}

type Database struct {
	DbUser  string `env:"db_user"`
	DbPass  string `env:"db_pass"`
	DbName  string `env:"db_name"`
	SslMode string `env:"ssl_mode"`
}

type Server struct {
	Timeout     int    `env:"timeout"`
	IdleTimeout int    `env:"idle_timeout"`
	Address     string `env:"address"`
}

func MustLoad() *Config {
	const op = "config.MustLoad"

	var config Config

	if err := env.PopulateWithEnv(&config); err != nil {
		slog.Error("couldn't load config file: ", slogResponse.SlogErr(err))
		os.Exit(404)
	}

	slog.Info("loaded conf")
}
