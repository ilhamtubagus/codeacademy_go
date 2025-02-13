package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// URL-encoded payload data
	formData := url.Values{
		"blogName":    {"My First Blog"},
		"blogContent": {"This is the content of my first blog."},
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", "http://localhost:4001/form", strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Log the response status
	log.Println("Response status:", resp.Status)
}
