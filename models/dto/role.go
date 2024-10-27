package dto

import (
	"ddCode-server/global"
	"ddCode-server/internal"
	"ddCode-server/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RoleSearchParam struct {
	Name string `json:"name" form:"name"`
	internal.Pagination
}

func (d *RoleSearchParam) Bind(c *gin.Context) error {
	err := c.ShouldBindQuery(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RoleCreateParam struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d *RoleCreateParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RoleUpdateParam struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d *RoleUpdateParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RoleDeleteParam struct {
	Ids []int `json:"ids" validate:"required"`
}

func (d *RoleDeleteParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RoleMenuUpdateParam struct {
	RoleId  int   `json:"roleId" validate:"required"`
	MenuIds []int `json:"menuIds"`
}

func (d *RoleMenuUpdateParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RoleMenuSearchParam struct {
	RoleId []int `json:"roleIds" form:"roleId" validate:"required"`
}

func (d *RoleMenuSearchParam) Bind(c *gin.Context) error {
	err := c.ShouldBindQuery(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RoleUserCreateParam struct {
	RoleName string `json:"roleName" validate:"required"`
	UserName string `json:"userName" validate:"required"`
}

func (d *RoleUserCreateParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RoleUserRemoveParam struct {
	RoleName string `json:"roleName" validate:"required"`
	UserName string `json:"userName" validate:"required"`
}
type RoleUserSearchParam struct {
	UserName string `json:"userName" form:"userName" validate:"required"`
}

func (d *RoleUserSearchParam) Bind(c *gin.Context) error {
	err := c.ShouldBindQuery(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}
func (d *RoleUserRemoveParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type PolicyAddParam struct {
	Url    string `json:"url" validate:"required"`
	Method string `json:"method" validate:"required"`
}
type RolePolicyCreateParam struct {
	RoleName string           `json:"roleName" validate:"required"`
	Policy   []PolicyAddParam `json:"policy" validate:"required"`
}

func (d *RolePolicyCreateParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RolePolicyUpdateParam struct {
	RoleName string           `json:"roleName" validate:"required"`
	Policy   []PolicyAddParam `json:"policy" validate:"required"`
}

func (d *RolePolicyUpdateParam) Bind(c *gin.Context) error {
	err := c.ShouldBind(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}

type RolePolicySearchParam struct {
	RoleName string `json:"roleName" form:"roleName" validate:"required"`
}

func (d *RolePolicySearchParam) Bind(c *gin.Context) error {
	err := c.ShouldBindQuery(d)
	if err != nil {
		return err
	}
	if msg, errCode := utils.Validate(d); errCode != 0 {
		global.Logger.Errorf("request params validate failed, err: %s, params: %v", msg, d)
		return fmt.Errorf("request params validate failed, err: %s, params: %v", msg, d)
	}
	return nil
}
