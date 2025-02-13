package main

import (
	"bytes"
	"log"
	"net/http"
)

func main() {
	// Make POST request here
	jsonData := []byte(`{"title":"New Blog Post","content":"This is the content of the new blog post."}`)
	buffer := bytes.NewBuffer(jsonData)

	resp, err := http.Post("http://localhost:4001/data", "application/json", buffer)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}
