package models

type Dept struct {
	Id          int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name" gorm:"column:name;unique;not null"`
	ParentId    int    `json:"parentId" gorm:"column:parent_id"`
	Description string `json:"description" gorm:"column:description"`
}

func (d *Dept) TableName() string {
	return "dept"
}
