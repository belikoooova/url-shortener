package main

import (
	"github.com/belikoooova/url-shortener/cmd/shortener/config"
	h "github.com/belikoooova/url-shortener/internal/app/handler"
	short "github.com/belikoooova/url-shortener/internal/app/shortener"
	stor "github.com/belikoooova/url-shortener/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	config.ParseFlags()
	cfg := config.Config{
		AppRunServerAddress:   config.AppRunAddress,
		RedirectServerAddress: config.RedirectAddress,
	}
	var shortener short.Shortener = short.NewSha256WithBase58Shortener()
	var storage stor.Storage = stor.NewInMemoryStorage()

	go startAppServer(&storage, &shortener, &cfg)
	go startRedirectServer(&storage, &cfg)

	select {}
}

func startAppServer(storage *stor.Storage, shortener *short.Shortener, cfg *config.Config) {
	createHandler := h.NewCreateHandler(*storage, *shortener, *cfg)
	appRouter := chi.NewRouter()
	appRouter.Post("/", createHandler.ServeHTTP)
	err := http.ListenAndServe(cfg.AppRunServerAddress, appRouter)
	if err != nil {
		panic(err)
	}
}

func startRedirectServer(storage *stor.Storage, cfg *config.Config) {
	redirectHandler := h.NewRedirectHandler(*storage)
	redirectRouter := chi.NewRouter()
	redirectRouter.Get("/{id}", redirectHandler.ServeHTTP)
	err := http.ListenAndServe(cfg.RedirectServerAddress, redirectRouter)
	if err != nil {
		panic(err)
	}
}
