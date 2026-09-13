package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	port := "8081"

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "up"})
	})

	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
