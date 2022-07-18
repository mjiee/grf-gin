package model

import (
	"time"

	"gorm.io/gorm"
)

// primary key
type ID struct {
	ID uint `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
}

// create, update time
type Timestamps struct {
	CreateAt time.Time `json:"-" gorm:"default:current_timestamp"`
	UpdateAt time.Time `json:"-" gorm:"default:current_timestamp"`
}

// soft delete time
type SoftDeletes struct {
	DeleteAt gorm.DeletedAt `json:"-" gorm:"index"`
}
