package utils

import "sort"

// SliceDeduplication 字符串切片分组isSort 否表示是排序过的切片，否则需要排序，双指针方式去重，没有额外的内存开销
func SliceDeduplication(isSort bool, slice []string) []string {
	if len(slice) < 2 {
		return slice
	}
	if isSort {
		sort.Strings(slice)
	}
	index := 0
	for i := 0; i < len(slice); i++ {
		if slice[index] != slice[i] {
			index++
			slice[index] = slice[i]
		}
	}
	return slice[:index+1]
}
