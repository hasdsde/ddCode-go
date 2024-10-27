package router

import (
	"ddCode-server/api"
	"github.com/gin-gonic/gin"
)

func registerDeptRouter(g *gin.RouterGroup) {
	deptApi := api.Dept{}
	deptGroup := g.Group("/dept")
	deptGroup.GET("", deptApi.GetDeptList)
	deptGroup.POST("", deptApi.CreateDept)
	deptGroup.PUT("", deptApi.UpdateDept)
	deptGroup.DELETE("", deptApi.DeleteDept)
}
