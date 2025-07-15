package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/model"
	"github.com/service"
	"github.com/vo"
	"log"
	"net/http"
	"strconv"
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

func Login(c *gin.Context) {
	var req vo.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"数据解析失败": err.Error(),
		})
		return
	}

	var user model.User
	copier.Copy(&user, &req)
	err := service.ValidateLogin(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"数据校验失败": err.Error(),
		})
		return
	}

	loginResult, err := service.Login(&user)

	if err != nil {
		log.Println("登录用户失败:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"登录失败": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": loginResult,
	})
}

func GetUserByName(c *gin.Context) {
	name := c.Query("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户名不能为空",
		})
		return
	}

	user, err := service.GetUserByName(name)
	if err != nil {
		log.Println("获取用户失败:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteById(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户ID不能为空",
		})
		return
	}

	err = service.DeleteById(id)
	if err != nil {
		log.Println("删除用户失败:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}
