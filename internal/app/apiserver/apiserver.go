package apiserver

import (
	"fmt"
	"net/http"
	"url_shortener/internal/app/generator"
	"url_shortener/internal/store"
	"url_shortener/internal/version"

	"github.com/gorilla/mux"
)

type Server struct {
	generator *generator.Generator
	router    *mux.Router
	store     store.Store
}

func New(inStore store.Store) *Server {
	server := &Server{
		generator: generator.New(inStore),
		router:    mux.NewRouter(),
		store:     inStore,
	}
	server.configureRouter()
	return server
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/not_found", s.handlerNotFound).Methods(http.MethodGet)
	s.router.HandleFunc("/version", s.handlerVersion).Methods(http.MethodGet)
	s.router.HandleFunc("/{shortLink}", s.handlerGet).Methods(http.MethodGet)
	s.router.HandleFunc("/", s.handlerPost).Methods(http.MethodPost)
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8000", s.router)
}

func (s *Server) handlerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortLink := vars["shortLink"]
	originalLink, err := s.store.Get(shortLink)
	if err != nil {
		w.Header().Add("Location", "not_found")
		w.WriteHeader(http.StatusMovedPermanently)
		return
	}
	w.Header().Add("Location", originalLink)
	w.WriteHeader(http.StatusMovedPermanently)
}

func (s *Server) handlerNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go go")
	w.WriteHeader(http.StatusNotFound)
}

func (s *Server) handlerPost(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handlerVersion(w http.ResponseWriter, r *http.Request) {
	programName, gitTag, gitCommit, gitBranch := version.Version()
	fmt.Fprintf(w, "pragramName=%s\tgitTag=%s\tgitCommit=%s\tgitBranch=%s", programName, gitTag, gitCommit, gitBranch)
}
