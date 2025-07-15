package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/service"
	"log"
	"net/http"
)

func Migrate(c *gin.Context) {
	// 1. Create the database
	err := service.Migrate()
	if err != nil {
		log.Println("迁移表结构：", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}
