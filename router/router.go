package router

import (
	"ddCode-server/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(middlewares.GinLogger())
	r.Use(middlewares.GinRecovery(true))
	r.Use(middlewares.Cors())
	r.Use(middlewares.TokenAuthMiddleware())
	r.Use(middlewares.CasbinMiddlewares())

	//路由写到这里
	group := r.Group("")
	g := group.Group("/api/v1")

	registerUserRouter(g)
	registerDeptRouter(g)
	registerMenuRouter(g)
	registerPolicyRouter(g)
	registerLogRouter(g)
	registerRoleRouter(g)
	registerSwaggerRouter(g)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})
	return r
}
