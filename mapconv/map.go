package mapconv

import (
	"fmt"
	"strings"
)

// Map is wrapper of map[string]interface{}
type Map map[string]interface{}

// Map returns naked map
func (m *Map) Map() map[string]interface{} {
	return *m
}

// Set sets map value with dot-separated key
func (m *Map) Set(key string, value interface{}) {
	keys := strings.Split(key, ".")
	var dest map[string]interface{} = *m
	for i, k := range keys {
		last := i == len(keys)-1
		var v interface{}
		if last {
			v = value
		}
		setValueWithDefault(dest, k, v)
		if !last {
			dest = dest[k].(map[string]interface{})
		}
	}
}

// Get returns map value with dot-separated key
func (m *Map) Get(key string) (interface{}, error) {
	keys := strings.Split(key, ".")
	targetMap := *m
	for i, k := range keys {
		last := i == len(keys)-1
		value := targetMap[k]
		if value == nil {
			return nil, nil
		}
		if last {
			return value, nil
		}

		if v, ok := value.(map[string]interface{}); ok {
			targetMap = v
		} else {
			return nil, fmt.Errorf("key %q(part of %q) is not map[string]interface{}", k, key)
		}
	}

	return nil, fmt.Errorf("failed naked get tagged map: invalid state - key:%s values:%v", key, *m)
}

func setValueWithDefault(values map[string]interface{}, key string, value interface{}) {
	if value == nil {
		value = map[string]interface{}{}
	}
	if _, ok := values[key]; !ok {
		values[key] = value
	}
}
