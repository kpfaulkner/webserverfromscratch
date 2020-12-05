package main

import (
	"fmt"
	"github.com/kpfaulkner/webserverfromscratch/pkg/server"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func  InitializeLogging(logFile string) {

	var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)

	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	fmt.Printf("so it begins...\n")

	InitializeLogging("test.log")

	/*
	srv := server.NewServer()
	// adding logging middleware to server.
	loggerMW := server.WithLogging()
	serverWithLogging := loggerMW(srv) */

	serverWithLogging := server.NewServerWithLogging()
	err := http.ListenAndServe(":8080", serverWithLogging)
	fmt.Printf("err is %s\n", err.Error())
}
