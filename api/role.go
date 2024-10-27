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

type Role struct{}

// GetRoleList 分页查询
func (m *Role) GetRoleList(c *gin.Context) {
	p := new(dto.RoleSearchParam)
	s := new(service.RoleService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	list, count, err := s.GetRoleList(p)
	if err != nil {
		global.Logger.Errorf("查询角色信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.PageOK(c, list, int(count), p.PageNo, p.PageSize, "ok")
	return
}

// CreateRole 创建角色
func (m *Role) CreateRole(c *gin.Context) {
	p := new(dto.RoleCreateParam)
	s := new(service.RoleService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.CreateUser(&models.Role{
		Name:        p.Name,
		Description: p.Description,
	})
	if err != nil {
		global.Logger.Errorf("创建角色信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// UpdateRole 更新角色
func (m *Role) UpdateRole(c *gin.Context) {
	p := new(dto.RoleUpdateParam)
	s := new(service.RoleService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.UpdateRole(&models.Role{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
	})
	if err != nil {
		global.Logger.Errorf("创建角色信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// DeleteRole 删除角色
func (m *Role) DeleteRole(c *gin.Context) {
	p := new(dto.RoleDeleteParam)
	s := new(service.RoleService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.DeleteRole(p.Ids)
	if err != nil {
		global.Logger.Errorf("删除角色失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// UpdateRoleMenu 为角色添加菜单
func (m *Role) UpdateRoleMenu(c *gin.Context) {
	p := new(dto.RoleMenuUpdateParam)
	s := new(service.RoleMenuService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.UpdateRoleMenu(p.RoleId, p.MenuIds)
	if err != nil {
		global.Logger.Errorf("更新菜单失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// FindRoleMenu 获取角色菜单
func (m *Role) FindRoleMenu(c *gin.Context) {
	p := new(dto.RoleMenuSearchParam)
	s := new(service.RoleMenuService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	list, err := s.FindMenusByRoleIds(p.RoleId)
	if err != nil {
		global.Logger.Errorf("获取角色对应菜单菜单失败, parms: %v", p)
		resp.Error(c, http.StatusBadRequest, err, "获取角色对应菜单菜单失败")
		return
	}
	resp.OKList(c, list, "ok")
}

// CreateRoleUser 为用户添加角色
func (m *Role) CreateRoleUser(c *gin.Context) {
	p := new(dto.RoleUserCreateParam)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = global.Rbac.CreateUserRole(p.UserName, p.RoleName)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "添加权限失败")
		return
	}
	resp.OK(c, "", "ok")
}

// DeleteRoleUser 为用户删除角色
func (m *Role) DeleteRoleUser(c *gin.Context) {
	p := new(dto.RoleUserRemoveParam)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = global.Rbac.DeleteUserRole([]string{p.UserName, p.RoleName})
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "删除权限失败")
		return
	}
	resp.OK(c, "", "ok")
}

// FindRoleUser 根据用户名获取角色
func (m *Role) FindRoleUser(c *gin.Context) {
	p := new(dto.RoleUserSearchParam)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	roles, err := global.Rbac.GetRoleByUserName(p.UserName)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "获取用户角色失败")
		return
	}
	resp.OKList(c, roles, "ok")
}

func (m *Role) UpdateRolePolicy(c *gin.Context) {
	p := new(dto.RolePolicyUpdateParam)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	s := new(service.RoleService)
	err = s.UpdateRolePolicy(p)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "更新角色权限失败")
		return
	}
	resp.OK(c, "", "ok")
}

// FindRolePolicy 获取角色权限
func (m *Role) FindRolePolicy(c *gin.Context) {
	p := new(dto.RolePolicySearchParam)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	policy, err := global.Rbac.GetPolicyByRoleName(p.RoleName)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "获取角色权限失败")
		return
	}
	resp.OKList(c, policy, "ok")
}
