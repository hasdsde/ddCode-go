package vo

import "ddCode-server/models"

type UserLoginInfo struct {
	UserName string        `json:"userName"`
	NickName string        `json:"nickName"`
	Email    string        `json:"email"`
	Phone    string        `json:"phone"`
	Sex      int           `json:"sex"`
	Avatar   string        `json:"avatar"`
	DeptName string        `json:"DeptName"`
	Roles    []string      `json:"roles"`
	Menus    []models.Menu `json:"menus"`
}

// UserLogin 登录信息
type UserLogin struct {
	Info  UserLoginInfo `json:"info"`
	Token string        `json:"token"`
}

// UserTokenInfo token中存的信息
type UserTokenInfo struct {
	UserName string   `json:"userName"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	DeptId   int      `json:"deptId"`
	Roles    []string `json:"roles"`
}
