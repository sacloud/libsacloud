// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mapconv

import (
	"errors"
	"math"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

// ConvertTo converts struct which input by mapconv to plain models
func ConvertTo(source interface{}, dest interface{}) error {
	s := structs.New(source)
	destMap := Map(make(map[string]interface{}))

	fields := s.Fields()
	for _, f := range fields {
		//if !f.IsExported() || f.IsZero() {
		if !f.IsExported() {
			continue
		}

		tags := ParseMapConvTag(f.Tag("mapconv"))
		for _, key := range tags.SourceFields {
			destKey := f.Name()
			value := f.Value()

			if key != "" {
				destKey = key
			}
			if f.IsZero() {
				if tags.OmitEmpty {
					continue
				}
				if tags.DefaultValue != nil {
					value = tags.DefaultValue
				}
			}

			if tags.Squash {
				d := Map(make(map[string]interface{}))
				err := ConvertTo(value, &d)
				if err != nil {
					return err
				}
				for k, v := range d {
					destMap.Set(k, v)
				}
				continue
			}

			if tags.Recursive {
				var dest []interface{}
				values := valueToSlice(value)
				for _, v := range values {
					if structs.IsStruct(v) {
						destMap := Map(make(map[string]interface{}))
						if err := ConvertTo(v, &destMap); err != nil {
							return err
						}
						dest = append(dest, destMap)
					} else {
						dest = append(dest, v)
					}
				}
				if tags.IsSlice || dest == nil || len(dest) > 1 {
					value = dest
				} else {
					value = dest[0]
				}
			}

			destMap.Set(destKey, value)
		}
	}

	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           dest,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(destMap.Map())
}

// ConvertFrom converts struct which input by mapconv from plain models
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

		tags := ParseMapConvTag(f.Tag("mapconv"))
		if tags.Squash {
			return errors.New("ConvertFrom is not allowed squash")
		}
		for _, key := range tags.SourceFields {
			sourceKey := f.Name()
			if key != "" {
				sourceKey = key
			}

			value, err := sourceMap.Get(sourceKey)
			if err != nil {
				return err
			}
			if value == nil || isZero(reflect.ValueOf(value)) {
				continue
			}

			if tags.Recursive {
				t := reflect.TypeOf(f.Value())
				if t.Kind() == reflect.Slice {
					t = t.Elem().Elem()
				} else {
					t = t.Elem()
				}

				var dest []interface{}
				values := valueToSlice(value)
				for _, v := range values {
					if v == nil {
						dest = append(dest, v)
						continue
					}
					d := reflect.New(t).Interface()
					if err := ConvertFrom(v, d); err != nil {
						return err
					}
					dest = append(dest, d)
				}

				if dest != nil {
					if tags.IsSlice || len(dest) > 1 {
						value = dest
					} else {
						value = dest[0]
					}
				}
			}

			destMap.Set(f.Name(), value)
		}
	}
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           dest,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(destMap.Map())
}

// TagInfo mapconvタグの情報
type TagInfo struct {
	SourceFields []string
	DefaultValue interface{}
	OmitEmpty    bool
	Recursive    bool
	Squash       bool
	IsSlice      bool
}

// ParseMapConvTag mapconvタグを文字列で受け取りパースしてTagInfoを返す
func ParseMapConvTag(tagBody string) TagInfo {
	tokens := strings.Split(tagBody, ",")
	key := tokens[0]

	keys := strings.Split(key, "/")
	var defaultValue interface{}
	var omitEmpty, recursive, squash, isSlice bool

	for _, k := range keys {
		if strings.Contains(k, "[]") {
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
		case strings.HasPrefix(token, "squash"):
			squash = true
		case strings.HasPrefix(token, "default"):
			keyValue := strings.Split(token, "=")
			if len(keyValue) > 1 {
				defaultValue = strings.Join(keyValue[1:], "")
			}
		}
	}
	return TagInfo{
		SourceFields: keys,
		DefaultValue: defaultValue,
		OmitEmpty:    omitEmpty,
		Recursive:    recursive,
		Squash:       squash,
		IsSlice:      isSlice,
	}
}

// isZero go 1.13 からのポート
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(&reflect.ValueError{
			Method: "reflect.Value.IsZero",
			Kind:   v.Kind(),
		})
	}
}
