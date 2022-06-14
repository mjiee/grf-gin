package lib

import "github.com/mjiee/scaffold-gin/app/pkg/request"

// Register 注册信息
type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`                 // user name
	Phone    string `form:"phone" json:"phone" binding:"required,len=11,number"` // phone
	Password string `form:"password" json:"password" binding:"required,gte=6"`   // password
}

// GetMessages 实现request.Validator接口
func (register Register) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Name.required":     "用户名不能为空",
		"Phone.required":    "手机号不能为空",
		"Phone.len":         "手机号长度必须是11位",
		"Phone.number":      "手机号必须是数字",
		"Password.required": "密码不能为空",
		"Password.gte":      "密码长度必须大于6",
	}
}

// Login 登录信息
type Login struct {
	Phone    string `form:"phone" json:"phone" binding:"required,len=11,number"`
	Password string `form:"password" json:"password" binding:"required,gte=6"`
}

// GetMessages 实现request.Validator接口
func (login Login) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Phone.required":    "手机号不能为空",
		"Phone.len":         "手机号长度必须是11位",
		"Phone.number":      "手机号必须是数字",
		"Password.required": "密码不能为空",
		"Password.gte":      "密码长度必须大于6",
	}
}
