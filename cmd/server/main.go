package main

import (
	"log"

	"api_gateway.com/internal/routes"
	"api_gateway.com/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	env:= utils.GetEnv()
	router := gin.Default()
	routes.RegisterRoutes(router)
	utils.SetMode(env, router)
	log.Printf("Starting server in %s mode", env)
	port := utils.GetPort()
	err:=router.Run("0.0.0.0:"+port);
	if err == nil {
		log.Printf("Server is running at %s", port);
	} else {
		log.Printf("Failed to start server: %v", err)
	}

}