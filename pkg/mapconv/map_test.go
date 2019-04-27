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
			caseName: "with error",
			keyValues: map[string]interface{}{
				"test.A.B": "test",
			},
			source: map[string]interface{}{
				"test": map[string]interface{}{
					"A": "test",
				},
			},
			err: errors.New(`key "A"(part of "test.A.B") is not map[string]interface{}`),
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
