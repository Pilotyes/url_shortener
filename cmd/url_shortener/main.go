package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
	"url_shortener/internal/app/apiserver"
	"url_shortener/internal/config"
	"url_shortener/internal/store/cache"
)

func main() {
	rand.Seed(time.Now().Unix())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	newCache := cache.New()
	newConfig := config.NewConfig()
	server := apiserver.New(newCache, newConfig)
	if err := server.Start(ctx); err != nil {
		log.Fatalln(err)
	}
}
