package main

import (
	"fmt"
	"log"
	"os"
)

// Function demonstrating log.Fatal
func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	_, err = fmt.Fscan(file, &content)
	if err != nil {
		return "", err
	}
	return content, nil
}

// Function demonstrating panic
func checkConfig() {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		panic("Config file is missing") // Trigger panic if config file doesn't exist
	}
}

// Main function demonstrating the three scenarios
func main() {

	// Scenario 1: Panic
	fmt.Println("Checking configuration file...")
	checkConfig()

	// Scenario 2: Log.Fatal
	fmt.Println("Reading data from input.txt...")
	content, err := readFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Scenario 3: fmt.Print
	fmt.Println("Successfully read file content:", content)
}
