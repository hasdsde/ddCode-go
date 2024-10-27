package api

import (
	"ddCode-server/global"
	"ddCode-server/models/dto"
	"ddCode-server/pkg/resp"
	"ddCode-server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Log struct{}

// GetLogList 分页查询
func (m *Log) GetLogList(c *gin.Context) {
	p := new(dto.LogSearchParam)
	s := new(service.LogService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	list, count, err := s.GetLogList(p)
	if err != nil {
		global.Logger.Errorf("查询日志信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.PageOK(c, list, int(count), p.PageNo, p.PageSize, "ok")
	return
}
