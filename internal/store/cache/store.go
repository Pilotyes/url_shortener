package cache

import "errors"

// Cache is a local storage struct
type Cache struct {
	store map[string]string
}

// New returns pointer to a new Cache
func New() *Cache {
	return &Cache{
		store: map[string]string{},
	}
}

// Get returns long link and error, if short link not found
func (c *Cache) Get(shortLink string) (string, error) {
	link, ok := c.store[shortLink]
	if !ok {
		return "", errors.New("not found")
	}
	return link, nil
}

// Put returns error
func (c *Cache) Put(shortLink, link string) error {
	if _, ok := c.store[shortLink]; ok {
		return errors.New("already exist")
	}
	c.store[shortLink] = link
	return nil
}

// IsShortLinkExists checks short link in cache
func (c *Cache) IsShortLinkExists(shortLink string) (bool, error) {
	_, ok := c.store[shortLink]
	return ok, nil
}
