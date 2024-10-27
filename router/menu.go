package router

import (
	"ddCode-server/api"
	"github.com/gin-gonic/gin"
)

func registerMenuRouter(g *gin.RouterGroup) {
	menuApi := api.Menu{}
	menuGroup := g.Group("/menu")
	menuGroup.GET("/parent", menuApi.GetAllParentMenu)
	menuGroup.GET("", menuApi.GetMenuList)
	menuGroup.POST("", menuApi.CreateMenu)
	menuGroup.PUT("", menuApi.UpdateMenu)
	menuGroup.DELETE("", menuApi.DeleteMenu)
}
