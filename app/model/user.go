package model

import (
	"strconv"
)

// User 用户
type User struct {
	ID
	Name     string `json:"name,omitempty" gorm:"index;not null"`
	Phone    string `json:"phone,omitempty" gorm:"size:11;uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null;" `
	Avatar   string `json:"avatar,omitempty" gorm:"default:''"`
	Timestamps
	SoftDeletes
}

// GetUid 实现lib.JwtUser接口
func (u *User) GetUid() string {
	return strconv.Itoa(int(u.ID.ID))
}

// GetPwd 实现lib.JwtUser接口
func (u *User) GetPwd() string {
	return u.Password
}

// Manager 管理员, 用户登陆后台系统的账户
type Manager struct {
	User
	// 管理人员角色: super: 0, admin: 1, general: 2
	Role    uint8 `json:"role" gorm:"not null;default:2"`
	Actived bool  `json:"-" gorm:"not null;default:false"`
}
