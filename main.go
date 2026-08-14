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

	for {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Create task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Update task")
		fmt.Println("4. Exit")
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
		}

		choice := strings.TrimSpace(input)
		fmt.Println(choice)
		switch choice {
		case "1":
			CreateTask(utils.Task{
				Title:      "Testing create task",
				ID:         utils.CreateTaskId(),
				Time_stamp: currentTime.Format(time.RFC3339),
				Completed:  false,
			})
		case "2":
			fmt.Println("Listing tasks...")
			ListTask()
		case "3":
			fmt.Println("Update task...")

		case "4":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}

}
