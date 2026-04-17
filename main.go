package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/JanisJuska/Go-task-cli/task"
	"github.com/JanisJuska/Go-task-cli/utils"
)

var tasks = returnListFromFile("todos.json")
var idCount uint = utils.ReturnIdCount("todos.json")

func main() {

	allArgs := os.Args[1:]
	firstArg := allArgs[0]
	afterArgs := allArgs[1:]
	argString := strings.Join(afterArgs, " ")

	switch strings.ToLower(firstArg) {
	case "add":
		idCount++
		var newTask task.Task
		newTask.ID = idCount
		newTask.Title = argString
		newTask.Done = false

		tasks = addNewTaskToJSON(newTask, "todos.json")

		fmt.Printf("'%v' succesfully added to the list\n", argString)
	case "list":

		fmt.Println()
		fmt.Printf("%-4s | %-30s | %s\n", "ID", "Title", "Done")
		fmt.Println("----------------------------------------------")

		for _, task := range tasks {
			fmt.Println(task.String())
		}

		fmt.Println()
	case "done":
		id, err := strconv.Atoi(argString)
		if err != nil {
			log.Fatalf("Cannot convert string to number due to: %v\n", err)
		}

		var taskTitle string

		for i, t := range tasks {
			if t.ID == uint(id) {
				t.Done = true

				taskTitle = t.Title

				tasks[i] = t
			}
		}

		fileData := openAndReadFile("todos.json")

		err = os.WriteFile("todos.json", fileData, 0644)
		if err != nil {
			log.Fatalf("Cannot write to file due to: %v\n", err)
		}

		fmt.Printf("'%v' task marked as Done ✔️\n", taskTitle)
	default:
		log.Fatalf("No argument passed.\n")
	}

}

func addNewTaskToJSON(newTask task.Task, filename string) []task.Task {

	tasksList := returnListFromFile(filename)
	fileData := openAndReadFile(filename)

	tasksList = append(tasksList, newTask)

	fileData, err := json.MarshalIndent(tasksList, "", "  ")
	if err != nil {
		log.Fatalf("Cannot Marshal file due to: %v\n", err)
	}

	err = os.WriteFile(filename, fileData, 0644)
	if err != nil {
		log.Fatalf("Cannot write to file due to: %v\n", err)
	}

	return tasksList

}

func returnListFromFile(filename string) []task.Task {
	fileData := openAndReadFile(filename)

	var tasksList []task.Task
	json.Unmarshal(fileData, &tasksList)

	return tasksList
}

func openAndReadFile(filename string) []byte {
	dataFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Cannot open the file due to: %v\n", err)
	}
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Cannot read the file due to: %v\n", err)
	}

	defer dataFile.Close()

	return fileData
}
