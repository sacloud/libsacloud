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

package sacloud

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestServer_BandWidthAt(t *testing.T) {
	cases := []struct {
		msg    string
		in     *Server
		index  int
		expect int
	}{
		{
			msg:    "no NICs",
			in:     &Server{},
			expect: -1,
		},
		{
			msg: "invalid NIC index",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Unknown,
					},
				},
			},
			index:  1,
			expect: -1,
		},
		{
			msg: "unknown upstream type",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Unknown,
					},
				},
			},
			expect: -1,
		},
		{
			msg: "shared",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Shared,
					},
				},
			},
			expect: 100,
		},
		{
			msg: "switch with private host",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Switch,
					},
				},
				PrivateHostID: 1,
			},
			expect: 0,
		},
		{
			msg: "memory less than 32GB",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Switch,
					},
				},
				MemoryMB: 31 * size.GiB,
			},
			expect: 1000,
		},
		{
			msg: "switch with memory 32GB",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Switch,
					},
				},
				MemoryMB: 32 * size.GiB,
			},
			expect: 2000,
		},
		{
			msg: "memory less than 128GB",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Switch,
					},
				},
				MemoryMB: 127 * size.GiB,
			},
			expect: 2000,
		},
		{
			msg: "switch with memory 128GB",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Switch,
					},
				},
				MemoryMB: 128 * size.GiB,
			},
			expect: 5000,
		},
		{
			msg: "memory less than 224GB",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Switch,
					},
				},
				MemoryMB: 223 * size.GiB,
			},
			expect: 5000,
		},
		{
			msg: "switch with memory 224GB",
			in: &Server{
				Interfaces: []*InterfaceView{
					{
						UpstreamType: types.UpstreamNetworkTypes.Switch,
					},
				},
				MemoryMB: 224 * size.GiB,
			},
			expect: 10000,
		},
	}

	for _, tc := range cases {
		actual := tc.in.BandWidthAt(tc.index)
		require.Equal(t, tc.expect, actual, tc.msg)
	}
}
