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

type Dept struct{}

// GetDeptList 分页查询
func (m *Dept) GetDeptList(c *gin.Context) {
	p := new(dto.DeptSearchParam)
	s := new(service.DeptService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	list, err := s.GetDeptList(p)
	if err != nil {
		global.Logger.Errorf("查询部门信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OKList(c, list, "ok")
	return
}

// CreateDept 创建部门
func (m *Dept) CreateDept(c *gin.Context) {
	p := new(dto.DeptCreateParam)
	s := new(service.DeptService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.AddDept(&models.Dept{
		Name:        p.Name,
		ParentId:    p.ParentId,
		Description: p.Description,
	})
	if err != nil {
		global.Logger.Errorf("创建部门信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// UpdateDept 更新部门
func (m *Dept) UpdateDept(c *gin.Context) {
	p := new(dto.DeptUpdateParam)
	s := new(service.DeptService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.UpdateDept(&models.Dept{
		Id:          p.Id,
		Name:        p.Name,
		ParentId:    p.ParentId,
		Description: p.Description,
	})
	if err != nil {
		global.Logger.Errorf("创建部门信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// DeleteDept 删除部门
func (m *Dept) DeleteDept(c *gin.Context) {
	p := new(dto.DeptDeleteParam)
	s := new(service.DeptService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.DeleteDept(p.Ids)
	if err != nil {
		global.Logger.Errorf("删除部门失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}
