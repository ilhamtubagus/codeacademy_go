package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /form", formHandler)

	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server:", err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	blogName := r.FormValue("blogName")
	blogContent := r.FormValue("blogContent")

	// Retrieve form value here

	log.Printf("Blog Name: %s\n", blogName)
	log.Printf("Blog Content: %s\n", blogContent)
}
