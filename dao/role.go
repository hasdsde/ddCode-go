package dao

import (
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type RoleDao struct {
	Dao
}

var tableRole models.Role

func (d *RoleDao) GetRoleList(param *dto.RoleSearchParam) (list []*models.Role, count int64, err error) {
	list = make([]*models.Role, 0)
	orm := global.Db.DB().Table(tableRole.TableName())
	if param.Name != "" {
		orm = orm.Where("name like ?", "%"+param.Name+"%")
	}
	err = orm.
		Count(&count).
		Limit(param.PageLimit()).
		Offset(param.PageOffset()).
		Find(&list).Error
	return
}
func (d *RoleDao) GetRoleById(id int64) (role *models.Role, err error) {
	role = new(models.Role)
	err = global.Db.DB().Table(tableRole.TableName()).Where("id = ?", id).First(role).Error
	return role, err
}
func (d *RoleDao) CreateRole(Role *models.Role) (err error) {
	err = global.Db.DB().Save(Role).Error
	return
}

func (d *RoleDao) UpdateRole(Role *models.Role) (err error) {
	err = global.Db.DB().Save(Role).Error
	return
}

func (d *RoleDao) DeleteRoleById(ids []int) (err error) {
	err = global.Db.DB().Where("id in (?)", ids).Delete(&models.Role{}).Error
	return
}

func (d *RoleDao) FindRolesByRoleNames(names []string) ([]*models.Role, error) {
	roles := make([]*models.Role, 0)
	orm := global.Db.DB().Table(tableRole.TableName())
	err := orm.Where("name in (?)", names).Find(&roles).Error
	return roles, err
}
