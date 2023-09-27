package main

import (
	"fmt"
	"go-cloud-native-rest-api/api/router"
	"go-cloud-native-rest-api/config"
	"log"
	"net/http"
)

//  @title          go-cloud-native-rest-api API
//  @version        1.0
//  @description    This is a sample RESTful API with a CRUD

//  @contact.name   Dumindu Madunuwan
//  @contact.url    https://learning-cloud-native-go.github.io

//  @license.name   MIT License
//  @license.url    https://github.com/learning-cloud-native-go/myapp/blob/master/LICENSE

// @host       localhost:8080
// @basePath   /v1
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
