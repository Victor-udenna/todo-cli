package main

import (
	"cli-todo/utils"
	"encoding/json"
	"fmt"
	"os"
)

func ListTask() {
	data, err := os.ReadFile("task.json")
	if err != nil {
		return
	}
	var taskList []utils.Task

	err = json.Unmarshal(data, &taskList)
	fmt.Println(taskList)
}
