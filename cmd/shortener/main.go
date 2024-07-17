package main

import (
	"github.com/belikoooova/url-shortener/cmd/shortener/config"
	h "github.com/belikoooova/url-shortener/internal/app/handler"
	"github.com/belikoooova/url-shortener/internal/app/logger"
	short "github.com/belikoooova/url-shortener/internal/app/shortener"
	stor "github.com/belikoooova/url-shortener/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	cfg := config.Configure()
	var shortener short.Shortener = short.NewSha256WithBase58Shortener()
	var storage stor.Storage = stor.NewInMemoryStorage()

	createHandler := h.NewCreateHandler(storage, shortener, *cfg)
	redirectHandler := h.NewRedirectHandler(storage)

	runLogger(cfg.LogLevel)
	runServer(createHandler, redirectHandler, cfg.ServerAddress)
}

func runLogger(logLevel string) {
	err := logger.Initialize(logLevel)
	if err != nil {
		panic(err)
	}
}

func runServer(createHandler *h.CreateHandler, redirectHandler *h.RedirectHandler, address string) {
	router := chi.NewRouter()
	router.Post("/", logger.WithLogging(createHandler).ServeHTTP)
	router.Get("/{id}", logger.WithLogging(redirectHandler).ServeHTTP)
	err := http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
}
