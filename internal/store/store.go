package store

type Store interface {
	Get(shortLink string) (string, error)
	Put(shortLink, link string) error
	IsShortLinkExists(shortLink string) (bool, error)
}
