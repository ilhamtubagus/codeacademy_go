package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Add GET route here
	mux.HandleFunc("GET /blog", handleGetBlogRequest)

	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server: ", err)
	}
}

func handleGetBlogRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of blog posts")
}
