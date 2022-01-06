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
	ipv6netAPIName     = "IPv6Net"
	ipv6netAPIPathName = "ipv6net"
)

var ipv6netAPI = &dsl.Resource{
	Name:       ipv6netAPIName,
	PathName:   ipv6netAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// list Note: Findのパラメータなしと同等だが、後方互換のために残しておく
		ops.List(ipv6netAPIName, ipv6netNakedType, ipv6netView),

		// find
		ops.Find(ipv6netAPIName, ipv6netNakedType, findParameter, ipv6netView),

		// read
		ops.Read(ipv6netAPIName, ipv6netNakedType, ipv6netView),
	},
}
var (
	ipv6netNakedType = meta.Static(naked.IPv6Net{})

	ipv6netView = &dsl.Model{
		Name:      ipv6netAPIName,
		NakedType: ipv6netNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Def("ServiceID", meta.TypeID),
			fields.Def("IPv6Prefix", meta.TypeString),
			fields.Def("IPv6PrefixLen", meta.TypeInt),
			fields.Def("IPv6PrefixTail", meta.TypeString),
			fields.Def("ServiceClass", meta.TypeString),
			fields.Def("IPv6TableID", meta.TypeID, mapConvTag("IPv6Table.ID")),
			fields.Def("NamedIPv6AddrCount", meta.TypeInt),
			fields.CreatedAt(),
			fields.SwitchID(),
		},
	}
)
