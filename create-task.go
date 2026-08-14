package main

import (
	"bufio"
	"cli-todo/utils"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func CreateTask(item utils.Task) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter new task | ")
	taskTitle, _ := reader.ReadString('\n')
	formatedTitle := strings.TrimSpace(taskTitle)
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
	tasks = append(tasks, utils.Task{
		Title:      formatedTitle,
		ID:         utils.CreateTaskId(),
		Time_stamp: currentTime.Format(time.RFC3339),
		Completed:  false,
	})
	data, err = json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	fmt.Print("Task Created succesfully")
	return os.WriteFile("task.json", data, 0644)

}
