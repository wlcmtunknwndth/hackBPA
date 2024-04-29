package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/wlcmtunknwndth/hackBPA/internal/auth"
	"github.com/wlcmtunknwndth/hackBPA/internal/broker/nats"
	"github.com/wlcmtunknwndth/hackBPA/internal/config"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage/postgres"
	"log/slog"
	"net/http"
)

const scope = "main"

func main() {
	cfg := config.MustLoad()
	slog.Info("Config: ", slog.Attr{Key: "Config", Value: slog.AnyValue(*cfg)})

	ns, err := nats.New(cfg)
	if err != nil {
		slog.Error("couldn't run nats:", slogResponse.SlogErr(err))
	}
	defer func(ns *nats.Nats) {
		err := ns.Close()
		if err != nil {
			slog.Error("couldn't close NATS:", slogResponse.SlogErr(err))
		}
	}(ns)
	slog.Info("successfully initialized NATS")

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

	db, err := postgres.New(&cfg.DB)
	if err != nil {
		slog.Error("couldn't connect to storage", slogResponse.SlogOp(scope), slogResponse.SlogErr(err))
		return
	}
	defer func(db *postgres.Storage) {
		err := db.Close()
		if err != nil {
			slog.Error("couldn't close connection to storage", slogResponse.SlogOp(scope), slogResponse.SlogErr(err))
			return
		}
	}(db)
	slog.Info("successfully initialized storage")

	authService := auth.Auth{Db: db}

	router.Post("/register", authService.Register)
	router.Post("/login", authService.LogIn)
	router.Post("/logout", authService.LogOut)
	router.Delete("/delete_user", authService.DeleteUser)

	if err = srv.ListenAndServe(); err != nil {
		slog.Error("failed to run server: ", slogResponse.SlogErr(err))
	}
	slog.Info("server closed")
}
