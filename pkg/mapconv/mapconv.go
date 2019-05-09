package mapconv

import (
	"encoding/json"

	"github.com/fatih/structs"
)

// ToNaked converts struct naked naked models by using mapconv tag
func ToNaked(source interface{}, dest interface{}) error {
	s := structs.New(source)
	destMap := Map(make(map[string]interface{}))

	fields := s.Fields()
	for _, f := range fields {
		if !f.IsExported() || f.IsZero() {
			continue
		}
		value := f.Value()

		destKey := f.Name()
		tag := f.Tag("mapconv")
		if tag != "" {
			destKey = tag
		}
		destMap.Set(destKey, value)
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

		key := f.Name()
		tag := f.Tag("mapconv")
		if tag != "" {
			key = tag
		}

		value, err := sourceMap.Get(key)
		if err != nil {
			return err
		}
		if value == nil {
			continue
		}
		destMap.Set(f.Name(), value)
	}

	data, err := json.Marshal(destMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}
