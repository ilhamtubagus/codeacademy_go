package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createLogDirectory() {
	fmt.Println("Enter the name of the log directory you'd like to create:")
	var dirName string
	fmt.Scanln(&dirName)

	/* Make Directory here */
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	fmt.Printf("Directory %s successfully created", dirName)
}

func changeLogDirectory() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Directory:", cwd)

	fmt.Print("Enter the name of the log directory you want to change to: ")
	var newDirName string
	fmt.Scanln(&newDirName)

	/* Change Directory here */
	if err := os.Chdir(newDirName); err != nil {
		fmt.Println(err)

		return
	}

	newCwd := getWorkingDirectory()
	fmt.Println("Directory changed successfully. Current Directory:", newCwd)
}

func printWorkingDirectory() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Working Directory:", cwd)
}

func listDirectoryContents() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Working Directory:", cwd)

	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}

	fmt.Println("Directory Contents:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func getWorkingDirectory() string {
	/* Get Working Directory here */
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Print(err)

		return ""
	}

	return cwd
}
