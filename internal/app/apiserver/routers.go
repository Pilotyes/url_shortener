package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (s *Server) newRouter() *mux.Router {
	var routes = Routes{
		Route{
			"Index",
			http.MethodGet,
			"/",
			s.index,
		},

		Route{
			"CreateShortLink",
			http.MethodPost,
			"/",
			s.createShortLink,
		},

		Route{
			"Version",
			http.MethodGet,
			"/version",
			s.version,
		},

		Route{
			"Redirect",
			http.MethodGet,
			"/{shortLink}",
			s.redirect,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	router.Use(Logger)
	for _, route := range routes {
		var handler http.Handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
