package accessor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type dummySizeMBAccessor struct {
	size int
}

func (d *dummySizeMBAccessor) GetSizeMB() int {
	return d.size
}

func (d *dummySizeMBAccessor) SetSizeMB(size int) {
	d.size = size
}

func TestSizeMBAccessor(t *testing.T) {
	expects := []struct {
		input  int
		expect int
	}{
		{
			input:  0,
			expect: 0,
		},
		{
			input:  1,
			expect: 1024 * 1,
		},
		{
			input:  2,
			expect: 1024 * 2,
		},
	}

	for _, tc := range expects {
		var target SizeMB = &dummySizeMBAccessor{}

		SetSizeGB(target, tc.input)
		require.Equal(t, tc.input, GetSizeGB(target))
		require.Equal(t, tc.expect, target.GetSizeMB())
	}
}
