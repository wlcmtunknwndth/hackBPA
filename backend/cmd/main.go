package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/wlcmtunknwndth/hackBPA/internal/broker/nats"
	"github.com/wlcmtunknwndth/hackBPA/internal/config"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"log/slog"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	slog.Info("Config: ", slog.Attr{Key: "Config", Value: slog.AnyValue(*cfg)})

	ns, err := nats.New(cfg)
	if err != nil {
		slog.Error("couldn't run nats:", slogResponse.SlogErr(err))
	}
	defer ns.Close()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Logger)

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	if err = srv.ListenAndServe(); err != nil {
		slog.Error("failed to run server: ", slogResponse.SlogErr(err))
	}
	slog.Info("server closed")
}
