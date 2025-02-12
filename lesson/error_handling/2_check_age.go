package main

import (
	"errors"
	"fmt"
)

var errorAgeLimit error = errors.New("age must be at least 18")

func checkAge(age int) error {
	if age < 18 {
		return errorAgeLimit

	}
	return nil
}

func main() {
	age := 16
	err := checkAge(age)
	switch err {
	case errorAgeLimit:
		fmt.Println(err)
	default:
		fmt.Println("Age is valid.")
	}
}
