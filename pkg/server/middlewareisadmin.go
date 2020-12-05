package server

import (
	"net/http"
)

func (s *Server) isAdmin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// just execute original handler for moment
		h(w, r)
	}
}
