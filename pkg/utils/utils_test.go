package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitArray(t *testing.T) {
	arr1 := []int{}
	//数组0~24
	for i := 0; i < 25; i++ {
		arr1 = append(arr1, i)
	}
	res1 := SplitArray[int](arr1, 4)
	assert.Len(t, res1, 4)

	res2 := SplitArray[int](arr1, 7)
	assert.Len(t, res2, 7)

	res3 := SplitArray[int](arr1, 12)
	assert.Len(t, res3, 12)

	res4 := SplitArray[int](arr1, 30)
	assert.Len(t, res4, 25)
}

func TestSplitArrayOne(t *testing.T) {
	arr1 := []int{1}
	res1 := SplitArray[int](arr1, 2)
	assert.Len(t, res1, 1)

	res2 := SplitArray[int](arr1, 1)
	assert.Len(t, res2, 1)

	// 不能为0
	res3 := SplitArray[int](arr1, 0)
	assert.Len(t, res3, 1)

	res4 := SplitArray[int](arr1, 10)
	assert.Len(t, res4, 1)
}
