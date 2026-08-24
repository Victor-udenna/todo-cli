package main

import (
	"bufio"
	"cli-todo/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UpdateTask() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task id to change status | ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}
	data, err := os.ReadFile("task.json")
	choice := strings.TrimSpace(input)
	var taskList []utils.Task
	// var filtered []utils.Task
	err = json.Unmarshal(data, &taskList)
	if err != nil {
		return err
	}

	convertIdToInt, err := strconv.Atoi(choice)
	if err != nil {
		return err
	}

	found := false

	for i := range taskList {
		if convertIdToInt == taskList[i].ID {
			found = true

			if taskList[i].Completed {
				fmt.Println("Task is already completed")
				return nil
			}

			taskList[i].Completed = true

			fmt.Println("Updated task:", taskList[i])
			break
		}
	}

	if !found {
		fmt.Println("Task not found")
		return nil
	}

	data, err = json.MarshalIndent(taskList, "", " ")
	if err != nil {
		return err
	}

	fmt.Println("Task updated successfully")

	return os.WriteFile("task.json", data, 0644)
}
