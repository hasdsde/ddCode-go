package router

import (
	"ddCode-server/api"
	"github.com/gin-gonic/gin"
)

func registerLogRouter(g *gin.RouterGroup) {
	logApi := api.Log{}
	logGroup := g.Group("/log")
	logGroup.GET("", logApi.GetLogList)
}
