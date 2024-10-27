package utils

import (
	"reflect"
	"strings"
)

// AddKeyPrefixToMap 给 map 添加 key 前缀
// data    数据源 map
// target  目标   map
// prefix  要添加的前缀
// ignores 排除的字段
func AddKeyPrefixToMap(data, target map[string]interface{}, prefix string, ignores ...string) map[string]interface{} {
	passMap := make(map[string]struct{})
	for _, ig := range ignores {
		passMap[ig] = struct{}{}
	}
	for k, m := range data {
		if m == nil {
			continue
		}
		_, pass := passMap[k]
		mType := reflect.TypeOf(m)
		mValue := reflect.ValueOf(m)
		var d interface{}
		if mType.Kind() == reflect.Map {
			m1 := m.(map[string]interface{})
			var subTarget = make(map[string]interface{})
			subTarget = AddKeyPrefixToMap(m1, subTarget, prefix, ignores...)
			d = subTarget
		} else if mType.Kind() == reflect.Slice {
			l := mValue.Len()
			arr := make([]interface{}, 0, l)
			for i := 0; i < l; i++ {
				value := mValue.Index(i) // Value of item
				itemType := value.Type() // Type of item
				switch itemType.Kind() {
				case reflect.Map:
					// 递归加前缀
					m2 := value.Interface().(map[string]interface{})
					var subTarget = make(map[string]interface{})
					subTarget = AddKeyPrefixToMap(m2, subTarget, prefix, ignores...)
					arr = append(arr, subTarget)
				case reflect.Interface:
					val := value.Interface()
					if m2, ok := val.(map[string]interface{}); ok {
						var subTarget = make(map[string]interface{})
						subTarget = AddKeyPrefixToMap(m2, subTarget, prefix, ignores...)
						arr = append(arr, subTarget)
					} else {
						arr = append(arr, val)
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					arr = append(arr, value.Int())
				case reflect.String:
					arr = append(arr, value.String())
				case reflect.Bool:
					arr = append(arr, value.Bool())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					arr = append(arr, value.Uint())
				default:
					continue
				}
			}
			d = arr
		} else {
			d = m
		}
		if pass {
			target[k] = d
		} else {
			target[prefix+k] = d
		}
	}
	return target
}

// TrimKeyPrefixToMap 给 map 去除 key 前缀
// data    数据源 map
// target  目标   map
// prefix  要删除的前缀
// ignores 排除的字段
func TrimKeyPrefixToMap(data, target map[string]interface{}, prefix string, ignores ...string) map[string]interface{} {
	passMap := make(map[string]struct{})
	for _, ig := range ignores {
		passMap[ig] = struct{}{}
	}
	for k, m := range data {
		if m == nil {
			continue
		}
		_, pass := passMap[k]
		mType := reflect.TypeOf(m)
		mValue := reflect.ValueOf(m)
		var d interface{}
		if mType.Kind() == reflect.Map {
			m1 := m.(map[string]interface{})
			var subTarget = make(map[string]interface{})
			subTarget = TrimKeyPrefixToMap(m1, subTarget, prefix, ignores...)
			d = subTarget
		} else if mType.Kind() == reflect.Slice {
			l := mValue.Len()
			arr := make([]interface{}, 0, l)
			for i := 0; i < l; i++ {
				value := mValue.Index(i) // Value of item
				itemType := value.Type() // Type of item
				switch itemType.Kind() {
				case reflect.Map:
					// 递归去前缀
					m2 := value.Interface().(map[string]interface{})
					var subTarget = make(map[string]interface{})
					subTarget = TrimKeyPrefixToMap(m2, subTarget, prefix, ignores...)
					arr = append(arr, subTarget)
				case reflect.Interface:
					val := value.Interface()
					if m2, ok := val.(map[string]interface{}); ok {
						var subTarget = make(map[string]interface{})
						subTarget = TrimKeyPrefixToMap(m2, subTarget, prefix, ignores...)
						arr = append(arr, subTarget)
					} else {
						arr = append(arr, val)
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					arr = append(arr, value.Int())
				case reflect.String:
					arr = append(arr, value.String())
				case reflect.Bool:
					arr = append(arr, value.Bool())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					arr = append(arr, value.Uint())
				default:
					continue
				}
			}
			d = arr
		} else {
			d = m
		}
		if pass {
			target[k] = d
		} else {
			target[strings.TrimPrefix(k, prefix)] = d
		}
	}
	return target
}

// DeleteAbsentElementKeys 去除不存在的元素
// keys 要保留的 key 值
func DeleteAbsentElementKeys(data, target map[string]interface{}, keys []string) map[string]interface{} {
	if len(keys) == 0 {
		return target
	}
	keysMap := make(map[string]struct{})
	for _, k := range keys {
		keysMap[k] = struct{}{}
	}
	// 遍历 map 集合
	for k, m := range data {
		if m == nil {
			continue
		}
		if _, ok := keysMap[k]; !ok {
			continue
		}
		mType := reflect.TypeOf(m)
		mValue := reflect.ValueOf(m)
		var d interface{}
		if mType.Kind() == reflect.Map {
			// 如果类型是 map, 继续向下递归
			m1 := m.(map[string]interface{})
			var subTarget = make(map[string]interface{})
			subTarget = DeleteAbsentElementKeys(m1, subTarget, keys)
			d = subTarget
		} else if mType.Kind() == reflect.Slice {
			// 如果类型是 切片, 遍历切片
			l := mValue.Len()
			arr := make([]interface{}, 0, l)
			for i := 0; i < l; i++ {
				value := mValue.Index(i) // Value of item
				itemType := value.Type() // Type of item
				switch itemType.Kind() {
				case reflect.Map:
					// 递归去前缀
					m2 := value.Interface().(map[string]interface{})
					var subTarget = make(map[string]interface{})
					subTarget = DeleteAbsentElementKeys(m2, subTarget, keys)
					arr = append(arr, subTarget)
				case reflect.Interface:
					val := value.Interface()
					if m2, ok := val.(map[string]interface{}); ok {
						var subTarget = make(map[string]interface{})
						subTarget = DeleteAbsentElementKeys(m2, subTarget, keys)
						arr = append(arr, subTarget)
					} else {
						arr = append(arr, val)
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					arr = append(arr, value.Int())
				case reflect.String:
					arr = append(arr, value.String())
				case reflect.Bool:
					arr = append(arr, value.Bool())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					arr = append(arr, value.Uint())
				default:
					continue
				}
			}
			d = arr
		} else {
			d = m
		}
		target[k] = d
	}
	return target
}

// DeleteAbsentElementByMapping 去除不存在的元素
// mapping 要保留的 mapping
// 注意：方法里只取 mapping 的 key 值，尽可能给没有子集 value 为空结构体 struct{}
// 支持 mapping 结构如下
//
//	{
//	   "one":"",
//	   "onesub":{
//	       "tow":"",
//	       "towsub":{
//	           "three":"",
//	           "threesub":{
//
//	           }
//	       }
//	   }
//	}
func DeleteAbsentElementByMapping(data, target, mapping map[string]interface{}) map[string]interface{} {
	if len(mapping) == 0 {
		return target
	}

	// 遍历 map 集合
	for k, m := range data {
		if m == nil {
			continue
		}
		if _, ok := mapping[k]; !ok {
			continue
		}
		if _, ok := mapping[k].(map[string]interface{}); !ok {
			target[k] = m
			continue
		}
		mp := mapping[k].(map[string]interface{})
		mType := reflect.TypeOf(m)
		mValue := reflect.ValueOf(m)
		var d interface{}
		if mType.Kind() == reflect.Map {
			// 如果类型是 map, 继续向下递归
			m1 := m.(map[string]interface{})
			var subTarget = make(map[string]interface{})
			subTarget = DeleteAbsentElementByMapping(m1, subTarget, mp)
			d = subTarget
		} else if mType.Kind() == reflect.Slice {
			// 如果类型是 切片, 遍历切片
			l := mValue.Len()
			arr := make([]interface{}, 0, l)
			for i := 0; i < l; i++ {
				value := mValue.Index(i) // Value of item
				itemType := value.Type() // Type of item
				switch itemType.Kind() {
				case reflect.Map:
					// 递归去前缀
					m2 := value.Interface().(map[string]interface{})
					var subTarget = make(map[string]interface{})
					subTarget = DeleteAbsentElementByMapping(m2, subTarget, mp)
					arr = append(arr, subTarget)
				case reflect.Interface:
					val := value.Interface()
					if m2, ok := val.(map[string]interface{}); ok {
						var subTarget = make(map[string]interface{})
						subTarget = DeleteAbsentElementByMapping(m2, subTarget, mp)
						arr = append(arr, subTarget)
					} else {
						arr = append(arr, val)
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					arr = append(arr, value.Int())
				case reflect.String:
					arr = append(arr, value.String())
				case reflect.Bool:
					arr = append(arr, value.Bool())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					arr = append(arr, value.Uint())
				default:
					continue
				}
			}
			d = arr
		} else {
			d = m
		}
		target[k] = d
	}
	return target
}

func StructTagToMap(data interface{}, tag string) map[string]interface{} {
	if data == nil {
		return map[string]interface{}{}
	}
	if reflect.TypeOf(data).Kind() != reflect.Struct && reflect.TypeOf(data).Kind() != reflect.Pointer {
		return map[string]interface{}{}
	}
	var ts reflect.Type
	var vs reflect.Value
	if reflect.TypeOf(data).Kind() == reflect.Struct {
		ts = reflect.TypeOf(data)
		vs = reflect.ValueOf(data)
	}
	if reflect.TypeOf(data).Kind() == reflect.Pointer {
		ts = reflect.TypeOf(data).Elem()
		vs = reflect.ValueOf(data).Elem()
	}
	values := make(map[string]interface{})
	for i := 0; i < ts.NumField(); i++ {
		field := ts.Field(i)
		// TODO 未来考虑深层遍历
		jsonKey := field.Tag.Get(tag) // victim_phone,omitempty
		if strings.Contains(jsonKey, ",") {
			jsonKey = jsonKey[0:strings.Index(jsonKey, ",")]
		}
		values[jsonKey] = vs.Field(i).Interface()
	}
	return values
}
