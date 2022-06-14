package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// BcryptPwd 加密密码
func BcryptPwd(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", errors.New("密码格式错误导致加密失败")
	}
	return string(hash), nil
}

// PwdCheck 密码检测
func PwdCheck(pwd []byte, hashedPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), pwd); err != nil {
		return false
	}
	return true
}
