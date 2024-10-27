package router

import (
	"ddCode-server/api"
	"github.com/gin-gonic/gin"
)

func registerPolicyRouter(g *gin.RouterGroup) {
	policyApi := api.Policy{}
	policyGroup := g.Group("/policy")
	policyGroup.GET("", policyApi.GetPolicyList)
	policyGroup.POST("", policyApi.CreatePolicy)
	policyGroup.PUT("", policyApi.UpdatePolicy)
	policyGroup.DELETE("", policyApi.DeletePolicy)
}
