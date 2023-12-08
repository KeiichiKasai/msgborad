package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/service"
	"main.go/utils"
)

func say(ctx *gin.Context) {
	user, e := ctx.Get("username")
	if !e {
		fmt.Println("没有")
		ctx.JSON(400, gin.H{
			"msg": "Something wrong",
		})
		return
	}
	username := user.(string)

	context := ctx.PostForm("context")
	if utils.IsEmpty(context) {
		ctx.JSON(400, gin.H{
			"msg": "Can't be empty",
		})
		return
	}
	err := service.LeaveMessage(username, context)
	if err != nil {
		ctx.JSON(500, gin.H{
			"msg": "leave message failed",
		})
	}
	ctx.JSON(200, gin.H{
		"msg": "leave message successfully",
	})
}
func message(ctx *gin.Context) {
	msg, err := service.GetMessage()
	if err != nil {
		ctx.JSON(500, gin.H{
			"msg": "get massage failed",
		})
	}
	ctx.JSON(200, msg)
}
