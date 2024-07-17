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
	cfg := config.Configure()
	var shortener short.Shortener = short.NewSha256WithBase58Shortener()
	var storage stor.Storage = stor.NewInMemoryStorage()

	createHandler := h.NewCreateHandler(storage, shortener, *cfg)
	redirectHandler := h.NewRedirectHandler(storage)

	startGeneralServer(createHandler, redirectHandler, cfg.ServerAddress)
}

func startAppServer(createHandler *h.CreateHandler, address string) {
	appRouter := chi.NewRouter()
	appRouter.Post("/", createHandler.ServeHTTP)
	err := http.ListenAndServe(address, appRouter)
	if err != nil {
		panic(err)
	}
}

func startRedirectServer(redirectHandler *h.RedirectHandler, address string) {
	redirectRouter := chi.NewRouter()
	redirectRouter.Get("/{id}", redirectHandler.ServeHTTP)
	err := http.ListenAndServe(address, redirectRouter)
	if err != nil {
		panic(err)
	}
}

func startGeneralServer(createHandler *h.CreateHandler, redirectHandler *h.RedirectHandler, address string) {
	router := chi.NewRouter()
	router.Post("/", createHandler.ServeHTTP)
	router.Get("/{id}", redirectHandler.ServeHTTP)
	err := http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
}
