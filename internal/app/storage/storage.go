package storage

import m "github.com/belikoooova/url-shortener/internal/app/model"

type Storage interface {
	Save(url m.URL) (*m.URL, error)
	FindByID(id string) (*m.URL, error)
}
