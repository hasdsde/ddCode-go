package models

import "time"

type Log struct {
	Id       int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserName string    `json:"UserName" gorm:"column:user_name;type:varchar(50);not null"`
	UserId   int       `json:"UserId" gorm:"column:user_id;type:int(11);not null"`
	Info     string    `json:"info" gorm:"column:info;type:text;not null"`
	CreateAt time.Time `json:"createAt" gorm:"column:created_at;type:int(11);not null"`
	Method   string    `json:"method" gorm:"column:method;type:varchar(50);not null"`
}

func (l *Log) TableName() string {
	return "log"
}
