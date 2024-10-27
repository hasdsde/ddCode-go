package rbac

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCasbinService_AddUserRole(t *testing.T) {
	s, err := InitTest()
	assert.Equal(t, err, nil)
	err = s.CreateUserRole("zhangsan", "test1")
	assert.Equal(t, err, nil)
}

func TestCasbinService_GetAllRoleUser(t *testing.T) {
	s, err := InitTest()
	assert.Equal(t, err, nil)
	roles, err := s.GetAllUserRole()
	fmt.Println(roles)
	assert.Equal(t, err, nil)
}

func TestCasbinService_AddUserRolesBatch(t *testing.T) {
	s, err := InitTest()
	if err != nil {
		fmt.Println(err)
	}
	err = s.CreateUserRolesBatch([][]string{{"zhangsan", "admin222"}, {"lisi", "user222"}})
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestCasbinService_UpdateUserRole(t *testing.T) {
	s, err := InitTest()
	if err != nil {
		fmt.Println(err)
	}
	err = s.UpdateUserRole([]string{"zhangsan", "admin222"}, []string{"zhangsan", "admin333"})
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestCasbinService_UpdateUserRolesBatch(t *testing.T) {
	s, err := InitTest()
	if err != nil {
		fmt.Println(err)
	}
	err = s.UpdateUserRolesBatch([][]string{{"lisi", "user2"}, {"lisi", "user3"}}, [][]string{{"lisi", "user4"}, {"lisi", "user5"}})
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestCasbinService_DeleteUserRole(t *testing.T) {
	s, err := InitTest()
	if err != nil {
		fmt.Println(err)
	}
	err = s.DeleteUserRole([]string{"zhangsan", "admin"})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestCasbinService_DeleteUserRolesBatch(t *testing.T) {
	s, err := InitTest()
	if err != nil {
		fmt.Println(err)
	}
	err = s.RemoveUserRoleBatch([][]string{{"guest", "admin"}, {"lisi", "user2"}})
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestCasbinService_GetPolicyByRoleName(t *testing.T) {
	s, err := InitTest()
	assert.Equal(t, err, nil)
	roles, err := s.GetRoleByUserName("admin")
	assert.Equal(t, err, nil)
	fmt.Println(roles)
}
