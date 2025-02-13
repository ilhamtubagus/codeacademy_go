package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleBlogRequest)

	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server:", err)
	}
}

func handleBlogRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of blog posts")
}
