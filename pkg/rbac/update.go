package rbac

// GetRoles 获取所有用户
func (c *CasbinService) GetRoles() ([]string, error) {
	return c.e.GetAllRoles()
}

// UpdateRole 更新角色名字
func (c *CasbinService) UpdateRole(oldName, newName string) error {
	err := c.UpdateRoleForPolicies(oldName, newName)
	if err != nil {
		return err
	}
	err = c.UpdateRoleForUsers(oldName, newName)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRoleForPolicies 更新更新角色-权限关系
func (c *CasbinService) UpdateRoleForPolicies(oldName, newName string) error {
	userPermissions, err := c.e.GetPermissionsForUser(oldName)
	for _, v := range userPermissions {
		_, err = c.e.RemovePolicy(v)
		if err != nil {
			return err
		}
		v[0] = newName
		_, err = c.e.AddPolicy(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateRoleForUsers 更新角色-用户关系
func (c *CasbinService) UpdateRoleForUsers(oldName, newName string) error {
	users, err := c.e.GetUsersForRole(oldName)
	for _, v := range users {
		_, err = c.e.DeleteRoleForUser(v, oldName)
		if err != nil {
			return err
		}
		_, err = c.e.AddRoleForUser(v, newName)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdatePolicy 更新权限url或method
func (c *CasbinService) UpdatePolicy(oldPolicy []string, newPolicy []string) error {
	oldPolicies, err := c.e.GetFilteredPolicy(1, oldPolicy...)
	if err != nil {
		return err
	}
	for _, v := range oldPolicies {
		_, err := c.e.DeletePermission(oldPolicy...)
		if err != nil {
			return err
		}
		err = c.CreateRolePolicy(v[0], newPolicy[0], newPolicy[1])
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *CasbinService) UpdateUser(oldName string, newName string) error {
	roles, err := c.GetRoleByUserName(oldName)
	if err != nil {
		return err
	}
	for _, v := range roles {
		err := c.DeleteUserRole([]string{oldName, v})
		if err != nil {
			return err
		}
		err = c.CreateUserRole(newName, v)
		if err != nil {
			return err
		}
	}
	return nil
}
