package mapconv

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

// ConvertTo converts struct which tagged by mapconv to plain models
func ConvertTo(source interface{}, dest interface{}) error {
	s := structs.New(source)
	destMap := Map(make(map[string]interface{}))

	fields := s.Fields()
	for _, f := range fields {
		//if !f.IsExported() || f.IsZero() {
		if !f.IsExported() {
			continue
		}

		tags := mapConv(f.Tag("mapconv")).value()
		for _, key := range tags.keys {
			destKey := f.Name()
			value := f.Value()

			if key != "" {
				destKey = key
			}
			if f.IsZero() {
				if tags.omitEmpty {
					continue
				}
				if tags.defaultValue != nil {
					value = tags.defaultValue
				}
			}

			if tags.recursive {
				var dest []interface{}
				values := valueToSlice(value)
				for _, v := range values {
					destMap := Map(make(map[string]interface{}))
					if err := ConvertTo(v, &destMap); err != nil {
						return err
					}
					dest = append(dest, destMap)
				}
				if tags.isSlice || len(dest) > 1 {
					value = dest
				} else {
					value = dest[0]
				}
			}

			destMap.Set(destKey, value)
		}
	}

	data, err := json.Marshal(destMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// ConvertFrom converts struct which tagged by mapconv from plain models
func ConvertFrom(source interface{}, dest interface{}) error {
	var sourceMap Map
	if m, ok := source.(map[string]interface{}); ok {
		sourceMap = Map(m)
	} else {
		sourceMap = Map(structs.New(source).Map())
	}
	destMap := Map(make(map[string]interface{}))

	s := structs.New(dest)
	fields := s.Fields()
	for _, f := range fields {
		if !f.IsExported() {
			continue
		}

		tags := mapConv(f.Tag("mapconv")).value()
		for _, key := range tags.keys {
			sourceKey := f.Name()
			if key != "" {
				sourceKey = key
			}

			value, err := sourceMap.Get(sourceKey)
			if err != nil {
				return err
			}
			if value == nil {
				continue
			}

			if tags.recursive {
				t := reflect.TypeOf(f.Value())
				if t.Kind() == reflect.Slice {
					t = t.Elem().Elem()
				} else {
					t = t.Elem()
				}

				var dest []interface{}
				values := valueToSlice(value)
				for _, v := range values {

					d := reflect.New(t).Interface()
					if err := ConvertFrom(v, d); err != nil {
						return err
					}
					dest = append(dest, d)
				}

				if dest != nil {
					if tags.isSlice || len(dest) > 1 {
						value = dest
					} else {
						value = dest[0]
					}
				}
			}

			destMap.Set(f.Name(), value)
		}

	}

	data, err := json.Marshal(destMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

type mapConv string

type mapConvValue struct {
	keys         []string
	defaultValue interface{}
	omitEmpty    bool
	recursive    bool
	isSlice      bool
}

func (m mapConv) value() mapConvValue {
	tokens := strings.Split(string(m), ",")
	key := tokens[0]

	keys := strings.Split(key, "/")
	var defaultValue interface{}
	var omitEmpty, recursive, isSlice bool

	if len(keys) > 0 {
		if strings.HasPrefix(keys[0], "[]") {
			isSlice = true
		}
	}

	for i, token := range tokens {
		if i == 0 {
			continue
		}

		switch {
		case strings.HasPrefix(token, "omitempty"):
			omitEmpty = true
		case strings.HasPrefix(token, "recursive"):
			recursive = true
		case strings.HasPrefix(token, "default"):
			keyValue := strings.Split(token, "=")
			if len(keyValue) > 1 {
				defaultValue = strings.Join(keyValue[1:], "")
			}
		}

	}
	return mapConvValue{
		keys:         keys,
		defaultValue: defaultValue,
		omitEmpty:    omitEmpty,
		recursive:    recursive,
		isSlice:      isSlice,
	}
}
