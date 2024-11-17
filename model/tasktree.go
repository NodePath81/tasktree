package model

type TaskTree struct {
	TaskID     string   `json:"task_id"`
	ParentID   string   `json:"parent_id"`
	SubtaskIDs []string `json:"subtask_ids"`
}
