package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handleGreeting(msg string) http.HandlerFunc {

	type request struct {
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			decoder := json.NewDecoder(r.Body)
			var t request
			err := decoder.Decode(&t)
			if err != nil {
				fmt.Printf("BOOM %s\n", err.Error())
				fmt.Fprintf(w, "BREAKAGE")
				return
			}

			fmt.Fprintf(w, fmt.Sprintf("custom message for %s: %s", t.Name, msg))
			
		case http.MethodGet:
			fmt.Fprintf(w, "GET GET GET BLAH BLAH BLAH")
		}
	}
}
