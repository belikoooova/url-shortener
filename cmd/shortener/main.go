package main

import (
	h "github.com/belikoooova/url-shortener/internal/app/handler"
	short "github.com/belikoooova/url-shortener/internal/app/shortener"
	stor "github.com/belikoooova/url-shortener/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	var shortener short.Shortener = short.NewSha256WithBase58Shortener()
	var storage stor.Storage = stor.NewInMemoryStorage()
	createHandler := h.NewCreateHandler(storage, shortener)
	redirectHandler := h.NewRedirectHandler(storage)

	r := chi.NewRouter()
	r.Post("/", createHandler.ServeHTTP)
	r.Get("/{id}", redirectHandler.ServeHTTP)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
