package main

import (
	"fmt"
	"net/http"

	"github.com/dazeyawn/maketen/handler"
)

func main() {
	http.HandleFunc("/solve", handler.SolveHandler)

	const port = "8080"
	fmt.Printf("Server running on: http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
