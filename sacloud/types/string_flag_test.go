package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringFlag(t *testing.T) {

	expects := []struct {
		input  string
		expect StringFlag
	}{
		{input: `"True"`, expect: StringTrue},
		{input: `"true"`, expect: StringTrue},
		{input: `"False"`, expect: StringFalse},
		{input: `"false"`, expect: StringFalse},
		{input: `true`, expect: StringTrue},
		{input: `false`, expect: StringFalse},
		{input: `"On"`, expect: StringTrue},
		{input: `"on"`, expect: StringTrue},
		{input: `"Off"`, expect: StringFalse},
		{input: `"off"`, expect: StringFalse},
		{input: `"1"`, expect: StringTrue},
		{input: `"0"`, expect: StringFalse},
		{input: `1`, expect: StringTrue},
		{input: `0`, expect: StringFalse},
		{input: `""`, expect: StringFalse},
		{input: `"hoge"`, expect: StringFalse},
	}

	for _, tc := range expects {
		var f StringFlag
		err := json.Unmarshal([]byte(tc.input), &f)

		require.NotNil(t, f)
		require.NoError(t, err, "expect: %#v", tc)
		require.Equal(t, tc.expect, f)
	}

}
