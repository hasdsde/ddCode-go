package utils

import "reflect"

var ignoreField = map[string]struct{}{
	"fields":       {},
	"norms":        {},
	"type":         {},
	"ignore_above": {},
	"raw":          {},
	"format":       {},
	"keyword":      {},
}

// GetMappingTree 获取 mapping 树结构
func GetMappingTree(mapping map[string]interface{}) map[string]interface{} {
	return doGetMappingTree(mapping, make(map[string]interface{}))
}

// doGetMappingTree
func doGetMappingTree(mapping, target map[string]interface{}) map[string]interface{} {
	if target == nil {
		target = make(map[string]interface{})
	}
	// 如果没有 mapping
	if _, ok := mapping["mappings"]; ok {
		mapping = mapping["mappings"].(map[string]interface{})
	}
	for k, m := range mapping {
		// 跳过不需要的字段
		if _, ok := ignoreField[k]; ok {
			continue
		}
		mType := reflect.TypeOf(m)

		// mValue := reflect.ValueOf(m)
		if k == "location" {
			target[k] = map[string]interface{}{
				"lat": struct{}{},
				"lon": struct{}{},
			}
		} else if mType.Kind() == reflect.Map {
			m1 := m.(map[string]interface{})
			var subTarget = make(map[string]interface{})
			subTarget = doGetMappingTree(m1, subTarget)
			target[k] = subTarget
		} else {
			target[k] = struct{}{}
		}
	}
	if _, ok := target["properties"]; ok {
		return target["properties"].(map[string]interface{})
	}
	return target
}
