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

func UpdateTask() {
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
		return
	}

	convertIdToInt, err := strconv.Atoi(choice)
	if err != nil {
		return
	}

	for _, item := range taskList {
		if convertIdToInt == item.ID {
			item.Completed = true
			fmt.Println(item)
		}
	}

	fmt.Println(choice, "this is the choice")
	fmt.Println(taskList)
}
