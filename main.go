package main

import (
	"github.com/controller"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func checkMiddle() gin.HandlerFunc {
	// 自定义中间件示例
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is not success",
			})
			return
		}
		//继续执行
		c.Next()
	}

}

func main() {
	// This is the entry point of the application.
	// You can initialize your application here.

	//日志
	f, _ := os.Create("gin.log")

	// 设置gin的日志输出到文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 设置gin的错误日志输出到文件和控制台
	gin.DefaultErrorWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	//使用日志,也是注册中间件的方法
	r.Use(gin.Logger())
	r.Use(checkMiddle()) // 使用自定义中间件

	// 处理GET请求，获取URL中的路径参数
	r.GET("/index", func(c *gin.Context) {
		// 获取请求URL中的查询参数keyword
		q := c.Query("keyword")

		c.JSON(http.StatusOK, gin.H{
			"message": "查询关键词为" + q,
		})
	})

	// 处理POST请求，获取表单数据
	r.POST("/login", func(c *gin.Context) {
		//获取表单数据
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"message": "用户名为" + username + "，密码为" + password,
		})
	})

	// 处理POST请求，获取JSON数据
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		c.String(http.StatusOK, "用户数据：%+v", user)
	})

	//复用smart 路由组
	// smart/hello
	smart := r.Group("/smart")
	{
		smart.GET("/hello", func(c *gin.Context) {
			// 获取请求头中的User-Agent
			userAgent := c.GetHeader("User-Agent")
			c.JSON(http.StatusOK, gin.H{
				"message":    "Hello, Smart API!",
				"user-agent": userAgent,
			})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		// 处理未找到的路由
		c.JSON(http.StatusNotFound, gin.H{
			"error": "this is my 404 Not Found",
		})
	})

	r.NoMethod(func(c *gin.Context) {
		// 处理不支持的方法
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "this is my 405 Method Not Allowed",
		})
	})

	//命令行处理数据库内容
	r.POST("db/sync", controller.Migrate)

	//用户相关的路由
	user := r.Group("/user")
	{
		// 用户注册
		user.POST("/register", controller.Register)
	}

	r.Run(":9080")
}
