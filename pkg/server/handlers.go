package server

import (
	"fmt"
	"net/http"
)

func (s *Server) handleAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do something with it.
	}
}

func (s *Server) handleGreeting(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do something with it.
		fmt.Fprintf(w, fmt.Sprintf("custom message : %s", msg))
	}
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do something with it.
	}
}
