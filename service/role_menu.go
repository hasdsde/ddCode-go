package service

import (
	"ddCode-server/dao"
	"ddCode-server/global"
	"ddCode-server/models"
)

type RoleMenuService struct {
	Service
}

// UpdateRoleMenu 更新ROleMenus
func (s *RoleMenuService) UpdateRoleMenu(roleId int, menuIds []int) error {
	roleMenuDao := &dao.RoleMenuDao{}
	err := roleMenuDao.DeleteRoleMenuByRoleId(roleId)
	if err != nil {
		global.Logger.Errorf("根据id删除roleMenu失败:%s", err.Error())
		return err
	}
	roleMenus := make([]models.RoleMenu, 0, len(menuIds))
	for _, menuId := range menuIds {
		roleMenus = append(roleMenus, models.RoleMenu{
			RoleId: roleId,
			MenuId: menuId,
		})
	}
	err = roleMenuDao.CreateRoleMenuByRoleId(roleMenus)
	if err != nil {
		global.Logger.Errorf("新增roleMenu失败:%s", err.Error())
		return err
	}
	return nil
}

// FindMenusByRoleIds  根据roleIds获取Menus
func (s *RoleMenuService) FindMenusByRoleIds(roleIds []int) ([]models.Menu, error) {
	roleMenuDao := &dao.RoleMenuDao{}
	menuDao := &dao.MenuDao{}
	menuRoles, err := roleMenuDao.FindAllRoleMenuByRoleId(roleIds)
	menuIds := make([]int, 0, len(menuRoles))
	if err != nil {
		return nil, err
	}
	for _, v := range menuRoles {
		menuIds = append(menuIds, v.MenuId)
	}
	menus, err := menuDao.FindMenusByIds(menuIds)
	if err != nil {
		return nil, err
	}
	return menus, nil
}
func (s *RoleMenuService) FindRoleMenuByRoleNames(roleNames []string) ([]models.Menu, error) {
	roleDao := &dao.RoleDao{}
	roles, err := roleDao.FindRolesByRoleNames(roleNames)
	roleIds := make([]int, 0, len(roles))
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}
	if err != nil {
		return nil, err
	}
	menus, err := s.FindMenusByRoleIds(roleIds)
	if err != nil {
		return nil, err
	}
	return menus, nil
}
