package main

import (
	"fmt"
	"os"
)

func getAndPrintWorkingDirectory() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)

		return err
	}

	fmt.Println(cwd)

	return nil
}

func main() {
	err := getAndPrintWorkingDirectory()
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Change Working Directory here */
	err = os.Chdir("error_logs")
	if err != nil {
		fmt.Println(err)

		return
	}

	err = getAndPrintWorkingDirectory()
	if err != nil {
		fmt.Println(err)
		return
	}
}
