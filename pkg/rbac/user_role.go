package rbac

// GetAllUserRole  获取用户和角色表
func (c *CasbinService) GetAllUserRole() ([][]string, error) {
	policy, err := c.e.GetGroupingPolicy()
	return policy, err
}

// GetRoleByUserName 根据用户名获取角色
func (c *CasbinService) GetRoleByUserName(userName string) ([]string, error) {
	roles, err := c.e.GetFilteredGroupingPolicy(0, userName)
	names := make([]string, 0, len(roles))
	for _, v := range roles {
		names = append(names, v[1])
	}
	return names, err
}

// CreateUserRole 为用户赋予角色
func (c *CasbinService) CreateUserRole(userName, roleName string) error {
	_, err := c.e.AddGroupingPolicy(userName, roleName)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// CreateUserRolesBatch 批量为用户赋予角色
func (c *CasbinService) CreateUserRolesBatch(params [][]string) error {
	_, err := c.e.AddGroupingPolicies(params)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// UpdateUserRole 更新角色用户
func (c *CasbinService) UpdateUserRole(oldRule, newRule []string) error {
	_, err := c.e.UpdateGroupingPolicy(oldRule, newRule)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// UpdateUserRolesBatch 更新角色用户
func (c *CasbinService) UpdateUserRolesBatch(oldRule, newRule [][]string) error {
	_, err := c.e.UpdateGroupingPolicies(oldRule, newRule)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// DeleteUserRole 将用户从角色中删除
func (c *CasbinService) DeleteUserRole(params []string) error {
	_, err := c.e.RemoveGroupingPolicy(params)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// RemoveUserRoleBatch 批量将用户从角色中删除
func (c *CasbinService) RemoveUserRoleBatch(params [][]string) error {
	_, err := c.e.RemoveGroupingPolicies(params)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}
