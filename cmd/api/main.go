package main

import (
	"fmt"
	"go-cloud-native-rest-api/api/router"
	"go-cloud-native-rest-api/config"
	"log"
	"net/http"
)

func main() {
	config := config.New()
	router := router.New()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      router,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
		IdleTimeout:  config.Server.TimeoutIdle,
	}

	log.Printf("Starting server at '%s'\n", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed: %s", err)
	}
}
