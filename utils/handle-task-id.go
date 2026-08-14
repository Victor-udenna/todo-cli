package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateTaskId() int {
	CreateFile()
	data, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	var tasks []Task
	if len(data) > 0 {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println(err)
			return 1
		}
	}

	return len(tasks) + 1
}
