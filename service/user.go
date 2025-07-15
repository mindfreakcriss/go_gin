package service

import (
	"errors"
	"github.com/common"
	"github.com/model"
	"github.com/util"
)

func ValidateUser(user *model.User) error {
	if len(user.Name) == 0 {
		return errors.New("用户名不能为空")
	}

	if len(user.Password) == 0 {
		return errors.New("密码不能为空")
	}

	if len(user.NickName) == 0 {
		return errors.New("昵称不能为空")
	}
	return nil
}

func RegisterUser(user *model.User) error {
	// 这里可以添加更多的业务逻辑，比如检查用户名是否已存在等
	user.Password = util.MD5(user.Password)

	return common.DB.FirstOrCreate(&model.User{}, user).Error
}
