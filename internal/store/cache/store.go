package cache

import "errors"

type Cache struct {
	store map[string]string
}

func New() *Cache {
	return &Cache{
		store: map[string]string{
			"aaa": "https://yandex.ru",
			"bbb": "https://google.com",
		},
	}
}

func (c *Cache) Get(shortLink string) (string, error) {
	link, ok := c.store[shortLink]
	if !ok {
		return "", errors.New("not found")
	}
	return link, nil
}

func (c *Cache) Put(shortLink, link string) error {
	if _, ok := c.store[shortLink]; ok {
		return errors.New("already exist")
	}
	c.store[shortLink] = link
	return nil
}

func (c *Cache) IsShortLinkExists(shortLink string) (bool, error) {
	_, ok := c.store[shortLink]
	return ok, nil
}
