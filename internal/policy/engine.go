package policy

import (
	"api_gateway.com/internal/constants"
	"api_gateway.com/internal/models"
	"github.com/gin-gonic/gin"
)

type PolicyEngine struct {
	context *models.PolicyContext
}

func NewPolicyEngine(ctx *gin.Context) *PolicyEngine {
	// Extract headers as map[string]string
	headers := make(map[string]string)
	for k, v := range ctx.Request.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	// Use RemoteAddr as client
	client := ctx.Request.RemoteAddr
	return &PolicyEngine{
		context: models.NewPolicyContext(
			ctx.Request.Method,
			ctx.Request.URL.Path,
			headers,
			client,
		),
	}
}

func (pe *PolicyEngine) Evaluate() *models.PolicyDecision {
	// Evaluate policy
	if pe.context.Method == "GET" || pe.context.Path == "/admin" {
		return models.NewPolicyDecision(constants.Deny, constants.Unauthorized)
	}
	return models.NewPolicyDecision(constants.Allow, "No foul detected")
}
