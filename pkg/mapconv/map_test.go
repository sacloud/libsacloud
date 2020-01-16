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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap_Set(t *testing.T) {
	expects := []struct {
		caseName string
		source   map[string]interface{}
		dest     map[string]interface{}
	}{
		{
			caseName: "minimum",
			source: map[string]interface{}{
				"test": "test",
			},
			dest: map[string]interface{}{
				"test": "test",
			},
		},
		{
			caseName: "nested",
			source: map[string]interface{}{
				"test.A":    "A",
				"test.B":    "B",
				"test.C.C1": "C1",
				"test.C.C2": "C2",
				"outer":     "outer",
				"int":       1,
				"float":     1.1,
			},
			dest: map[string]interface{}{
				"test": map[string]interface{}{
					"A": "A",
					"B": "B",
					"C": map[string]interface{}{
						"C1": "C1",
						"C2": "C2",
					},
				},
				"outer": "outer",
				"int":   1,
				"float": 1.1,
			},
		},
		{
			caseName: "slice",
			source: map[string]interface{}{
				"slice.slice.value": []interface{}{"value4", "value5"},
			},
			dest: map[string]interface{}{
				"slice": map[string]interface{}{
					"slice": map[string]interface{}{
						"value": []interface{}{"value4", "value5"},
					},
				},
			},
		},
		{
			caseName: "expanded slice",
			source: map[string]interface{}{
				"[]slice.value": []interface{}{"value1", "value2"},
			},
			dest: map[string]interface{}{
				"slice": []map[string]interface{}{
					{"value": "value1"},
					{"value": "value2"},
				},
			},
		},
		{
			caseName: "expanded nested slice",
			source: map[string]interface{}{
				"[]slice.slice.value": []interface{}{"value4", "value5"},
			},
			dest: map[string]interface{}{
				"slice": []map[string]interface{}{
					{
						"slice": map[string]interface{}{
							"value": "value4",
						},
					},
					{
						"slice": map[string]interface{}{
							"value": "value5",
						},
					},
				},
			},
		},
		{
			caseName: "expanded nested slice with middle slice",
			source: map[string]interface{}{
				"slice.[]slice.value": []interface{}{"value4", "value5"},
			},
			dest: map[string]interface{}{
				"slice": map[string]interface{}{
					"slice": []map[string]interface{}{
						{"value": "value4"},
						{"value": "value5"},
					},
				},
			},
		},
		{
			caseName: "expanded nested slice with last slice",
			source: map[string]interface{}{
				"slice.slice.[]value": []interface{}{"value4", "value5"},
			},
			dest: map[string]interface{}{
				"slice": map[string]interface{}{
					"slice": map[string]interface{}{
						"value": []interface{}{"value4", "value5"},
					},
				},
			},
		},
		{
			caseName: "expanded deep nested slice",
			source: map[string]interface{}{
				"[]slice.[]slice.value": []interface{}{"value4", "value5"},
			},
			dest: map[string]interface{}{
				"slice": []map[string]interface{}{
					{
						"slice": []map[string]interface{}{
							{"value": "value4"},
						},
					},
					{
						"slice": []map[string]interface{}{
							{"value": "value5"},
						},
					},
				},
			},
		},
	}

	for _, expect := range expects {
		t.Run(expect.caseName, func(t *testing.T) {
			m := Map(make(map[string]interface{}))
			for k, v := range expect.source {
				m.Set(k, v)
			}
			require.Equal(t, expect.dest, m.Map())
		})
	}
}

func TestMap_Get(t *testing.T) {
	expects := []struct {
		caseName  string
		keyValues map[string]interface{}
		source    map[string]interface{}
		err       error
	}{
		{
			caseName: "minimum",
			keyValues: map[string]interface{}{
				"test": "test",
			},
			source: map[string]interface{}{
				"test": "test",
			},
		},
		{
			caseName: "nested",
			keyValues: map[string]interface{}{
				"test.A":    "A",
				"test.B":    "B",
				"test.C.C1": "C1",
				"test.C.C2": "C2",
				"outer":     "outer",
				"int":       1,
				"float":     1.1,
			},
			source: map[string]interface{}{
				"test": map[string]interface{}{
					"A": "A",
					"B": "B",
					"C": map[string]interface{}{
						"C1": "C1",
						"C2": "C2",
					},
				},
				"outer": "outer",
				"int":   1,
				"float": 1.1,
			},
		},
		{
			caseName: "slice",
			keyValues: map[string]interface{}{
				"slice.value": []interface{}{"value1", "value2"},
			},
			source: map[string]interface{}{
				"slice": []map[string]interface{}{
					{"value": "value1"},
					{"value": "value2"},
				},
			},
		},
		{
			caseName: "nested slice",
			keyValues: map[string]interface{}{
				"slice.slice.value": []interface{}{"value4", "value5"},
			},
			source: map[string]interface{}{
				"slice": []map[string]interface{}{
					{"value": "value1"},
					{"value": "value2"},
					{
						"value": "value3",
						"slice": []map[string]interface{}{
							{"value": "value4"},
							{"value": "value5"},
						},
					},
				},
			},
		},
		{
			caseName: "with error",
			keyValues: map[string]interface{}{
				"test.A.B": "test",
			},
			source: map[string]interface{}{
				"test": map[string]interface{}{
					"A": "test",
				},
			},
			err: errors.New(`key "A"(part of "test.A.B") is not map[string]interface{} or []map[string]interface{}`),
		},
	}

	for _, expect := range expects {
		t.Run(expect.caseName, func(t *testing.T) {
			m := Map(expect.source)
			for k, v := range expect.keyValues {
				value, err := m.Get(k)
				require.Equal(t, expect.err, err)
				if err == nil {
					require.Equal(t, v, value)
				}
			}
		})
	}
}
