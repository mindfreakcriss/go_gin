package service

import (
	"errors"
	"github.com/common"
	"github.com/model"
	"github.com/util"
	"github.com/vo"
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

func ValidateLogin(user *model.User) error {
	if len(user.Name) == 0 {
		return errors.New("用户名不能为空")
	}

	if len(user.Password) == 0 {
		return errors.New("密码不能为空")
	}
	return nil
}

func Login(user *model.User) (bool, error) {
	user.Password = util.MD5(user.Password)
	var users []model.User
	err := common.DB.Find(&users, &user).Error

	if err != nil || len(users) == 0 {
		return false, err
	}
	return true, nil
}

func GetUserByName(name string) (*vo.UserVo, error) {

	queryUser := &model.User{Name: name}

	var users []model.User
	err := common.DB.Find(&users, queryUser).Error
	if err != nil || len(users) == 0 {
		return nil, err
	}

	user := users[0]

	userVo := vo.UserVo{
		ID:       user.ID,
		Name:     user.Name,
		NickName: user.NickName,
	}

	return &userVo, nil
}

func DeleteById(id int) error {
	var user model.User
	user.ID = id

	result := common.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	result = common.DB.Delete(&user)
	return result.Error
}
