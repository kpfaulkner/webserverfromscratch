package server

import (
	"github.com/kpfaulkner/webserverfromscratch/pkg/server"
	"net/http"
)

func (s *server.Server ) isAdmin( h http.HandlerFunc) http.HandlerFunc {
	return func( w http.ResponseWriter, r *http.Request){
		// just execute original handler for moment
		h(w,r)
	}
}
