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

package fake

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextSubnet(t *testing.T) {
	DataStore = NewInMemoryStore()
	// internal init
	vp = initValuePool(DataStore)

	first := pool().nextSubnet(24)
	require.Equal(t, "24.0.1.0", first.networkAddress)
	require.Equal(t, 24, first.networkMaskLen)
	require.Len(t, first.addresses, 251)
	require.Equal(t, "24.0.1.4", first.addresses[0])
	require.Equal(t, "24.0.1.254", first.addresses[len(first.addresses)-1])

	next := pool().nextSubnet(24)
	require.Equal(t, "24.0.2.0", next.networkAddress)
	require.Equal(t, 24, next.networkMaskLen)
	require.Len(t, next.addresses, 251)
	require.Equal(t, "24.0.2.4", next.addresses[0])
	require.Equal(t, "24.0.2.254", next.addresses[len(next.addresses)-1])
}
