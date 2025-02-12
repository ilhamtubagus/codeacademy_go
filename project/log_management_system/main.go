package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the Log Management System!")

	for {
		printOptions()

		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		// Convert the user input from a string to an integer
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid option provided.")
			continue
		}

		switch option {
		case 1:
			createLogDirectory()
		case 2:
			changeLogDirectory()
		case 3:
			createLogFile()
		case 4:
			readLogFile()
		case 5:
			writeLogEntry()
		case 6:
			printWorkingDirectory()
		case 7:
			listDirectoryContents()
		default:
			fmt.Println("Invalid option provided.")
		}
	}
}

func printOptions() {
	fmt.Println("--------------")
	fmt.Println("Please input a number to select an option or type 'exit' to quit:")
	fmt.Println("1 - Create a log directory")
	fmt.Println("2 - Change to a log directory")
	fmt.Println("3 - Create a log file")
	fmt.Println("4 - Read a log file")
	fmt.Println("5 - Write a log entry")
	fmt.Println("6 - Print working directory (pwd)")
	fmt.Println("7 - List directory contents (ls)")
}
