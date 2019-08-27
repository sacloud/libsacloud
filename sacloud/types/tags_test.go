package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTags_SortWhenUnmarshal(t *testing.T) {

	cases := []struct {
		input  string
		output Tags
	}{
		{
			input:  `["b","a","c"]`,
			output: Tags{"a", "b", "c"},
		},
	}
	for _, tc := range cases {
		var tags Tags
		if err := json.Unmarshal([]byte(tc.input), &tags); err != nil {
			t.Fatal(err)
		}
		require.Equal(t, tc.output, tags)
	}
}

func TestTags_SortWhenMarshal(t *testing.T) {

	cases := []struct {
		input  Tags
		output string
	}{
		{
			input:  Tags{"b", "a", "c"},
			output: `["a","b","c"]`,
		},
	}
	for _, tc := range cases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, tc.output, string(data))
	}
}
