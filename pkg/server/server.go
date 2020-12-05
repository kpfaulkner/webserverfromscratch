package server

import (
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	s := Server{}
	s.router = http.NewServeMux()
	s.routes()
	return &s
}

func NewServerWithLogging() http.Handler {
	server := NewServer()
	loggerMW := WithLogging()
	serverWithLogging := loggerMW(server)
	return serverWithLogging
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
