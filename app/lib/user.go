package lib

import (
	"errors"
	"strconv"

	"github.com/mjiee/grf-gin/app/model"
	"gorm.io/gorm"
)

// UserService 提供user相关服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建user服务
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

// GetUserInfo 获取单个用户信息
func (s *UserService) GetUserInfo(id string, isAdmin bool) (JwtUser, error) {
	var result JwtUser
	userId, err := strconv.Atoi(id)

	if isAdmin {
		var manager = &model.Manager{}
		err = s.db.Where("id = ?", userId).Select("id").Take(manager).Error
		result = manager
	} else {
		var user = &model.User{}
		err = s.db.Where("id = ?", userId).Select("id").Take(user).Error
		result = user
	}

	if err != nil {
		err = errors.New("该用户不存在")
	}

	return result, err
}
