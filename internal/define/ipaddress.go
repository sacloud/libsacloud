package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	ipAPIName     = "IPAddress"
	ipAPIPathName = "ipaddress"
)

var ipAPI = &dsl.Resource{
	Name:       ipAPIName,
	PathName:   ipAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		{
			ResourceName:     ipAPIName,
			Name:             "List",
			PathFormat:       dsl.DefaultPathFormat,
			Method:           http.MethodGet,
			UseWrappedResult: true,
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: ipNakedType,
				Name: ipAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: ipAPIName,
					DestField:   names.ResourceFieldName(ipAPIName, dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       ipView,
				},
			},
		},
		// read
		{
			ResourceName: ipAPIName,
			Name:         "Read",
			PathFormat:   dsl.DefaultPathFormat + "/{{.ipAddress}}",
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				{
					Name: "ipAddress",
					Type: meta.TypeString,
				},
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: ipNakedType,
				Name: ipAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: ipAPIName,
					DestField:   ipAPIName,
					IsPlural:    false,
					Model:       ipView,
				},
			},
		},

		// set reverse
		{
			ResourceName: ipAPIName,
			Name:         "UpdateHostName",
			PathFormat:   dsl.DefaultPathFormat + "/{{.ipAddress}}",
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: ipNakedType,
				Name: ipAPIName,
			}),
			Arguments: dsl.Arguments{
				{
					Name: "ipAddress",
					Type: meta.TypeString,
				},
				{
					Name:       "hostName",
					Type:       meta.TypeString,
					MapConvTag: "IPAddress.HostName",
				},
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: ipNakedType,
				Name: ipAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: ipAPIName,
					DestField:   ipAPIName,
					IsPlural:    false,
					Model:       ipView,
				},
			},
		},
	},
}
var (
	ipNakedType = meta.Static(naked.IPAddress{})

	ipView = &dsl.Model{
		Name:      ipAPIName,
		NakedType: ipNakedType,
		Fields: []*dsl.FieldDesc{
			fields.HostName(),
			fields.IPAddress(),
			fields.InterfaceID(),
			fields.SubnetID(),
			// Note: InterfaceとSubnetはIDにのみ対応。その他のフィールドは今後必要になったら対応を検討する。
		},
	}

	ipUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(ipAPIName),
		NakedType: ipNakedType,
		Fields: []*dsl.FieldDesc{
			fields.UserIPAddress(),
		},
	}
)
