package mapconv

import (
	"fmt"
	"reflect"
	"strings"
)

// Map is wrapper of map[string]interface{}
type Map map[string]interface{}

// Map returns output map
func (m *Map) Map() map[string]interface{} {
	return *m
}

// Set sets map value with dot-separated key
func (m *Map) Set(key string, value interface{}) {
	keys := strings.Split(key, ".")
	var dest map[string]interface{} = *m
	for i, k := range keys {
		last := i == len(keys)-1
		isSlice := strings.HasPrefix(k, "[]")
		k = strings.Replace(k, "[]", "", -1)

		var v interface{}
		if last {
			v = value
		}

		if !last && isSlice {
			values := valueToSlice(value)
			var nestMap []map[string]interface{}
			for _, value := range values {
				nested := Map(map[string]interface{}{})
				key := strings.Join(keys[i+1:], ".")
				nested.Set(key, value)
				nestMap = append(nestMap, nested)
			}
			if _, ok := dest[k]; !ok {
				dest[k] = nestMap
			} else {
				existed, ok := dest[k].([]map[string]interface{})
				if !ok {
					dest[k] = nestMap
				}
				dest[k] = append(existed, nestMap...)
			}
			return
		}

		setValueWithDefault(dest, k, v)
		if !last {
			dest = dest[k].(map[string]interface{})
		}
	}
}

func valueToSlice(value interface{}) []interface{} {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Slice {
		ret := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			ret[i] = v.Index(i).Interface()
		}
		return ret
	}
	return []interface{}{value}
}

func setValueWithDefault(values map[string]interface{}, key string, value interface{}) {
	if value == nil {
		value = map[string]interface{}{}
	}
	if _, ok := values[key]; !ok {
		values[key] = value
	}
}

// Get returns map value with dot-separated key
func (m *Map) Get(key string) (interface{}, error) {
	keys := strings.Split(key, ".")
	targetMap := *m
	for i, k := range keys {
		last := i == len(keys)-1
		k = strings.Replace(k, "[]", "", -1)

		value := targetMap[k]
		if value == nil {
			return nil, nil
		}
		if last {
			return value, nil
		}

		switch value := value.(type) {
		case map[string]interface{}:
			targetMap = value
		case []interface{}:
			var values []interface{}
			for _, v := range value {
				if _, ok := v.(map[string]interface{}); !ok {
					return nil, fmt.Errorf("elements of key %q(part of %q) are not map[string]interface{}", k, key)
				}
				nested := Map(v.(map[string]interface{}))
				key := strings.Join(keys[i+1:], ".")
				nv, err := nested.Get(key)
				if err != nil {
					return nil, err
				}
				if nv != nil {
					nvs, ok := nv.([]interface{})
					if ok {
						values = append(values, nvs...)
					} else {
						values = append(values, nv)
					}
				}
			}
			return values, nil
		case []map[string]interface{}:
			var values []interface{}
			for _, v := range value {
				nested := Map(v)
				key := strings.Join(keys[i+1:], ".")
				nv, err := nested.Get(key)
				if err != nil {
					return nil, err
				}
				if nv != nil {
					nvs, ok := nv.([]interface{})
					if ok {
						values = append(values, nvs...)
					} else {
						values = append(values, nv)
					}
				}
			}
			return values, nil
		default:
			return nil, fmt.Errorf("key %q(part of %q) is not map[string]interface{} or []map[string]interface{}", k, key)
		}
	}

	return nil, fmt.Errorf("failed output get input map: invalid state - key:%s values:%v", key, *m)
}
