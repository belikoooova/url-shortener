package shortener

type Shortener interface {
	Shorten(url string) (string, error)
}
