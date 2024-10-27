package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func registerSwaggerRouter(g *gin.RouterGroup) {
	sg := g.Group("/swagger")
	sg.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
