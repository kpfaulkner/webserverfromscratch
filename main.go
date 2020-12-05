package main

import (
	"fmt"
	"github.com/kpfaulkner/webserverfromscratch/pkg/server"
	"net/http"
)

func main() {
	fmt.Printf("so it begins...\n")

  srv := server.NewServer()
  err := http.ListenAndServe(":8080", srv)
  fmt.Printf("err is %s\n", err.Error())
}
