package server

import (
	"net/http"
	"strings"
)

func CheckJWT() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			// check JWT token...  somehow. If failed, then return 401
			jwt := r.Header["Authorization"]
			if len(jwt) > 0 && strings.TrimSpace(jwt[0]) != "" {
				// check JWT.
			} else {
				// no authorization header... so *has* to fail.
				http.Error(w, "No JWT, bad person!", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
