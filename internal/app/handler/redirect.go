package handler

import (
	stor "github.com/belikoooova/url-shortener/internal/app/storage"
	"net/http"
)

type RedirectHandler struct {
	stor.Storage
}

func NewRedirectHandler(storage stor.Storage) *RedirectHandler {
	return &RedirectHandler{storage}
}

func (h *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[1:len(r.URL.Path)]
	url, err := h.Storage.FindByID(id)
	w.Header().Add("Location", url.OriginalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
