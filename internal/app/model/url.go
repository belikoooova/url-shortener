package model

const BaseURL = "http://localhost:8080/"

type URL struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}
