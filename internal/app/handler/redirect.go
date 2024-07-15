package handler

import (
	stor "github.com/belikoooova/url-shortener/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const locationHeader = "Location"

type RedirectHandler struct {
	stor.Storage
}

func NewRedirectHandler(storage stor.Storage) *RedirectHandler {
	return &RedirectHandler{storage}
}

func (h *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url, err := h.Storage.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add(locationHeader, url.OriginalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
