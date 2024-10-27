package service

import (
	"ddCode-server/dao"
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
	"ddCode-server/models/vo"
	"fmt"
)

type RoleService struct {
	Service
}

// GetRoleList 获取角色列表
func (s *RoleService) GetRoleList(p *dto.RoleSearchParam) (list []*models.Role, count int64, err error) {
	d := new(dao.RoleDao)
	list, count, err = d.GetRoleList(p)
	if err != nil {
		global.Logger.Errorf("GetRoleList err:%v", err)
		return
	}
	return list, count, nil
}
func (s *RoleService) AddRole(p *models.Role) (err error) {
	d := new(dao.RoleDao)
	err = d.CreateRole(p)
	if err != nil {
		global.Logger.Errorf("AddRole err:%v", err)
		return
	}
	return err
}

func (s *RoleService) CreateUser(p *models.Role) (err error) {
	d := new(dao.RoleDao)
	err = d.CreateRole(p)
	if err != nil {
		global.Logger.Errorf("UpdateRole err:%v", err)
		return
	}
	return err
}

func (s *RoleService) UpdateRole(p *models.Role) (err error) {
	d := new(dao.RoleDao)
	oldRole, err := d.GetRoleById(int64(p.Id))
	if err != nil {
		return err
	}
	err = d.UpdateRole(p)
	if err != nil {
		global.Logger.Errorf("UpdateRole err:%v", err)
		return
	}
	if oldRole.Name != p.Name {
		err := global.Rbac.UpdateRole(oldRole.Name, p.Name)
		if err != nil {
			global.Logger.Errorf("UpdateRole err:%v", err)
			return err
		}
	}
	return err
}
func (s *RoleService) DeleteRole(ids []int) (err error) {
	d := new(dao.RoleDao)
	err = d.DeleteRoleById(ids)
	if err != nil {
		global.Logger.Errorf("DeleteRole err:%v", err)
		return
	}
	return err
}

// UpdateRolePolicy 更新角色权限
func (s *RoleService) UpdateRolePolicy(p *dto.RolePolicyUpdateParam) error {
	err := global.Rbac.DeleteAllRolePolicies(p.RoleName)
	if err != nil {
		return fmt.Errorf("更新权限失败")
	}
	rolePolicy := vo.RolePolicyToSlice(p.RoleName, p.Policy)
	err = global.Rbac.CreateRolePolicies(rolePolicy)
	if err != nil {
		return fmt.Errorf("更新权限失败")
	}
	return nil
}
