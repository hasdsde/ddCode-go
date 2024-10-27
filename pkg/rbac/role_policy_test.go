package rbac

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCasbinService_CreateRolePolicy(t *testing.T) {
	casbinService, err := InitTest()
	assert.Nil(t, err)
	err = casbinService.CreateRolePolicy("admin", "/towers", "GET")
	assert.Equal(t, nil, err)
}
func TestCasbinService_CreateUserRolesBatch(t *testing.T) {
	casbinService, err := InitTest()
	assert.Equal(t, nil, err)
	err = casbinService.CreateRolePolicies([][]string{{"test1", "info1", "get"}, {"test1", "info2", "GET"}})
	assert.Equal(t, nil, err)
}
func TestCasbinService_GetPolicyFromUser(t *testing.T) {
	casbinService, err := InitTest()
	assert.Equal(t, nil, err)
	userProlity, err := casbinService.GetPolicyByRoleName("admin")
	assert.Equal(t, nil, err)
	fmt.Println(userProlity)
}
func TestCasbinService_UpdateRolePolicy(t *testing.T) {
	casbinService, err := InitTest()
	assert.Equal(t, nil, err)
	err = casbinService.UpdateRolePolicy([]string{"admin", "/towers", "GET"}, []string{"admin", "/towers", "DELETE"})
	assert.Equal(t, nil, err)
}
func TestCasbinService_UpdateRolePolicies(t *testing.T) {
	casbinService, err := InitTest()
	assert.Equal(t, nil, err)
	err = casbinService.UpdateRolePolicies([][]string{{"admin", "/towers", "GET"}}, [][]string{{"admin", "/towers", "DELETE"}})
	assert.Equal(t, nil, err)
}

func TestCasbinService_DeleteRolePolicy(t *testing.T) {
	casbinService, err := InitTest()
	assert.Equal(t, nil, err)
	err = casbinService.DeleteRolePolicy("admin", "tower", "GET")
	assert.Equal(t, nil, err)
}

func TestCasbinService_DeleteRolePolicies(t *testing.T) {
	casbinService, err := InitTest()
	assert.Equal(t, nil, err)
	err = casbinService.DeleteRolePolicies([][]string{{"admin", "/towers", "GET"}})
	assert.Equal(t, nil, err)
}

func TestCasbinService_HasAccess(t *testing.T) {
	casbinService, err := InitTest()
	assert.Equal(t, nil, err)
	access, err := casbinService.HasAccess("zhangsans", "/towers", "GET")
	assert.Equal(t, nil, err)
	fmt.Println(access)
}

func TestCasbinService_DeleteAllRolePolicies(t *testing.T) {
	casbinService, err := InitTest()
	assert.Equal(t, nil, err)
	err = casbinService.DeleteAllRolePolicies("lisi")
	assert.Equal(t, nil, err)
}
