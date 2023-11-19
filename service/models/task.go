package models

type TaskStatus int

const (
	InProgress = TaskStatus(iota)
	Done
)

type Task struct {
	Id      int        `json:"Id"`
	Status  TaskStatus `json:"Status"`
	Message string     `json:"Message"`
}
