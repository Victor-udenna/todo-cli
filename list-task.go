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
	for _, item := range taskList {
		fmt.Println("-------------------------")
		fmt.Println("Id:", item.ID)
		fmt.Println("Title:", item.Title)
		fmt.Println("Time:", item.Time_stamp)
		fmt.Println("Completed:", item.Completed)
		fmt.Println("-------------------------")
	}
}
