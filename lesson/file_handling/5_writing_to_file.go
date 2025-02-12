package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	file, err := os.Create("user_data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write title to the file
	numBytesWritten, err := file.WriteString("User_Data\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Number of bytes written for title: %d\n", numBytesWritten)

	// Create a User struct instance
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Serialize the struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error serializing struct to JSON:", err)
		return
	}

	// Write JSON data to the file
	numBytesWritten, err = file.Write(jsonData)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Number of bytes written for JSON data: %d\n", numBytesWritten)

	content, err := os.ReadFile("user_data.json")
	if err != nil {
		return
	}
	defer file.Close()

	fmt.Println(string(content))
}
