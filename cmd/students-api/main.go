package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sujaysinhpatil2/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Println("server started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to staart server")
	}

}
