package main

import (
	"bufio"
	"cli-todo/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

var currentTime = time.Now()

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter new task | ")
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}
	task = strings.TrimSpace(task)

	CreateTask(utils.Task{
		Title:      task,
		ID:         utils.CreateTaskId(),
		Time_stamp: currentTime.Format(time.RFC3339),
		Completed:  false,
	})
}
