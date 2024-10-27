package rbac

func (c *CasbinService) GetPolicyByRoleName(username string) ([][]string, error) {
	namedPolicy, err := c.e.GetFilteredPolicy(0, username)
	return namedPolicy, err
}

// CreateRolePolicy 创建角色组权限
func (c *CasbinService) CreateRolePolicy(policy ...interface{}) error {
	_, err := c.e.AddPolicy(policy...)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// CreateRolePolicies 批量创建角色组权限
func (c *CasbinService) CreateRolePolicies(policy [][]string) error {
	_, err := c.e.AddPolicies(policy)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// UpdateRolePolicy 更新单条角色组权限
func (c *CasbinService) UpdateRolePolicy(oldPolicy []string, newPolicy []string) error {
	_, err := c.e.UpdatePolicy(oldPolicy, newPolicy)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// UpdateRolePolicies 批量更新角色组权限
func (c *CasbinService) UpdateRolePolicies(oldPolicy [][]string, newPolicy [][]string) error {
	_, err := c.e.UpdatePolicies(oldPolicy, newPolicy)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// DeleteRolePolicy 删除角色权限
func (c *CasbinService) DeleteRolePolicy(policy ...string) error {
	_, err := c.e.RemovePolicy(policy)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// DeleteRolePolicies 批量删除角色权限
func (c *CasbinService) DeleteRolePolicies(rule [][]string) error {
	_, err := c.e.RemovePolicies(rule)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

func (c *CasbinService) DeleteAllRolePolicies(roleName string) error {
	_, err := c.e.DeletePermissionsForUser(roleName)
	if err != nil {
		return err
	}
	return c.e.SavePolicy()
}

// HasAccess  验证用户权限
func (c *CasbinService) HasAccess(p, url, method string) (ok bool, err error) {
	return c.e.Enforce(p, url, method)
}
