package models

type Role struct {
	Id          int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name" gorm:"column:name;type:varchar(50);not null"`
	Description string `json:"description" gorm:"column:description;type:text;not null"`
}

func (r *Role) TableName() string {
	return "role"
}
