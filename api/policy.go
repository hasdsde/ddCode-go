package api

import (
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
	"ddCode-server/pkg/resp"
	"ddCode-server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Policy struct{}

// GetPolicyList 分页查询
func (m *Policy) GetPolicyList(c *gin.Context) {
	p := new(dto.PolicySearchParam)
	s := new(service.PolicyService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	list, count, err := s.GetPolicyList(p)
	if err != nil {
		global.Logger.Errorf("查询权限信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.PageOK(c, list, int(count), p.PageNo, p.PageSize, "ok")
	return
}

// CreatePolicy 创建权限
func (m *Policy) CreatePolicy(c *gin.Context) {
	p := new(dto.PolicyCreateParam)
	s := new(service.PolicyService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.UpdatePolicy(&models.Policy{
		Url:         p.Url,
		Method:      p.Method,
		Description: p.Description,
	})
	if err != nil {
		global.Logger.Errorf("创建权限信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// UpdatePolicy 更新权限
func (m *Policy) UpdatePolicy(c *gin.Context) {
	p := new(dto.PolicyUpdateParam)
	s := new(service.PolicyService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.UpdatePolicy(&models.Policy{
		Id:          p.Id,
		Url:         p.Url,
		Method:      p.Method,
		Description: p.Description,
	})
	if err != nil {
		global.Logger.Errorf("创建权限信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// DeletePolicy 删除权限
func (m *Policy) DeletePolicy(c *gin.Context) {
	p := new(dto.PolicyDeleteParam)
	s := new(service.PolicyService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.DeletePolicy(p.Ids)
	if err != nil {
		global.Logger.Errorf("删除权限失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}
