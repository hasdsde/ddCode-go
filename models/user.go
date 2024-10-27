package models

import "time"

type User struct {
	Id        int        `json:"id" gorm:"column:id;primary_key"`
	UserName  string     `json:"userName" gorm:"column:user_name;unique"`
	Password  string     `json:"password" gorm:"column:password"`
	NickName  string     `json:"nickName" gorm:"column:nick_name"`
	Email     string     `json:"email" gorm:"column:email"`
	Phone     string     `json:"phone" gorm:"column:phone"`
	Sex       int        `json:"sex" gorm:"column:sex"`
	Avatar    string     `json:"avatar" gorm:"column:avatar"`
	DeptId    int        `json:"deptId" gorm:"column:dept_id"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
	Comment   string     `json:"comment" gorm:"column:comment"`
}

func (u *User) TableName() string {
	return "user"
}
