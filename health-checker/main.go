package main

import (
	"health-checker/config"
	"health-checker/handlers"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.LoadFromYAML("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	router := handlers.NewRouter(cfg)
	router.RegisterRoutes()

	server := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}

	log.Printf("Server starting on port %s", "8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
