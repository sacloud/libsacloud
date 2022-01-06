// Copyright 2016-2022 The Libsacloud Authors
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

	"github.com/sacloud/libsacloud/v2/pkg/size"
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
			expect: 1 * size.GiB,
		},
		{
			input:  2,
			expect: 2 * size.GiB,
		},
	}

	for _, tc := range expects {
		var target SizeMB = &dummySizeMBAccessor{}

		SetSizeGB(target, tc.input)
		require.Equal(t, tc.input, GetSizeGB(target))
		require.Equal(t, tc.expect, target.GetSizeMB())
	}
}
