package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dazeyawn/maketen/handler"
)

func main() {
	http.HandleFunc("/solve", handler.SolveHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
