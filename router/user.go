package router

import (
	"ddCode-server/api"
	"github.com/gin-gonic/gin"
)

func registerUserRouter(g *gin.RouterGroup) {
	userApi := api.User{}
	userGroup := g.Group("/user")
	userGroup.POST("/login", userApi.Login)
	userGroup.POST("/", userApi.CreateUser)
	userGroup.PUT("/", userApi.UpdateUser)
}
