package routes

import (
	"api_gateway.com/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Lets define an health check endpoint
	router.GET("/health", controllers.HealthCheck)

	// And then a proxy route this is important for the gateway and the network policy engine to function
	router.Any("/apiGateway/*proxyPath", controllers.GatewayHandler)
}