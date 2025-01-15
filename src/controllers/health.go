package controllers

import "github.com/gin-gonic/gin"

func Health(context *gin.Context){
	context.JSON(200, gin.H{
		"message": "Health check is OK",
	})
}