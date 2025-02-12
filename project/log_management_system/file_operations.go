package main

import (
	"bufio"
	"fmt"
	"os"
)

func createLogFile() {
	fmt.Println("Enter the name of the log file you'd like to create:")
	var fileName string
	fmt.Scanln(&fileName)

	/* Create File here */
	file, err := os.Create(fileName)
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

	fmt.Printf("Log file '%s' created successfully.\n", fileName)
}

func readLogFile() {
	fmt.Println("Enter the name of the log file you'd like to read from:")
	var fileName string
	fmt.Scanln(&fileName)

	/* Read File here */
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Log file content:")
	fmt.Println(string(content))
}

func writeLogEntry() {
	fmt.Println("Enter the name of the log file you'd like to write to:")
	var fileName string
	fmt.Scanln(&fileName)

	fmt.Println("Enter the log entry:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	/* Write string here */
	bytesWritten, err := file.WriteString(data)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Log entry successfully written to file. %d bytes written.\n", bytesWritten)
}
