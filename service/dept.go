package service

import (
	"ddCode-server/dao"
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type DeptService struct {
	Service
}

func (s *DeptService) GetDeptList(p *dto.DeptSearchParam) (list []models.Dept, err error) {
	d := new(dao.DeptDao)
	list, err = d.GetDeptAvailableList(p)
	if err != nil {
		global.Logger.Errorf("GetDeptAvailableList err:%v", err)
		return
	}
	return list, nil
}
func (s *DeptService) AddDept(p *models.Dept) (err error) {
	d := new(dao.DeptDao)
	err = d.CreateDept(p)
	if err != nil {
		global.Logger.Errorf("AddDept err:%v", err)
		return
	}
	return err
}
func (s *DeptService) UpdateDept(p *models.Dept) (err error) {
	d := new(dao.DeptDao)
	err = d.UpdateDept(p)
	if err != nil {
		global.Logger.Errorf("UpdateDept err:%v", err)
		return
	}
	return err
}
func (s *DeptService) DeleteDept(ids []int) (err error) {
	d := new(dao.DeptDao)
	err = d.DeleteDeptById(ids)
	if err != nil {
		global.Logger.Errorf("DeleteDept err:%v", err)
		return
	}
	return err
}
