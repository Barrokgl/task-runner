package model

import "time"

type SystemTask struct {
	BasicModel
	TimeOut    time.Time `json:"timeout" gorm:"column:timeout"`
	WorkType   string    `json:"workType" gorm:"column:workType"`
	Status     string    `json:"status" gorm:"column:status"`
	Repeatable bool      `json:"repeatable" gorm:"column:repeatable"`
	PayloadInt int64     `json:"payload" gorm:"column:payloadInt"`
}

func (SystemTask) TableName() string {
	return "SystemTask"
}
