package dto

import (
	"ddCode-server/global"
	"ddCode-server/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DeptSearchParam struct {
	ParentId string `form:"parentId" json:"parentId"`
}

func (d *DeptSearchParam) Bind(c *gin.Context) error {
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

type DeptCreateParam struct {
	Name        string `json:"name"`
	ParentId    int    `json:"parentId"`
	Description string `json:"description"`
}

func (d *DeptCreateParam) Bind(c *gin.Context) error {
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

type DeptUpdateParam struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name"`
	ParentId    int    `json:"parentId"`
	Description string `json:"description"`
}

func (d *DeptUpdateParam) Bind(c *gin.Context) error {
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

type DeptDeleteParam struct {
	Ids []int `json:"ids" validate:"required"`
}

func (d *DeptDeleteParam) Bind(c *gin.Context) error {
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
