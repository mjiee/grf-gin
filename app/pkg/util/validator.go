package util

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CheckAddr 校验网络资源连接地址地址
func CheckAddr(fl validator.FieldLevel) bool {
	pattern := `^((0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])\.){3}(0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5]):([2-9]\d\d\d|[1-5]\d\d\d\d)$`
	addr := fl.Field().String()
	if ok, _ := regexp.MatchString(pattern, addr); !ok {
		return false
	}
	return true
}
