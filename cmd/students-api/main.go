package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"students-api/internal/config"
)

func main() {
	// Define the `-config` flag to pass the path to the configuration file
	configPath := flag.String("config", "", "path to the configuration file")
	flag.Parse()

	// Load the configuration
	cfg := config.MustRead(*configPath)

	// Database setup

	// Setup router
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students API"))
	})

	// Setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	fmt.Println("server started at", cfg.HTTPServer.Address)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
