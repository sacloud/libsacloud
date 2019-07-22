package search

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEqualExpression(t *testing.T) {
	targetTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	targetTimeString := targetTime.Format(time.RFC3339)

	cases := []struct {
		input  *EqualExpression
		expect string
	}{
		{
			input:  AndEqual("c1", "c2"),
			expect: `"c1%20c2"`,
		},
		{
			input:  AndEqual("c1", "", "c2"),
			expect: `"c1%20%20c2"`,
		},
		{
			input:  OrEqual("c1", "c2"),
			expect: `["c1","c2"]`,
		},
		{
			input:  OrEqual("c1", nil, "c2"),
			expect: `["c1","c2"]`,
		},
		{
			input:  OrEqual(1, 2),
			expect: `["1","2"]`,
		},
		{
			input:  OrEqual(targetTime),
			expect: fmt.Sprintf(`["%s"]`, targetTimeString),
		},
	}

	for _, tc := range cases {
		data, err := json.Marshal(tc.input)
		require.NoError(t, err)
		require.Equal(t, tc.expect, string(data))
	}
}
