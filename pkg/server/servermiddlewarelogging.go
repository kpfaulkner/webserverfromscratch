package server

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"time"
)

func WithLogging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			end := time.Now()
			log.Info(fmt.Sprintf("%s tool %d ms", r.URL.EscapedPath(), end.Sub(start).Milliseconds()))
		}
		return http.HandlerFunc(fn)
	}
}
