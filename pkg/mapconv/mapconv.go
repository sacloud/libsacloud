package mapconv

import (
	"encoding/json"
	"strings"

	"github.com/fatih/structs"
)

// ToNaked converts struct naked naked models by using mapconv tag
func ToNaked(source interface{}, dest interface{}) error {
	s := structs.New(source)
	destMap := Map(make(map[string]interface{}))

	fields := s.Fields()
	for _, f := range fields {
		//if !f.IsExported() || f.IsZero() {
		if !f.IsExported() {
			continue
		}

		tags := mapConv(f.Tag("mapconv")).values()
		for _, tag := range tags {
			destKey := f.Name()
			value := f.Value()

			if tag.key != "" {
				destKey = tag.key
			}
			if f.IsZero() && tag.defaultValue != nil {
				value = tag.defaultValue
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

// FromNaked converts naked models naked struct by using mapconv tag
func FromNaked(source interface{}, dest interface{}) error {
	sourceMap := Map(structs.New(source).Map())
	destMap := Map(make(map[string]interface{}))

	s := structs.New(dest)
	fields := s.Fields()
	for _, f := range fields {
		if !f.IsExported() {
			continue
		}

		tags := mapConv(f.Tag("mapconv")).values()
		for _, tag := range tags {
			key := f.Name()
			if tag.key != "" {
				key = tag.key
			}
			value := f.Value()

			value, err := sourceMap.Get(key)
			if err != nil {
				return err
			}
			if value == nil {
				continue
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
	key          string
	defaultValue interface{}
}

func (m mapConv) values() []*mapConvValue {
	var values []*mapConvValue
	tokens := strings.Split(string(m), ",")
	for _, token := range tokens {
		keyValues := strings.Split(token, ":")
		key := keyValues[0]
		var def interface{}
		if len(keyValues) > 1 {
			def = strings.Join(keyValues[1:], "")
		}
		values = append(values, &mapConvValue{
			key:          key,
			defaultValue: def,
		})
	}
	return values
}
