package controllers

import (
	"github.com/gin-gonic/gin"
)

func GatewayHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is a proxy endpoint",
	})
}