package main

import (
	"cli-todo/utils"
	"encoding/json"
	"os"
)

func CreateTask(item utils.Task) error {
	data, err := os.ReadFile("task.json")
	if err != nil {
		return err
	}
	var tasks []utils.Task
	if len(data) > 0 {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			return err
		}
	}
	tasks = append(tasks, item)
	data, err = json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("task.json", data, 0644)
}
