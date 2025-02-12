package main

import (
	"errors"
	"fmt"
	"os"
)

type FileError struct {
	Message string
}

func (e *FileError) Error() string {
	return e.Message
}

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("readFile: %w", &FileError{Message: "file not found"})

	}
	defer file.Close()
	return nil
}

func main() {
	if err := readFile("nonexistent.txt"); err != nil {
		fmt.Println("Error:", err)

		var fileErr *FileError
		// Check if the error is of type FileError
		if errors.As(err, &fileErr) {
			fmt.Println(fileErr.Message)
		}
	}
}
