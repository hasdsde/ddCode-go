package dao

import (
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type LogDao struct {
	Dao
}

var tableLog models.Log

func (d *LogDao) GetLogList(param *dto.LogSearchParam) (list []*models.Log, count int64, err error) {
	list = make([]*models.Log, 0)
	orm := global.Db.DB().Table(tableLog.TableName())
	err = orm.
		Count(&count).
		Limit(param.PageLimit()).
		Offset(param.PageOffset()).
		Find(&list).Error
	return
}

func (d *LogDao) CreateLog(Log *models.Log) (err error) {
	err = global.Db.DB().Save(Log).Error
	return
}
