package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//  Make GET request here
	resp, err := http.Get("http://localhost:4001/")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(string(body))
}
