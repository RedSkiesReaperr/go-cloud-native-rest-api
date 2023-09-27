package main

import (
	"fmt"
	"go-cloud-native-rest-api/config"
	"io"
	"log"
	"net/http"
)

func main() {
	config := config.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      mux,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
		IdleTimeout:  config.Server.TimeoutIdle,
	}

	log.Printf("Starting server at '%s'\n", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed: %s", err)
	}

}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world!")
}
