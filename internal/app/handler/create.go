package handler

import (
	m "github.com/belikoooova/url-shortener/internal/app/model"
	short "github.com/belikoooova/url-shortener/internal/app/shortener"
	stor "github.com/belikoooova/url-shortener/internal/app/storage"
	"io"
	"net/http"
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

	if r.Header.Get("Content-Type") != "text/plain" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	originalUrlBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	originalUrl := string(originalUrlBytes)

	hash, err := h.shortener.Shorten(originalUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	url := m.Url{Id: hash, OriginalUrl: originalUrl, ShortUrl: m.BaseUrl + hash}

	var savedUrl *m.Url
	savedUrl, err = h.storage.Save(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(savedUrl.ShortUrl))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
