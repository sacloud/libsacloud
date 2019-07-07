package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var internetAPI = &schema.Resource{
	Name:       "Internet",
	PathName:   "internet",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{

			// find
			r.DefineOperationFind(internetNakedType, findParameter, internetView),

			// create
			r.DefineOperationCreate(internetNakedType, internetCreateParam, internetView),

			// read
			r.DefineOperationRead(internetNakedType, internetView),

			// update
			r.DefineOperationUpdate(internetNakedType, internetUpdateParam, internetView),

			// delete
			r.DefineOperationDelete(),

			// UpdateBandWidth
			{
				Resource:   r,
				Name:       "UpdateBandWidth",
				PathFormat: schema.IDAndSuffixPathFormat("bandwidth"),
				Method:     http.MethodPut,
				RequestEnvelope: schema.RequestEnvelope(
					&schema.EnvelopePayloadDesc{
						Type: internetNakedType,
						Name: "Internet",
					},
				),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.MappableArgument("param", internetUpdateBandWidthParam, "Internet"),
				},
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Type: internetNakedType,
					Name: "Internet",
				}),
				Results: schema.Results{
					{
						SourceField: "Internet",
						DestField:   "Internet",
						IsPlural:    false,
						Model:       internetView,
					},
				},
			},

			// AddSubnet
			{
				Resource:        r,
				Name:            "AddSubnet",
				PathFormat:      schema.IDAndSuffixPathFormat("subnet"),
				Method:          http.MethodPost,
				RequestEnvelope: schema.RequestEnvelopeFromModel(internetAddSubnetParam),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.PassthroughModelArgument("param", internetAddSubnetParam),
				},
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Type: meta.Static(naked.Subnet{}),
					Name: "Subnet",
				}),
				Results: schema.Results{
					{
						SourceField: "Subnet",
						DestField:   "Subnet",
						IsPlural:    false,
						Model:       models.internetSubnetOperationResult(),
					},
				},
			},

			// UpdateSubnet
			{
				Resource:        r,
				Name:            "UpdateSubnet",
				PathFormat:      schema.IDAndSuffixPathFormat("subnet/{{.subnetID}}"),
				Method:          http.MethodPut,
				RequestEnvelope: schema.RequestEnvelopeFromModel(internetUpdateSubnetParam),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					{
						Name: "subnetID",
						Type: meta.TypeID,
					},
					schema.PassthroughModelArgument("param", internetUpdateSubnetParam),
				},
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Type: meta.Static(naked.Subnet{}),
					Name: "Subnet",
				}),
				Results: schema.Results{
					{
						SourceField: "Subnet",
						DestField:   "Subnet",
						IsPlural:    false,
						Model:       models.internetSubnetOperationResult(),
					},
				},
			},

			// DeleteSubnet
			r.DefineSimpleOperation("DeleteSubnet", http.MethodDelete, "subnet/{{.subnetID}}",
				&schema.Argument{
					Name: "subnetID",
					Type: meta.TypeID,
				},
			),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.routerModel()),

			// TODO IPv6関連は後回し
		}
	},
}
var (
	internetNakedType = meta.Static(naked.Internet{})

	internetView = models.internetModel()

	internetCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.NetworkMaskLen(),
			fields.BandWidthMbps(),
		},
	}

	internetUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	internetUpdateBandWidthParam = &schema.Model{
		Name:      "InternetUpdateBandWidthRequest",
		NakedType: internetNakedType,
		Fields: []*schema.FieldDesc{
			fields.BandWidthMbps(),
		},
	}

	internetAddSubnetParam = &schema.Model{
		Name:      "InternetAddSubnetRequest",
		NakedType: meta.Static(naked.SubnetOperationRequest{}),
		Fields: []*schema.FieldDesc{
			fields.NetworkMaskLen(),
			fields.NextHop(),
		},
	}
	internetUpdateSubnetParam = &schema.Model{
		Name:      "InternetUpdateSubnetRequest",
		NakedType: meta.Static(naked.SubnetOperationRequest{}),
		Fields: []*schema.FieldDesc{
			fields.NextHop(),
		},
	}
)
