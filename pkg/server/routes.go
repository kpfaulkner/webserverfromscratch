package server

func (s *Server) routes() {
	s.router.HandleFunc("/api/", s.handleAPI())
	s.router.HandleFunc("/greetings", s.handleGreeting("hello"))
	s.router.HandleFunc("/", s.handleIndex())
}
