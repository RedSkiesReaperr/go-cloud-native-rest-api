package main

import (
	"fmt"
	"go-cloud-native-rest-api/api/router"
	"go-cloud-native-rest-api/config"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
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

	logLevel := gormlogger.Error
	if config.DB.Debug {
		logLevel = gormlogger.Info
	}

	const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	dbString := fmt.Sprintf(fmtDBString, config.DB.Host, config.DB.Username, config.DB.Password, config.DB.DBName, config.DB.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("DB connection start failure")
		return
	}

	router := router.New(db)
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
