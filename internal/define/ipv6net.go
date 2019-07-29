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
		// List
		ops.List(ipv6netAPIName, ipv6netNakedType, ipv6netView),

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
