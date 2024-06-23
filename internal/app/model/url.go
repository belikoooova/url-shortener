package model

const BaseUrl = "http://localhost:8080/"

type Url struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"original_url"`
	ShortUrl    string `json:"short_url"`
}
