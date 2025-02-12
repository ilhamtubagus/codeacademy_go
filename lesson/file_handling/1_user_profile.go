package main

import (
	"fmt"
	"os"
)

func main() {
	/* Create user_directory here */
	err := os.Mkdir("user_profiles", 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
