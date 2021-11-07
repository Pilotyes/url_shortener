package apiserver

import (
	"context"
	"log"
	"net/http"
	"time"
	"url_shortener/internal/app/generator"
	"url_shortener/internal/config"
	"url_shortener/internal/store"

	"github.com/gorilla/mux"
)

const (
	shutdownTimeout = 5
)

type Server struct {
	config     *config.MainConfig
	generator  *generator.Generator
	httpServer *http.Server
	router     *mux.Router
	store      store.Store
}

func New(inStore store.Store, inConfig *config.MainConfig) *Server {
	newGenerator := generator.New(inStore)

	server := &Server{
		config:    inConfig,
		generator: newGenerator,
		store:     inStore,
	}

	router := server.newRouter()
	server.router = router

	httpServer := &http.Server{
		Addr:    inConfig.BindAddr,
		Handler: router,
	}
	server.httpServer = httpServer

	return server
}

func (s *Server) Start(ctx context.Context) error {
	var err error
	go func() {
		if err = s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %+s\n", err)
		}
	}()

	log.Printf("URL shortener server started")

	<-ctx.Done()

	log.Printf("URL shortener server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
	defer cancel()

	err = s.httpServer.Shutdown(ctxShutDown)
	if err != nil {
		return err
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}
