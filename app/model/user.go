package model

import (
	"strconv"
	"time"
)

type User struct {
	ID
	Name     string    `json:"name" gorm:"index;not null"`
	Phone    string    `json:"phone" gorm:"size:11;uniqueIndex;not null"`
	Email    string    `json:"email" gorm:"default:'';"`
	Password string    `json:"password" gorm:"not null;" `
	Birthday time.Time `json:"birthday" gorm:"not null;default:current_timestamp"`
	Timestamps
	SoftDeletes
}

// GetUid 实现lib.JwtUser接口
func (u *User) GetUid() string {
	return strconv.Itoa(int(u.ID.ID))
}
