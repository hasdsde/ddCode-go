package dto

import (
	"ddCode-server/global"
	"ddCode-server/internal"
	"ddCode-server/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MenuSearchParam struct {
	Name     string `form:"name" json:"name"`
	ParentId string `form:"parentId" json:"parentId"`
	internal.Pagination
}

func (d *MenuSearchParam) Bind(c *gin.Context) error {
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

type MenuCreateParam struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	ParentId int    `json:"parentId"`
	Orders   int    `json:"orders"`
	Icon     string `json:"icon"`
}

func (d *MenuCreateParam) Bind(c *gin.Context) error {
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

type MenuUpdateParam struct {
	Id       int    `json:"id" validate:"required"`
	Url      string `json:"url"`
	Name     string `json:"name"`
	ParentId int    `json:"parentId"`
	Orders   int    `json:"orders"`
	Icon     string `json:"icon"`
}

func (d *MenuUpdateParam) Bind(c *gin.Context) error {
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

type MenuDeleteParam struct {
	Ids []int `json:"ids" validate:"required"`
}

func (d *MenuDeleteParam) Bind(c *gin.Context) error {
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
