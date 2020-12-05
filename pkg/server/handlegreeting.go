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

		if r.Method == http.MethodPost {
			decoder := json.NewDecoder(r.Body)
			var t request
			err := decoder.Decode(&t)
			if err != nil {
				fmt.Printf("BOOM %s\n", err.Error())
				fmt.Fprintf(w, "BREAKAGE")
				return
			}

			fmt.Fprintf(w, fmt.Sprintf("custom message for %s: %s", t.Name, msg))

		} else {
			fmt.Fprintf(w, "Can only perform POSTs")
		}
	}
}
