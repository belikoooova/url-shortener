package handler

import (
	m "github.com/belikoooova/url-shortener/internal/app/model"
	short "github.com/belikoooova/url-shortener/internal/app/shortener"
	stor "github.com/belikoooova/url-shortener/internal/app/storage"
	"io"
	"net/http"
	"strings"
)

type CreateHandler struct {
	storage   stor.Storage
	shortener short.Shortener
}

func NewCreateHandler(storage stor.Storage, shortener short.Shortener) *CreateHandler {
	return &CreateHandler{storage: storage, shortener: shortener}
}

func (h CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if !strings.Contains(r.Header.Get("Content-Type"), "text/plain") {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	originalURLBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	originalURL := string(originalURLBytes)

	hash, err := h.shortener.Shorten(originalURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	URL := m.URL{ID: hash, OriginalURL: originalURL, ShortURL: m.BaseURL + hash}

	var savedURL *m.URL
	savedURL, err = h.storage.Save(URL)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(savedURL.ShortURL))
	if err != nil {
		return
	}
}
