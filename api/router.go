package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS()) //解决跨域
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/forgot", forgot)
	
	u := r.Group("/user")
	{
		u.Use(middleware.JWT())
		u.POST("/say", say)
		u.GET("/message", message)
	}

	r.Run(":8080")
}
