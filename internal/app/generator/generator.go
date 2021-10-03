package generator

import (
	"math/rand"
	"url_shortener/internal/store"
)

const (
	alphabet = "abcdefghABCDEFGH"
	length   = 8
)

type Generator struct {
	store store.Store
}

func New(inStore store.Store) *Generator {
	return &Generator{
		store: inStore,
	}
}

func (g *Generator) Generate() string {
	shortLink := ""
	for {
		shortLink = generate()
		if ok, _ := g.store.IsShortLinkExists(shortLink); !ok {
			break
		}
	}
	return shortLink
}

func generate() string {
	var shortLink string
	for i := 0; i < length; i++ {
		shortLink += string(alphabet[rand.Intn(len(alphabet))]) //nolint
	}
	return shortLink
}
