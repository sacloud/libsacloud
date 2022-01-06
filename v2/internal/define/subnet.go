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

package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	subnetAPIName     = "Subnet"
	subnetAPIPathName = "subnet"
)

var subnetAPI = &dsl.Resource{
	Name:       subnetAPIName,
	PathName:   subnetAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(subnetAPIName, subnetNakedType, findParameter, subnetView),
		// read
		ops.Read(subnetAPIName, subnetNakedType, subnetView),
	},
}
var (
	subnetNakedType = meta.Static(naked.Subnet{})

	subnetView = &dsl.Model{
		Name:      subnetAPIName,
		NakedType: subnetNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.SwitchID(),
			fields.Def("InternetID", meta.TypeID, mapConvTag("Switch.Internet.ID,omitempty")),
			fields.DefaultRoute(),
			fields.NextHop(),
			fields.StaticRoute(),
			fields.NetworkAddress(),
			fields.NetworkMaskLen(),
			{
				Name: "IPAddresses",
				Type: &dsl.Model{
					Name:      "SubnetIPAddress",
					NakedType: ipNakedType,
					IsArray:   true,
					Fields: []*dsl.FieldDesc{
						fields.HostName(),
						fields.IPAddress(),
					},
				},
				Tags: mapConvTag("[]IPAddresses,recursive"),
			},
		},
	}
)
