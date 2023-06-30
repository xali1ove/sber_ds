package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sber_ds/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err.Error())
	}

	router := gin.Default()

	addr := ":" + cfg.Port
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %s", err.Error())
	}
}
