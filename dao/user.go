package dao

import (
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type UserDao struct {
	Dao
}

var tableUser models.User

func (d *UserDao) GetUserInfoById(id int) (*models.User, error) {
	var user models.User
	err := global.Db.DB().
		Where("id = ?", id).
		First(&user).
		Error
	return &user, err
}
func (d *UserDao) GetUserInfo(userName string) (*models.User, error) {
	var user models.User
	err := global.Db.DB().
		Where("user_name = ?", userName).
		Where("deleted_at IS NULL").
		First(&user).
		Error
	return &user, err
}

func (d *UserDao) CreateUser(m *models.User) error {
	return global.Db.DB().
		Create(m).Error
}

func (d *UserDao) UpdateUserInfo(m *models.User) error {
	return global.Db.DB().Table(tableUser.TableName()).
		Where("id = ?", m.Id).
		Updates(m).
		Error
}

func (d *UserDao) UpdateUser(m *models.User) error {
	return global.Db.DB().Save(m).Error
}
func (d *UserDao) FindUserByParamAndPage(p *dto.UserListParam) ([]models.User, int64, error) {
	list := make([]models.User, 0)
	var count int64
	tx := global.Db.DB().Table(tableUser.TableName()).Where("deleted_at IS NULL")
	if p.UserName != "" {
		tx = tx.Where("user_name like ?", p.UserName)
	}
	if p.NickName != "" {
		tx = tx.Where("nick_name like ?", p.NickName)
	}
	err := tx.Count(&count).Offset(p.PageOffset()).Limit(p.PageLimit()).Find(&list).Error
	return list, count, err
}
