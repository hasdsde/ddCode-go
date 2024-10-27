package internal

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type ModelBModel struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type BaseModel struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type ControlBy struct {
	CreatedAt time.Time `json:"createBy" gorm:"index;default:1;comment:创建者"`
	UpdatedAt time.Time `json:"updateBy" gorm:"index;default:1;comment:更新者"`
	DeletedAt time.Time `json:"deletedAt"`
}

func (e *ControlBy) SetCreatedAt(createBy time.Time) {
	e.CreatedAt = createBy
}

func (e *ControlBy) SetUpdatedAt(updateBy time.Time) {
	e.UpdatedAt = updateBy
}
func (e *ControlBy) SetDeletedAt(deleteBy time.Time) {
	e.DeletedAt = deleteBy
}

type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	timeStr := fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))
	if timeStr == fmt.Sprintf("\"%v\"", "0001-01-01 00:00:00") {
		timeStr = "\"\""
	}
	return []byte(timeStr), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	// LocalTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}
func (t LocalTime) StrValue() string {
	// LocalTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05")
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	// 前端接收的时间字符串
	str := string(data)
	// 去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime(t1)
	return err
}
