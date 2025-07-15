package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/model"
	"github.com/service"
	"github.com/vo"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	var req vo.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"数据解析失败": err.Error(),
		})
		return
	}

	var user model.User
	copier.Copy(&user, &req)
	err := service.ValidateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"数据校验失败": err.Error(),
		})
		return
	}

	err = service.RegisterUser(&user)

	if err != nil {
		log.Println("注册用户失败:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"注册失败": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": user.ID,
	})
}
