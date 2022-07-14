package lib

import (
	"errors"

	"github.com/mjiee/scaffold-gin/app/model"
	"github.com/mjiee/scaffold-gin/app/pkg/util"
	"gorm.io/gorm"
)

// SignUp 用户注册
func (s *UserService) SignUp(params Register) (JwtUser, error) {
	var result *gorm.DB

	if params.IsAdmin {
		result = s.db.Where("phone = ?", params.Phone).Select("id").Take(&model.Manager{})
	} else {
		result = s.db.Where("phone = ?", params.Phone).Select("id").Take(&model.User{})
	}
	if result.RowsAffected != 0 {
		err := errors.New("用户已存在")
		return &model.User{}, err
	}

	pwd, err := util.BcryptPwd([]byte(params.Password))
	if err != nil {
		return &model.User{}, err
	}

	if params.IsAdmin {
		manager := model.Manager{User: model.User{Name: params.Name, Phone: params.Phone, Password: pwd}}
		err = s.db.Create(&manager).Error
		return &manager, err
	}

	user := model.User{Name: params.Name, Phone: params.Phone, Password: pwd}
	err = s.db.Create(&user).Error
	return &user, err
}

// SignIn 用户登录
func (s *UserService) SignIn(params Login) (JwtUser, error) {

	var result JwtUser
	var err error

	if params.IsAdmin {
		manager := &model.Manager{}
		err = s.db.Where("phone = ?", params.Phone).Select("id", "password").Take(manager).Error
		result = manager
	} else {
		user := &model.User{}
		err = s.db.Where("phone = ?", params.Phone).Select("id", "password").Take(user).Error
		result = user
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("该用户不存在")
	} else if !util.PwdCheck([]byte(params.Password), result.GetPwd()) {
		err = errors.New("密码错误")
	}

	return result, err
}
