package main

import (
	"fmt"
	"os"
)

func main() {
	/* Create File here */
	// will truncate if file already exists
	file, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
