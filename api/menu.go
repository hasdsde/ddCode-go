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

type Menu struct{}

// GetMenuList 分页查询
func (m *Menu) GetMenuList(c *gin.Context) {
	p := new(dto.MenuSearchParam)
	s := new(service.MenuService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	list, count, err := s.GetMenuList(p)
	if err != nil {
		global.Logger.Errorf("查询菜单信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.PageOK(c, list, int(count), p.PageNo, p.PageSize, "ok")
	return
}

// CreateMenu 创建菜单
func (m *Menu) CreateMenu(c *gin.Context) {
	p := new(dto.MenuCreateParam)
	s := new(service.MenuService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.UpdateMenu(&models.Menu{
		Url:      p.Url,
		Name:     p.Name,
		ParentId: p.ParentId,
		Orders:   p.Orders,
		Icon:     p.Icon,
	})
	if err != nil {
		global.Logger.Errorf("创建菜单信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// UpdateMenu 更新菜单
func (m *Menu) UpdateMenu(c *gin.Context) {
	p := new(dto.MenuUpdateParam)
	s := new(service.MenuService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.UpdateMenu(&models.Menu{
		Id:       p.Id,
		Url:      p.Url,
		Name:     p.Name,
		ParentId: p.ParentId,
		Orders:   p.Orders,
		Icon:     p.Icon,
	})
	if err != nil {
		global.Logger.Errorf("创建菜单信息失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

// DeleteMenu 删除菜单
func (m *Menu) DeleteMenu(c *gin.Context) {
	p := new(dto.MenuDeleteParam)
	s := new(service.MenuService)
	err := p.Bind(c)
	if err != nil {
		resp.Error(c, http.StatusBadRequest, err, "参数验证失败")
		return
	}
	err = s.DeleteMenu(p.Ids)
	if err != nil {
		global.Logger.Errorf("删除菜单失败, parms: %v", p)
		resp.Error(c, http.StatusInternalServerError, err, "错误")
		return
	}
	resp.OK(c, nil, "操作成功")
}

func (m *Menu) GetAllParentMenu(c *gin.Context) {
	s := new(service.MenuService)
	menus, err := s.GetAllParentMenu()
	if err != nil {
		global.Logger.Errorf("查询所有菜单失败:%s", err.Error())
		return
	}
	resp.OKList(c, menus, "")
}
