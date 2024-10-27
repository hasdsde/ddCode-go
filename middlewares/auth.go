package middlewares

import (
	"ddCode-server/cons"
	"ddCode-server/global"
	"ddCode-server/models/vo"
	"ddCode-server/pkg/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strings"
)

const NoAuthMassage = "用户认证失败"
const UnsAuthCode = http.StatusUnauthorized
const LoginApiAddress = "/api/v1/user/login"
const UserTokenName = "userTokenInfo"

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// nolint:all
		if c.Request.URL.Path == LoginApiAddress {
			c.Next()
			return
		}
		header := c.GetHeader(cons.Auth)
		if header == "" {
			resp.Error(c, http.StatusUnauthorized, nil, NoAuthMassage)
			redirectToLogin(c)
			c.Abort()
			return
		}
		if !strings.HasPrefix(header, cons.Bearer) {
			resp.Error(c, http.StatusUnauthorized, nil, NoAuthMassage)
			redirectToLogin(c)
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(header, cons.Bearer)

		payload, err := global.JwtMaker.VerifyToken(tokenString)
		if err != nil {
			resp.Error(c, http.StatusUnauthorized, nil, NoAuthMassage)
			redirectToLogin(c)
			c.Abort()
			return
		}
		if payload.Valid() != nil { // 验证过期时间
			resp.Error(c, http.StatusUnauthorized, nil, NoAuthMassage)
			redirectToLogin(c)
			c.Abort()
			return
		}

		userTokenInfo := vo.UserTokenInfo{}
		err = mapstructure.Decode(payload.Info, &userTokenInfo)
		if err != nil {
			resp.Error(c, http.StatusUnauthorized, nil, NoAuthMassage)
			redirectToLogin(c)
			c.Abort()
			return
		}
		c.Set(UserTokenName, userTokenInfo)
		//key := vo.GetUserRedisId(payload)
		//is := vo.IsExistKey(c, key)
		//if !is {
		//	resp.Error(c, http.StatusUnauthorized, nil, NoAuthMassage)
		//	redirectToLogin(c)
		//	c.Abort()
		//	return
		//}
		c.Next()
	}
}
func redirectToLogin(c *gin.Context) {
	if c.Request.URL.Path == LoginApiAddress {
		c.Redirect(http.StatusTemporaryRedirect, LoginApiAddress)
	}
}

func GetTokenInfo(c *gin.Context) (vo.UserTokenInfo, error) {
	user, exist := c.Get(UserTokenName)
	if !exist {
		return vo.UserTokenInfo{}, fmt.Errorf("获取token中用户信息错误")
	}
	return user.(vo.UserTokenInfo), nil
}
