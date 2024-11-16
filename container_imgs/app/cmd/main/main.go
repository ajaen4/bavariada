package main

import (
	"log"
	"os"
	"strconv"

	_ "app/internal/env"
	"app/internal/server"
)

func main() {

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal("APP_PORT env variable not set correctly")
	}

	server := server.NewServer(port)
	defer server.Close()
	log.Printf("Running server on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}
}
