package model

import "time"

const (
	TaskStatusPending  = "pending"
	TaskStatusProgress = "progress"
	TaskStatusDone     = "done"
	TaskStatusFailed   = "failed"
)

type TaskMeta struct {
	Status string `json:"status"`

	AddedAt time.Time `json:"added_at"`
	StartAt time.Time `json:"start_at"`
	NextAt  time.Time `json:"next_at"`
	EndAt   time.Time `json:"end_at"`

	RecurInterval time.Duration `json:"recur_interval"`
	RecurCount    uint          `json:"recur_count"`
}
