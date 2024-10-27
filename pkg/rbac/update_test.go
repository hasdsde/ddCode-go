package rbac

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCasbinService_GetRoles(t *testing.T) {
	casbinService, err := InitTest()
	if err != nil {
		fmt.Println(err)
		return
	}
	roles, err := casbinService.GetRoles()
	fmt.Println(roles)
}

func TestCasbinService_UpdateRole(t *testing.T) {
	casbinService, err := InitTest()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = casbinService.UpdateRole("test1", "test2")
	assert.Equal(t, nil, err)
}

func TestCasbinService_UpdateRoleNameUser(t *testing.T) {
	casbinService, err := InitTest()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = casbinService.UpdateRoleForUsers("test1", "test2")
	assert.Equal(t, nil, err)
}

func TestCasbinService_UpdateRoleNamePolicy(t *testing.T) {
	casbinService, err := InitTest()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = casbinService.UpdateRoleForPolicies("test2", "test3")
	assert.Equal(t, nil, err)
}

func TestCasbinService_UpdateRoleName(t *testing.T) {
	casbinService, err := InitTest()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = casbinService.UpdateRole("test1", "test2")
	assert.Equal(t, nil, err)
}

func TestCasbinService_UpdatePolicy(t *testing.T) {
	casbinService, err := InitTest()
	oldPolicy := []string{"info1", "GET"}
	newPolicy := []string{"info1", "GETTT"}
	err = casbinService.UpdatePolicy(oldPolicy, newPolicy)
	if err != nil {
		return
	}
}

func TestCasbinService_UpdateUser(t *testing.T) {
	casbinService, err := InitTest()
	if err != nil {
		fmt.Println(err)
	}
	err = casbinService.UpdateUser("lisi", "wangwu")
	if err != nil {
		return
	}
}
