package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up the server here
	err := http.ListenAndServe(":4001", nil)

	if err != nil {
		log.Fatal(err)
	}
}
