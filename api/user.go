package api

import (
	"ddCode-server/global"
	"ddCode-server/models/dto"
	"ddCode-server/pkg/resp"
	"ddCode-server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{}

// Login 登录
// @Tags 用户
// @Summary 用户登录
// @Description 用户登录
// @Param user body dto.UserLoginParam true "用户"
// @Router /user/login [post]
// @Produce json
// @Success 200 {object} resp.Response
func (u *User) Login(c *gin.Context) {
	param := new(dto.UserLoginParam)
	err := param.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	userService := service.UserService{}
	userInfo, err := userService.LoginByUserNameAndPassword(param)
	if err != nil {
		resp.Error(c, http.StatusInternalServerError, err, "登录失败")
		return
	}
	resp.OK(c, userInfo, "登录成功")
}

// CreateUser 创建用户
func (u *User) CreateUser(c *gin.Context) {
	regParam := new(dto.UserRegisterParam)
	err := regParam.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusInternalServerError, err, "参数验证失败")
		return
	}
	userService := service.UserService{}
	err = userService.CreateUser(regParam)
	if err != nil {
		resp.Error(c, http.StatusInternalServerError, err, "注册失败")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// Logout 登出
func (u *User) Logout(c *gin.Context) {

}

// UpdateUser 更新权限是本人或admin
func (u *User) UpdateUser(c *gin.Context) {
	upParam := new(dto.UserInfoUpdateParam)
	err := upParam.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusInternalServerError, err, "参数验证失败")
		return
	}
	us := service.UserService{}
	err = us.UpdateUser(upParam)
	if err != nil {
		global.Logger.Errorf("更新用户基本信息失败, parms: %v", upParam)
		resp.Error(c, http.StatusInternalServerError, err, "更新失败")
		return
	}
	resp.OK(c, nil, "操作成功")
}

func (u *User) UserList(c *gin.Context) {
	p := dto.UserListParam{}
	s := service.UserService{}
	err := p.Bind(c)
	list, count, err := s.FindUserByPage(&p)
	if err != nil {
		return
	}
	resp.PageOK(c, list, int(count), p.PageNo, p.PageNo, "")
}
