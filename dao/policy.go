package dao

import (
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type PolicyDao struct {
	Dao
}

var tablePolicy models.Policy

func (d *PolicyDao) GetPolicyList(param *dto.PolicySearchParam) (list []*models.Policy, count int64, err error) {
	list = make([]*models.Policy, 0)
	orm := global.Db.DB().Table(tablePolicy.TableName())
	if param.Url != "" {
		orm = orm.Where("url like ?", "%"+param.Url+"%")
	}
	if param.Method != "" {
		orm = orm.Where("method = ?", param.Method)
	}
	err = orm.
		Count(&count).
		Limit(param.PageLimit()).
		Offset(param.PageOffset()).
		Find(&list).Error
	return
}
func (d *PolicyDao) GetPolicyById(id int) (policy *models.Policy, err error) {
	policy = new(models.Policy)
	err = global.Db.DB().Where("id=?", id).First(&policy).Error
	return
}

func (d *PolicyDao) CreatePolicy(Policy *models.Policy) (err error) {
	err = global.Db.DB().Save(Policy).Error
	return
}

func (d *PolicyDao) UpdatePolicy(Policy *models.Policy) (err error) {
	err = global.Db.DB().Save(Policy).Error
	return
}

func (d *PolicyDao) DeletePolicyById(ids []int) (err error) {
	err = global.Db.DB().Where("id in (?)", ids).Delete(&models.Policy{}).Error
	return
}
