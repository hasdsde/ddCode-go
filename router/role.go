package router

import (
	"ddCode-server/api"
	"github.com/gin-gonic/gin"
)

func registerRoleRouter(g *gin.RouterGroup) {
	roleApi := api.Role{}
	roleGroup := g.Group("/role")
	roleGroup.GET("", roleApi.GetRoleList)
	roleGroup.POST("", roleApi.CreateRole)
	roleGroup.PUT("", roleApi.UpdateRole)
	roleGroup.DELETE("", roleApi.DeleteRole)

	roleGroup.GET("/roleMenu", roleApi.FindRoleMenu)    // 获取角色的菜单
	roleGroup.POST("/roleMenu", roleApi.UpdateRoleMenu) // 更新角色的菜单

	roleGroup.GET("/roleUser", roleApi.FindRoleUser)      // 获取用户的角色
	roleGroup.POST("/roleUser", roleApi.CreateRoleUser)   // 为用户新增角色
	roleGroup.DELETE("/roleUser", roleApi.DeleteRoleUser) // 为用户删除角色

	roleGroup.GET("/rolePolicy", roleApi.FindRolePolicy)    // 获取角色权限
	roleGroup.POST("/rolePolicy", roleApi.UpdateRolePolicy) // 更新角色权限

}
