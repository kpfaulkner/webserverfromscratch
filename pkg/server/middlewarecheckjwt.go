package server

import (
	"net/http"
	"strings"
)

func (s *Server) checkJWT(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check JWT token...  somehow. If failed, then return 401
		jwt := r.Header["Authorization"]
		if len(jwt) > 0 && strings.TrimSpace(jwt[0]) != "" {
			// check JWT.
		} else {
			// no authorization header... so *has* to fail.
			http.Error(w, "No JWT, bad person!", http.StatusUnauthorized)
			return
		}

		h(w, r)
	}
}
