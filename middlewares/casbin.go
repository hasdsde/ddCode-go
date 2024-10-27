package middlewares

import (
	"ddCode-server/global"
	"ddCode-server/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

const NoPolicyMessage = "暂时没有此权限"
const PolicyCheckFail = "权限鉴定失败"

func CasbinMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == LoginApiAddress {
			c.Next()
			return
		}
		tokenInfo, err := GetTokenInfo(c)
		if err != nil || tokenInfo.UserName == "" {
			global.Logger.Error(PolicyCheckFail)
			resp.Error(c, http.StatusForbidden, nil, PolicyCheckFail)
			c.Abort()
			return
		}
		url := c.Request.URL.Path
		method := c.Request.Method
		access, err := global.Rbac.HasAccess(tokenInfo.UserName, url, method)
		if err != nil {
			global.Logger.Error(PolicyCheckFail)
			return
		}
		if access {
			c.Next()
		} else {
			resp.Error(c, http.StatusForbidden, nil, NoPolicyMessage)
			redirectToLogin(c)
			c.Abort()
			return
		}
	}
}
