package models

type Policy struct {
	Id          int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Url         string `json:"url" gorm:"column:url;type:varchar(255);not null"`
	Method      string `json:"method" gorm:"column:method;type:varchar(255);not null"`
	Description string `json:"description" gorm:"column:description;type:varchar(255);not null"`
}

func (p *Policy) TableName() string {
	return "policy"
}
