package dao

import (
	"ddCode-server/global"
	"ddCode-server/models"
)

type RoleMenuDao struct {
	Dao
}

var tableRoleMenu models.RoleMenu

func (d *RoleMenuDao) CreateRoleMenuByRoleId(roleMenus []models.RoleMenu) error {
	return global.Db.DB().Table(tableRoleMenu.TableName()).CreateInBatches(roleMenus, len(roleMenus)).Error
}
func (d *RoleMenuDao) DeleteRoleMenuByRoleId(roleId int) error {
	return global.Db.DB().Table(tableRoleMenu.TableName()).Where("role_id = ?", roleId).Delete(&tableRoleMenu).Error
}
func (d *RoleMenuDao) FindAllRoleMenuByRoleId(roleIds []int) ([]models.RoleMenu, error) {
	list := make([]models.RoleMenu, 0)
	err := global.Db.DB().Table(tableRoleMenu.TableName()).Where("role_id in ?", roleIds).Find(&list).Error
	return list, err
}
