package utils

import (
	"os"
	"github.com/gin-gonic/gin"
)


func GetEnv() string {
	env:=os.Getenv("ENV")
	if env==""{
		return "dev"
	}
	return env
}

func SetMode(env string, router *gin.Engine) {
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}