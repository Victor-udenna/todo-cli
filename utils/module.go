package utils

type Task struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Time_stamp string `json:"time_stamp"`
	Completed  bool   `json:"completed"`
}
