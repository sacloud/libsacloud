package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringNumber(t *testing.T) {

	expects := []struct {
		input  string
		expect StringNumber
	}{
		{input: `""`, expect: StringNumber(0)},
		{input: `"1"`, expect: StringNumber(1)},
		{input: `0`, expect: StringNumber(0)},
		{input: `1`, expect: StringNumber(1)},
	}

	for _, tc := range expects {
		var n StringNumber
		err := json.Unmarshal([]byte(tc.input), &n)

		require.NotNil(t, n)
		require.NoError(t, err, "expect: %#v", tc)
		require.Equal(t, tc.expect, n)
	}

}
