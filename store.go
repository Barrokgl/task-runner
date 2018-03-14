package taskRunner

import (
	"time"
)

type PersistedTask struct {
	ID        uint64     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
	TimeOut    time.Time `json:"timeout" `
	WorkType   string    `json:"workType"`
	Status     string    `json:"status"`
	Repeatable bool      `json:"repeatable"`
	Payload string     	 `json:"payload"`
	Async bool 			 `json:"async"`
}

type TaskStore interface {
	Add(attributes PersistedTask) (*PersistedTask, error)
	Fetch() ([]PersistedTask, error)
	Remove(attributes PersistedTask) error
}