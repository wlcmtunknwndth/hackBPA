package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/wlcmtunknwndth/hackBPA/internal/auth"
	"github.com/wlcmtunknwndth/hackBPA/internal/broker/nats"
	"github.com/wlcmtunknwndth/hackBPA/internal/config"
	"github.com/wlcmtunknwndth/hackBPA/internal/handlers/event"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage/postgres"
	"log/slog"
	"net/http"
)

const scope = "main"

func main() {
	cfg := config.MustLoad()
	slog.Info("Config: ", slog.Attr{Key: "Config", Value: slog.AnyValue(*cfg)})

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	//router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Logger)

	//m := &autocert.Manager{
	//	Cache:      autocert.DirCache("golang-autocert"),
	//	Prompt:     autocert.AcceptTOS,
	//	HostPolicy: autocert.HostWhitelist("example.org", "www.example.org"),
	//}

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		//TLSConfig:    m.TLSConfig(),
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

	ns, err := nats.New(&cfg.Nats, db)
	if err != nil {
		slog.Error("couldn't run nats:", slogResponse.SlogErr(err))
		return
	}
	defer ns.Close()
	slog.Info("nats created")

	del, err := ns.EventDeleter(context.Background())
	if err != nil {
		slog.Error("couldn't run deleter", slogResponse.SlogErr(err))
		return
	}
	defer del.Unsubscribe()
	save, err := ns.EventSaver(context.Background())
	if err != nil {
		slog.Error("couldn't run saver", slogResponse.SlogErr(err))
		return
	}
	defer save.Unsubscribe()

	send, err := ns.EventSender(context.Background())
	if err != nil {
		slog.Error("couldn't run sender", slogResponse.SlogErr(err))
		return
	}
	defer send.Unsubscribe()

	slog.Info("successfully initialized NATS")

	authService := auth.Auth{Db: db}

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
	})
	router.Post("/register", authService.Register)
	router.Post("/login", authService.LogIn)
	router.Post("/logout", authService.LogOut)
	router.Delete("/delete_user", authService.DeleteUser)

	eventService := event.EventsHandler{Broker: ns}

	router.Post("/create_event", eventService.CreateEvent)
	router.Get("/event", eventService.GetEvent)
	router.Delete("/delete", eventService.DeleteEvent)

	if err = srv.ListenAndServe(); err != nil {
		slog.Error("failed to run server: ", slogResponse.SlogErr(err))
	}
	slog.Info("server closed")
}
