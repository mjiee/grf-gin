package model

import (
	"time"

	"gorm.io/gorm"
)

// primary key
type ID struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
}

// create, update time
type Timestamps struct {
	CreateAt time.Time `json:"create_at" gorm:"default:current_timestamp"`
	UpdateAt time.Time `json:"update_at" gorm:"default:current_timestamp"`
}

// soft delete time
type SoftDeletes struct {
	DeleteAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
