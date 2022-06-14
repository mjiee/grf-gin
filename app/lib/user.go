package lib

import (
	"errors"
	"strconv"

	"github.com/mjiee/scaffold-gin/app/model"
	"github.com/mjiee/scaffold-gin/app/pkg/util"
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

// Register 用户注册
func (s *UserService) Register(params Register) (model.User, error) {
	var result = s.db.Where("phone = ?", params.Phone).Select("id").First(&model.User{})
	if result.RowsAffected != 0 {
		err := errors.New("用户已存在")
		return model.User{}, err
	}

	pwd, err := util.BcryptPwd([]byte(params.Password))
	if err != nil {
		return model.User{}, err
	}

	user := model.User{Name: params.Name, Phone: params.Phone, Password: pwd}
	err = s.db.Create(&user).Error

	return user, err
}

// Login 用户登录
func (s *UserService) Login(params Login) (*model.User, error) {

	var user model.User

	err := s.db.Where("phone = ?", params.Phone).Take(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("该用户不存在")
	} else if !util.PwdCheck([]byte(params.Password), user.Password) {
		err = errors.New("密码错误")
	}

	return &user, err
}

// GetUserInfo 获取单个用户信息
func (s *UserService) GetUserInfo(id string) (*model.User, error) {
	var user model.User
	userId, err := strconv.Atoi(id)

	err = s.db.Where("id = ?", userId).Take(&user).Error
	if err != nil {
		err = errors.New("该用户不存在")
	}
	return &user, err
}
