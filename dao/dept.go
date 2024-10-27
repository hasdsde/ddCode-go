package dao

import (
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type DeptDao struct {
	Dao
}

var tableDept models.Dept

func (d *DeptDao) GetDeptNameById(id int) (string, error) {
	dept := models.Dept{}
	err := global.Db.DB().Where("id=?", id).Find(&dept).Error
	return dept.Name, err
}

// GetDeptAvailableList 查询父级id
func (d *DeptDao) GetDeptAvailableList(param *dto.DeptSearchParam) ([]models.Dept, error) {
	list := make([]models.Dept, 0)
	orm := global.Db.DB().Table(tableDept.TableName())
	err := orm.Where("parent_id = ?", param.ParentId).Scan(&list).Error
	return list, err
}

func (d *DeptDao) CreateDept(dept *models.Dept) (err error) {
	err = global.Db.DB().Save(dept).Error
	return
}

func (d *DeptDao) UpdateDept(dept *models.Dept) (err error) {
	err = global.Db.DB().Save(dept).Error
	return
}

func (d *DeptDao) DeleteDeptById(ids []int) (err error) {
	err = global.Db.DB().Where("id in (?)", ids).Delete(&models.Dept{}).Error
	return
}
