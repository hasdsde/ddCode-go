package service

import (
	"ddCode-server/dao"
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
)

type PolicyService struct {
	Service
}

func (s *PolicyService) GetPolicyList(p *dto.PolicySearchParam) (list []*models.Policy, count int64, err error) {
	d := new(dao.PolicyDao)
	list, count, err = d.GetPolicyList(p)
	if err != nil {
		global.Logger.Errorf("GetPolicyList err:%v", err)
		return
	}
	return list, count, nil
}
func (s *PolicyService) AddPolicy(p *models.Policy) (err error) {
	d := new(dao.PolicyDao)
	err = d.CreatePolicy(p)
	if err != nil {
		global.Logger.Errorf("AddPolicy err:%v", err)
		return
	}
	return err
}
func (s *PolicyService) UpdatePolicy(p *models.Policy) (err error) {
	d := new(dao.PolicyDao)
	err = d.UpdatePolicy(p)
	if err != nil {
		global.Logger.Errorf("UpdateRoleForPolicies err:%v", err)
		return
	}
	oldPolicy, err := d.GetPolicyById(p.Id)
	if err != nil {
		return err
	}
	if oldPolicy.Url != p.Url || oldPolicy.Method != p.Method {
		err := global.Rbac.UpdatePolicy([]string{oldPolicy.Url, oldPolicy.Method}, []string{p.Url, p.Method})
		if err != nil {
			return err
		}
	}
	return err
}
func (s *PolicyService) DeletePolicy(ids []int) (err error) {
	d := new(dao.PolicyDao)
	err = d.DeletePolicyById(ids)
	if err != nil {
		global.Logger.Errorf("DeletePolicy err:%v", err)
		return
	}
	return err
}
