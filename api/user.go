package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
)

func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	if utils.IsEmpty(username) || utils.IsEmpty(password) || utils.IsEmpty(phone) {
		ctx.JSON(400, gin.H{
			"msg": "Can't be empty",
		})
		return
	}
	if service.CheckUserIsExist(username) {
		ctx.JSON(202, gin.H{
			"msg": "Username has created",
		})
		return
	}
	err := service.CreateUser(username, password, phone)
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": "create user failed",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "create user successfully",
	})
}

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if utils.IsEmpty(username) || utils.IsEmpty(password) {
		ctx.JSON(400, gin.H{
			"msg": "Can't be empty",
		})
		return
	}
	if !service.CheckUserIsExist(username) {
		ctx.JSON(202, gin.H{
			"msg": "User hasn't created",
		})
		return
	}
	if !service.CheckPassword(username, password) {
		ctx.JSON(404, gin.H{
			"msg": "Wrong password",
		})
		return
	}
	token, err := middleware.GenToken(username)
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": "token failed",
		})
	}
	ctx.JSON(200, gin.H{
		"msg":   "login successfully",
		"token": token,
	})
}

func forgot(ctx *gin.Context) {
	username := ctx.PostForm("username")
	newPassword := ctx.PostForm("newpassword")
	phone := ctx.PostForm("phone")
	if utils.IsEmpty(username) || utils.IsEmpty(newPassword) || utils.IsEmpty(phone) {
		ctx.JSON(400, gin.H{
			"msg": "Can't be empty",
		})
		return
	}
	if !service.CheckUserIsExist(username) {
		ctx.JSON(202, gin.H{
			"msg": "User hasn't created",
		})
		return
	}
	err := service.ChangePassword(username, newPassword, phone)
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": "change password failed",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "change successfully",
	})
	return
}
