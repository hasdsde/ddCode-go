package dto

import (
	"ddCode-server/global"
	"ddCode-server/internal"
	"ddCode-server/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogSearchParam struct {
	internal.Pagination
}

func (d *LogSearchParam) Bind(c *gin.Context) error {
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
