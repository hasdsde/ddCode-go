package dao

import (
	"ddCode-server/global"
	"ddCode-server/models/dto"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeptDao_GetAllDeptList(t *testing.T) {
	global.SetUpTest()
	dao := DeptDao{}
	param := &dto.DeptSearchParam{
		ParentId: "6",
	}
	list, err := dao.GetDeptAvailableList(param)
	for _, v := range list {
		fmt.Println(v)
	}
	assert.Equal(t, nil, err)
}
