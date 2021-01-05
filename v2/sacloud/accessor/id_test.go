// Copyright 2016-2021 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
