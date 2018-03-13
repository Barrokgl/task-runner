package taskRunner

import (
"bitbucket.org/GromKri/go-qiwi-service/model"
)

type TaskStore interface {
	Add(attributes model.SystemTask) (*model.SystemTask, error)
	Fetch() ([]model.SystemTask, error)
	Remove(attributes model.SystemTask) error
}