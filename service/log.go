package service

import (
	"ddCode-server/dao"
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type LogService struct {
	Service
}

func (s *LogService) GetLogList(p *dto.LogSearchParam) (list []*models.Log, count int64, err error) {
	d := new(dao.LogDao)
	list, count, err = d.GetLogList(p)
	if err != nil {
		global.Logger.Errorf("GetLogList err:%v", err)
		return
	}
	return list, count, nil
}

func (s *LogService) AddLog(p *models.Log) (err error) {
	d := new(dao.LogDao)
	err = d.CreateLog(p)
	if err != nil {
		global.Logger.Errorf("AddLog err:%v", err)
		return
	}
	return err
}
