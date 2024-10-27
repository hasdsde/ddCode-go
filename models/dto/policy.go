package dto

import (
	"ddCode-server/global"
	"ddCode-server/internal"
	"ddCode-server/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PolicySearchParam struct {
	Url    string `form:"url" json:"url"`
	Method string `form:"method" json:"method"`
	internal.Pagination
}

func (d *PolicySearchParam) Bind(c *gin.Context) error {
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

type PolicyCreateParam struct {
	Url         string `form:"url" json:"url"`
	Method      string `form:"method" json:"method"`
	Description string `form:"description" json:"description"`
}

func (d *PolicyCreateParam) Bind(c *gin.Context) error {
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

type PolicyUpdateParam struct {
	Id          int    `json:"id" validate:"required"`
	Url         string `form:"url" json:"url"`
	Method      string `form:"method" json:"method"`
	Description string `form:"description" json:"description"`
}

func (d *PolicyUpdateParam) Bind(c *gin.Context) error {
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

type PolicyDeleteParam struct {
	Ids []int `json:"ids" validate:"required"`
}

func (d *PolicyDeleteParam) Bind(c *gin.Context) error {
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
