package main

import (
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdateAt    string `json:"updated_at"`
}

type TaskManager struct {
	Tasks      []Task `json:"tasks"`
	TotalTasks int    `json:"total_tasks"`
}

var taskManager TaskManager

func init() {
	loadTasks()
}

func AddTask(descriptor string) {
	taskManager.TotalTasks++
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	task := Task{
		Id:          taskManager.TotalTasks,
		Description: descriptor,
		Status:      "todo",
		CreatedAt:   currentTime,
		UpdateAt:    currentTime,
	}

	taskManager.Tasks = append(taskManager.Tasks, task)
	saveTasks()
	println("Task added successfully: ", descriptor)
}

func ListTasks(filter *string) {
	tasknumber := 0
	if taskManager.TotalTasks == 0 {
		println("No tasks to show")
		return
	}

	if filter == nil {
		ClearScreen()
		println("Listing tasks...\n")
		for _, task := range taskManager.Tasks {
			println("Task ID:", task.Id)
			println("Description:", task.Description)
			println("Status:", task.Status, "\n")
		}
	} else {
		switch *filter {
		case "todo":
			for _, task := range taskManager.Tasks {
				if task.Status == "todo" {
					tasknumber++
					println("Task ID:", task.Id)
					println("Description:", task.Description)
					println("Status:", task.Status, "\n")
				}
			}
			if tasknumber == 0 {
				println("No tasks to show")
			}
		case "done":
			for _, task := range taskManager.Tasks {
				if task.Status == "done" {
					tasknumber++
					println("Task ID:", task.Id)
					println("Description:", task.Description)
					println("Status:", task.Status, "\n")
				}

			}
			if tasknumber == 0 {
				println("No tasks to show")
			}
		case "in-progress":
			for _, task := range taskManager.Tasks {
				if task.Status == "in-progress" {
					tasknumber++
					println("Task ID:", task.Id)
					println("Description:", task.Description)
					println("Status:", task.Status, "\n")
				}
			}
			if tasknumber == 0 {
				println("No tasks to show")
			}
		default:
			println("Invalid filter")
		}
	}
}

func MarkTask(command string, id string) {
	taskStat := command[1:]
	if taskStat != "done" && taskStat != "in-progress" {
		println("Invalid status")
		return
	}
	task := FindTask(convert(id))
	println("Set task", id, "to", taskStat)
	task.Status = taskStat
	saveTasks()
}

func UpdateTask(id int, descriptor string) {
	task := FindTask(id)
	if task == nil {
		println("Task not found")
		return
	}
	task.Description = descriptor
	task.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	saveTasks()
	println("Task updated successfully")
}

func DeleteTask(id int) {
	index := -1
	for i, t := range taskManager.Tasks {
		if t.Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		println("Task not found")
		return
	}
	taskManager.Tasks = append(taskManager.Tasks[:index], taskManager.Tasks[index+1:]...)
	taskManager.TotalTasks--
	saveTasks()
	println("Task deleted successfully")
	UpdateIds()
}
