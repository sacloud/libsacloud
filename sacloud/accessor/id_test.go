package accessor

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

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
		var target ID = &dummyIDAccessor{}

		if _, ok := tc.input.(string); ok {
			SetStringID(target, tc.input.(string))
		} else {
			SetInt64ID(target, tc.input.(int64))
		}

		require.Equal(t, tc.expect, target.GetID())
	}

}
