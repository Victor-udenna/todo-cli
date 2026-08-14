package utils

import (
	"fmt"
	"os"
)

func CreateFile() error {
	if _, err := os.Stat("task.json"); os.IsNotExist(err) {
		file, err := os.Create("task.json")
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer file.Close()
	}
	return nil
}
