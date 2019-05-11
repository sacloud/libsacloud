package sacloud

import (
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
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
	id types.ID
}

func (d *dummyIDAccessor) GetID() types.ID {
	return d.id
}

func (d *dummyIDAccessor) SetID(id types.ID) {
	d.id = id
}

func TestIDAccessor(t *testing.T) {
	expects := []struct {
		input  interface{}
		expect types.ID
	}{
		{
			input:  "0",
			expect: types.Int64ID(0),
		},
		{
			input:  "1",
			expect: types.Int64ID(1),
		},
		{
			input:  "2",
			expect: types.Int64ID(2),
		},
	}

	for _, tc := range expects {
		var target idAccessor = &dummyIDAccessor{}

		if _, ok := tc.input.(string); ok {
			setStringID(target, tc.input.(string))
		} else {
			setInt64ID(target, tc.input.(int64))
		}

		require.Equal(t, tc.expect, target.GetID())
	}

}
