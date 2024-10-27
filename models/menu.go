package models

type Menu struct {
	Id       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Url      string `json:"url" gorm:"column:url;type:varchar(255)"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255)"`
	ParentId int    `json:"parentId" gorm:"column:parent_id;type:int(11)"`
	Orders   int    `json:"orders" gorm:"column:orders;type:int(11)"`
	Icon     string `json:"icon" gorm:"column:icon;type:varchar(255)"`
}

func (m *Menu) TableName() string {
	return "menu"
}
