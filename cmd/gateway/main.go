package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("HEALTH_PORT")

	http.HandleFunc("/healthcheck", healthHandler)

	fmt.Printf("Listening on %v", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status": "up",
	}
	json.NewEncoder(w).Encode(response)
}
