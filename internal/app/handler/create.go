package handler

import (
	"github.com/belikoooova/url-shortener/cmd/shortener/config"
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
	cfg       config.Config
}

func NewCreateHandler(storage stor.Storage, shortener short.Shortener, cfg config.Config) *CreateHandler {
	return &CreateHandler{storage: storage, shortener: shortener, cfg: cfg}
}

func (h CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	URL := m.URL{ID: hash, OriginalURL: originalURL, ShortURL: h.cfg.BaseUrl + "/" + hash}

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
