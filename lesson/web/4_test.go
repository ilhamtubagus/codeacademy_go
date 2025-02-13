package main

import (
	"bytes"
	"log"
	"net/http"
)

func main() {
	jsonData := []byte(`{"title":"New Blog Post","content":"This is the content of the new blog post."}`)
	resp, err := http.Post("http://localhost:4001/data", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("Response Status:", resp.Status)
}
