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

	for _, item := range taskList {
		if convertIdToInt == item.ID && item.Completed == false {
			item.Completed = true
			fmt.Println(item)
			// add the task to the json file
			taskList = append(taskList, item)
		} else {
			fmt.Printf("task is already completed")
			break
		}
	}

	fmt.Println(taskList)

	data, err = json.MarshalIndent(taskList, "", " ")
	fmt.Print("Task updated succesfully")
	return os.WriteFile("task.json", data, 06444)
}
