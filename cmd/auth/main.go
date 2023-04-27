package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"go-one-auth/config"
	httphandler "go-one-auth/internal/handlers/http"
	"go-one-auth/internal/storage"
	"go-one-auth/internal/usecase"
	"go-one-auth/pkg/logger"
	"log"
	"net/http"
)

func main() {
	cfg := config.New()
	repo := storage.New(cfg.DBConfig)
	logic := usecase.New(repo)
	loggerInstance := httplog.NewLogger("auth", httplog.Options{
		Concise: true,
	})

	if cfg.HTTP != "" {
		log.Println("Starting server on", cfg.HTTP)
		router := chi.NewRouter()
		router.Use(httplog.RequestLogger(loggerInstance))

		handler := httphandler.New(logic, logger.New(loggerInstance))
		router.Group(handler.PublicRoutes)

		if err := http.ListenAndServe(cfg.HTTP, router); err != nil {
			log.Fatal(err.Error())
		}
	}
}
