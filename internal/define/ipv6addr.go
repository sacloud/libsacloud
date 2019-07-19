package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	ipv6AddrAPIName     = "IPv6Addr"
	ipv6AddrAPIPathName = "ipv6addr"
)

var ipv6AddrAPI = &dsl.Resource{
	Name:       ipv6AddrAPIName,
	PathName:   ipv6AddrAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(ipv6AddrAPIName, ipv6AddrNakedType, findParameter, ipv6AddrView),

		// create
		ops.Create(ipv6AddrAPIName, ipv6AddrNakedType, ipv6AddrCreateParam, ipv6AddrView),

		// read (IDの代わりにipv6addrを利用)
		{
			ResourceName: ipv6AddrAPIName,
			Name:         "Read",
			PathFormat:   dsl.DefaultPathFormatWithID,
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				argIPv6Addr,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: ipv6AddrNakedType,
				Name: ipv6AddrAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: ipv6AddrAPIName,
					DestField:   ipv6AddrView.Name,
					IsPlural:    false,
					Model:       ipv6AddrView,
				},
			},
		},

		// update (IDの代わりにipv6addrを利用)
		{
			ResourceName: ipv6AddrAPIName,
			Name:         "Update",
			PathFormat:   dsl.DefaultPathFormatWithID,
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: ipv6AddrNakedType,
				Name: ipv6AddrAPIName,
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				argIPv6Addr,
				dsl.MappableArgument("param", ipv6AddrUpdateParam, ipv6AddrAPIName),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: ipv6AddrNakedType,
				Name: ipv6AddrAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: ipv6AddrAPIName,
					DestField:   ipv6AddrView.Name,
					IsPlural:    false,
					Model:       ipv6AddrView,
				},
			},
		},

		// delete (IDの代わりにipv6addrを利用)
		{
			ResourceName: ipv6AddrAPIName,
			Name:         "Delete",
			PathFormat:   dsl.DefaultPathFormatWithID,
			Method:       http.MethodDelete,
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				argIPv6Addr,
			},
		},
	},
}

var (
	argIPv6Addr = &dsl.Argument{
		Name:            "ipv6addr",
		Type:            meta.TypeString,
		PathFormatAlias: "id",
	}
)

var (
	ipv6AddrNakedType = meta.Static(naked.IPv6Addr{})

	ipv6AddrView = &dsl.Model{
		Name:      ipv6AddrAPIName,
		NakedType: ipv6AddrNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Def("IPv6Addr", meta.TypeString),
			fields.HostName(),
			fields.Def("IPv6NetID", meta.TypeID, mapConvTag("IPv6Net.ID")),
			fields.Def("SwitchID", meta.TypeID, mapConvTag("IPv6Net.Switch.ID")),
			fields.InterfaceID(),
		},
	}

	ipv6AddrCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(ipv6AddrAPIName),
		NakedType: ipv6AddrNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Def("IPv6Addr", meta.TypeString),
			fields.HostName(),
		},
	}

	ipv6AddrUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(ipv6AddrAPIName),
		NakedType: ipv6AddrNakedType,
		Fields: []*dsl.FieldDesc{
			fields.HostName(),
		},
	}
)
