package service

import (
	"ddCode-server/dao"
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type MenuService struct {
	Service
}

func (s *MenuService) GetMenuList(p *dto.MenuSearchParam) (list []*models.Menu, count int64, err error) {
	d := new(dao.MenuDao)
	list, count, err = d.GetMenuList(p)
	if err != nil {
		global.Logger.Errorf("GetMenuList err:%v", err)
		return
	}
	return list, count, nil
}
func (s *MenuService) AddMenu(p *models.Menu) (err error) {
	d := new(dao.MenuDao)
	err = d.CreateMenu(p)
	if err != nil {
		global.Logger.Errorf("AddMenu err:%v", err)
		return
	}
	return err
}
func (s *MenuService) UpdateMenu(p *models.Menu) (err error) {
	d := new(dao.MenuDao)
	err = d.UpdateMenuById(p)
	if err != nil {
		global.Logger.Errorf("UpdateMenuById err:%v", err)
		return
	}
	return err
}
func (s *MenuService) DeleteMenu(ids []int) (err error) {
	d := new(dao.MenuDao)
	err = d.RemoveMenuById(ids)
	if err != nil {
		global.Logger.Errorf("DeleteMenu err:%v", err)
		return
	}
	return err
}

func (s *MenuService) GetAllParentMenu() (list []models.Menu, err error) {
	d := new(dao.MenuDao)
	menus, err := d.GetAllParentMenu()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

//func (s *MenuService) GetMenuByRoleId(roleId int) (menu []models.Menu, err error) {
//	d := new(dao.MenuDao)
//	menuDao := new(dao.MenuDao)
//	roleMenuDao := new(dao.RoleMenuDao)
//	roleMenu, err := roleMenuDao.FindAllRoleMenuByRoleId(roleId)
//	if err != nil {
//		return nil, err
//	}
//	roleIds := make([]models.RoleMenu, 0, len(roleMenu))
//}
