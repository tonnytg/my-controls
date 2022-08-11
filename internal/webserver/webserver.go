package webserver

import (
	"log"
	"net/http"
	"os"
	"time"
)

var (
	WebPort string
)

func Start() {

	LoadHandlers()

	if os.Getenv("WEB_PORT") == "" {
		WebPort = ":8080"
	}

	s := &http.Server{
		Addr:           WebPort,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting web server on %s", WebPort)
	log.Fatal(s.ListenAndServe())
}
