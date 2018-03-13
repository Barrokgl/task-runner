package model

import "time"

type BasicModel struct {
	ID        uint64     `json:"id" required:"true" gorm:"primary_key;column:id"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index" gorm:"column:deleted_at"`
}
