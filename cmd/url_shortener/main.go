package main

import (
	"log"
	"math/rand"
	"time"
	"url_shortener/internal/app/apiserver"
	"url_shortener/internal/store/cache"
)

func main() {
	rand.Seed(time.Now().Unix())
	newCache := cache.New()
	server := apiserver.New(newCache)
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
