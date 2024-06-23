package storage

import m "github.com/belikoooova/url-shortener/internal/app/model"

type Storage interface {
	Save(url m.Url) (*m.Url, error)
	FindById(id string) (*m.Url, error)
}
