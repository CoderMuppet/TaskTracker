package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowMenu() {
	println("Please select an option:\n")
	println("add {task description} - Add new task")
	println("list {optional : done, todo, in-progress} - List tasks")
	println("update {task id}{new text} - Update task")
	println("delete {task id} - Delete a task")
	println("mark-in-progress {task id} - Mark task as in progress")
	println("mark-done {task id} - Mark task as done")
	println("exit - close the program\n")
}

func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func convert(input string) int {

	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return i
}
