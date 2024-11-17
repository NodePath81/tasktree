package model

type Task struct {
	TaskTree
	TaskMeta
	Title       string `json:"title"`
	Description string `json:"description"`
}
