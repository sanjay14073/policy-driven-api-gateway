// controllers/gateway.go
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"api_gateway.com/internal/policy"
)

func GatewayHandler(c *gin.Context) {

	// Step 1 : create the policy engine
	engine:= policy.NewPolicyEngine(c)
	// Step 2 : evaluate the request against policies
	decision := engine.Evaluate()


	if  !decision.Allowed {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "request blocked by policy",
		})
		return
	}

	// Step 5 will forward the request
	c.JSON(http.StatusOK, gin.H{
		"message": "request allowed (forwarding not implemented yet)",
	})
}
