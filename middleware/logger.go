package middleware

import (
	"log"
	"net/http"
)

func Logger(handler http.Handler) http.Handler {
	logger := func(w http.ResponseWriter, r *http.Request) {
		log.Print("START - ", r.RequestURI, "  -  ", r.Method)
		handler.ServeHTTP(w, r)
		log.Print("END - ", r.RequestURI, "  -  ", r.Method)
	}

	return http.HandlerFunc(logger)
}
