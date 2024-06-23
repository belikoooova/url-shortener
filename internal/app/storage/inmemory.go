package storage

import (
	"errors"
	m "github.com/belikoooova/url-shortener/internal/app/model"
)

type InMemoryStorage struct {
	urlForId map[string]m.Url
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{urlForId: make(map[string]m.Url)}
}

func (s *InMemoryStorage) Save(url m.Url) (*m.Url, error) {
	foundUrl, ok := s.urlForId[url.Id]
	if ok && foundUrl.OriginalUrl != url.OriginalUrl {
		return nil, errors.New("short url already exists")
	}

	s.urlForId[url.Id] = url
	return &url, nil
}

func (s *InMemoryStorage) FindById(id string) (*m.Url, error) {
	url, ok := s.urlForId[id]
	if ok {
		return &url, nil
	}

	return nil, errors.New("url does not exist")
}
