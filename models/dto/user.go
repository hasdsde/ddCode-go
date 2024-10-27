package dto

import (
	"ddCode-server/global"
	"ddCode-server/internal"
	"ddCode-server/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type UserLoginParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UserLoginParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(u)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(u); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, u)
		return err
	}
	return nil
}

type UserListParam struct {
	UserName string `json:"user_name" form:"userName"`
	NickName string `json:"nick_name" form:"nickName"`
	internal.Pagination
}

func (u *UserListParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(u)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(u); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, u)
		return err
	}
	return nil
}

type UserDeleteParam struct {
	Id int `json:"id" form:"id" validate:"required"`
}

func (u *UserDeleteParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(u)
	if err != nil {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", err.Error(), u)
		return err
	}
	if msg, errCode := utils.Validate(u); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, u)
		return err
	}
	return nil
}

type UserListDto struct {
	Id       int    `json:"id" gorm:"column:id;primary_key"`
	UserName string `json:"userName" gorm:"column:user_name;unique"`
	NickName string `json:"nickName" gorm:"column:nick_name"`
	Email    string `json:"email" gorm:"column:email"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Sex      int    `json:"sex" gorm:"column:sex"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
	Comment  string `json:"comment" gorm:"column:comment"`
}

// UserRegisterParam 用户注册用
type UserRegisterParam struct {
	Username        string   `json:"username" validate:"required"`
	Nickname        string   `json:"nickname"`
	Password        string   `json:"password" validate:"required"`
	ConFirmPassword string   `json:"confirmPassword"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	Sex             int      `json:"sex"`
	Avatar          string   `json:"avatar"`
	DeptId          int      `json:"deptId"`
	Roles           []string `json:"roles"`
	Comment         string   `json:"comment"`
}

func (p *UserRegisterParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(p)
	if err != nil {
		return err
	}
	if p.Password != p.ConFirmPassword {
		return fmt.Errorf("两次输入密码不一致")
	}
	if msg, errCode := utils.Validate(p); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, p)
		return err
	}
	return nil
}

// UserInfoUpdateParam 用户自己更新信息
type UserInfoUpdateParam struct {
	Id        int       `json:"id" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password" validate:"required"`
	DeptId    int       `json:"deptId"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Sex       int       `json:"sex"`
	Avatar    string    `json:"avatar"`
	Comment   string    `json:"comment"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (p *UserInfoUpdateParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(p)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(p); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, p)
		return err
	}
	return nil
}
