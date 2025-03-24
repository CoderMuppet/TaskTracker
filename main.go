package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	println("Welcome to TaskMonitor v0.1")
	start := true
	for start {

		ShowMenu()
		inputString := GetUserInput()

		input := strings.Fields(inputString)

		if input[0][0:4] == "mark" {

			MarkTask(input[0][4:], input[1])
			println("\nPress enter to continue...")
			fmt.Scanln()
			ClearScreen()
			continue
		}

		switch input[0] {
		case "add":
			AddTask(strings.Join(input[1:], " "))
		case "list":
			if len(input) > 1 {
				ListTasks(&input[1])
			} else {
				ListTasks(nil)
			}
		case "update":
			UpdateTask(convert(input[1]), input[2])
		case "delete":
			DeleteTask(convert(input[1]))
		case "exit":
			ClearScreen()
			os.Exit(0)
		default:
			println("Invalid input")
		}

		println("\nPress enter to continue...")
		fmt.Scanln()
		ClearScreen()

	}
}
