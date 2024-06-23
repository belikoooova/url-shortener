package storage

import (
	"errors"
	m "github.com/belikoooova/url-shortener/internal/app/model"
)

type InMemoryStorage struct {
	urlForID map[string]m.URL
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{urlForID: make(map[string]m.URL)}
}

func (s *InMemoryStorage) Save(url m.URL) (*m.URL, error) {
	foundURL, ok := s.urlForID[url.ID]
	if ok && foundURL.OriginalURL != url.OriginalURL {
		return nil, errors.New("short url already exists")
	}

	s.urlForID[url.ID] = url
	return &url, nil
}

func (s *InMemoryStorage) FindByID(id string) (*m.URL, error) {
	url, ok := s.urlForID[id]
	if ok {
		return &url, nil
	}

	return nil, errors.New("url does not exist")
}
