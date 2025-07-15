package service

import (
	"github.com/common"
	"github.com/model"
)

func Migrate() error {
	return common.DB.AutoMigrate(&model.User{})
}
