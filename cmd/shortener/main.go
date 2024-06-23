package main

import (
	h "github.com/belikoooova/url-shortener/internal/app/handler"
	short "github.com/belikoooova/url-shortener/internal/app/shortener"
	stor "github.com/belikoooova/url-shortener/internal/app/storage"
	"net/http"
)

func main() {
	var shortener short.Shortener = short.NewSha256WithBase58Shortener()
	var storage stor.Storage = stor.NewInMemoryStorage()
	createHandler := h.NewCreateHandler(storage, shortener)
	redirectHandler := h.NewRedirectHandler(storage)
	mux := http.NewServeMux()
	mux.Handle("/", createHandler)
	mux.Handle("/{id}", redirectHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
