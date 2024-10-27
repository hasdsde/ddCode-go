package dao

import (
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type MenuDao struct {
	Dao
}

var tableMenu models.Menu

func (d *MenuDao) GetAllParentMenu() (list []models.Menu, err error) {
	orm := global.Db.DB().Where("parent_id = 0").Table(tableMenu.TableName())
	err = orm.Find(&list).Error
	return
}
func (d *MenuDao) FindMenusByIds(ids []int) ([]models.Menu, error) {
	menus := make([]models.Menu, 0)
	orm := global.Db.DB().Table(tableMenu.TableName())
	err := orm.Where("id in ?", ids).Find(&menus).Error
	return menus, err
}

func (d *MenuDao) GetMenuList(param *dto.MenuSearchParam) (list []*models.Menu, count int64, err error) {
	list = make([]*models.Menu, 0)
	orm := global.Db.DB().Table(tableMenu.TableName())
	if param.Name != "" {
		orm = orm.Where("name like ?", "%"+param.Name+"%")
	}
	if param.ParentId != "" {
		orm = orm.Where("parent_id = ?", param.ParentId)
	}
	err = orm.
		Count(&count).
		Order("orders desc").
		Limit(param.PageLimit()).
		Offset(param.PageOffset()).
		Find(&list).Error
	return
}

func (d *MenuDao) CreateMenu(menu *models.Menu) (err error) {
	err = global.Db.DB().Save(menu).Error
	return
}

func (d *MenuDao) UpdateMenuById(menu *models.Menu) (err error) {
	err = global.Db.DB().Save(menu).Error
	return
}

func (d *MenuDao) RemoveMenuById(ids []int) (err error) {
	err = global.Db.DB().Where("id in (?)", ids).Delete(&models.Menu{}).Error
	return
}
