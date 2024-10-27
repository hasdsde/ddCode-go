package models

type RoleMenu struct {
	RoleId int `json:"roleId" gorm:"column:role_id"`
	MenuId int `json:"menuId" gorm:"column:menu_id"`
}

func (r *RoleMenu) TableName() string {
	return "role_menu"
}
