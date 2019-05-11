package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResult_UnmarshalJSON(t *testing.T) {

	expects := []struct {
		input    string
		expect   APIResult
		hasError bool
	}{
		{
			input:    "true",
			expect:   ResultSuccess,
			hasError: false,
		},
		{
			input:    "false",
			expect:   ResultFailed,
			hasError: false,
		},
		{
			input:    `"Accepted"`,
			expect:   ResultAccepted,
			hasError: false,
		},
		{
			input:    "2",
			expect:   ResultUnknown,
			hasError: true,
		},
		{
			input:    `"Deleted"`,
			expect:   ResultUnknown,
			hasError: false,
		},
	}

	for _, tc := range expects {
		var res APIResult
		err := json.Unmarshal([]byte(tc.input), &res)

		if tc.hasError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}

		require.Equal(t, tc.expect, res)
	}

}
