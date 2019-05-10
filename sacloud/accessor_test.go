package sacloud

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
		var target sizeMBAccessor = &dummySizeMBAccessor{}

		setSizeGB(target, tc.input)
		require.Equal(t, tc.input, getSizeGB(target))
		require.Equal(t, tc.expect, target.GetSizeMB())
	}
}

type dummyIDAccessor struct {
	id int64
}

func (d *dummyIDAccessor) GetID() int64 {
	return d.id
}

func (d *dummyIDAccessor) SetID(id int64) {
	d.id = id
}

func TestIDAccessor(t *testing.T) {
	expects := []struct {
		input  string
		expect int64
	}{
		{
			input:  "0",
			expect: 0,
		},
		{
			input:  "1",
			expect: 1,
		},
		{
			input:  "2",
			expect: 2,
		},
	}

	for _, tc := range expects {
		var target idAccessor = &dummyIDAccessor{}

		setStringID(target, tc.input)
		require.Equal(t, tc.input, getStringID(target))
		require.Equal(t, tc.expect, target.GetID())
	}

}
