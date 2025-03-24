package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveTasks() {
	data, err := json.Marshal(taskManager)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err)
	}
}

func loadTasks() {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			taskManager = TaskManager{Tasks: []Task{}, TotalTasks: 0}
		} else {
			fmt.Println("Error reading tasks from file:", err)
		}
		return
	}

	err = json.Unmarshal(data, &taskManager)
	if err != nil {
		fmt.Println("Error parsing tasks:", err)
		taskManager = TaskManager{Tasks: []Task{}, TotalTasks: 0}
	}
}

func FindTask(id int) *Task {
	for i, task := range taskManager.Tasks {
		if task.Id == id {
			return &taskManager.Tasks[i]
		}
	}
	return nil
}

func UpdateIds() {
	newId := 1
	for i := range taskManager.Tasks {
		taskManager.Tasks[i].Id = newId
		newId++
	}
	saveTasks()
}
