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
	internetAPIName     = "Internet"
	internetAPIPathName = "internet"
)

var internetAPI = &dsl.Resource{
	Name:       internetAPIName,
	PathName:   internetAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{

		// find
		ops.Find(internetAPIName, internetNakedType, findParameter, internetView),

		// create
		ops.Create(internetAPIName, internetNakedType, internetCreateParam, internetView),

		// read
		ops.Read(internetAPIName, internetNakedType, internetView),

		// update
		ops.Update(internetAPIName, internetNakedType, internetUpdateParam, internetView),

		// delete
		ops.Delete(internetAPIName),

		// UpdateBandWidth
		{
			ResourceName: internetAPIName,
			Name:         "UpdateBandWidth",
			PathFormat:   dsl.IDAndSuffixPathFormat("bandwidth"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: internetNakedType,
					Name: "Internet",
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", internetUpdateBandWidthParam, "Internet"),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: internetNakedType,
				Name: "Internet",
			}),
			Results: dsl.Results{
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
			ResourceName:    internetAPIName,
			Name:            "AddSubnet",
			PathFormat:      dsl.IDAndSuffixPathFormat("subnet"),
			Method:          http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelopeFromModel(internetAddSubnetParam),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.PassthroughModelArgument("param", internetAddSubnetParam),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.Subnet{}),
				Name: "Subnet",
			}),
			Results: dsl.Results{
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
			ResourceName:    internetAPIName,
			Name:            "UpdateSubnet",
			PathFormat:      dsl.IDAndSuffixPathFormat("subnet/{{.subnetID}}"),
			Method:          http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelopeFromModel(internetUpdateSubnetParam),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "subnetID",
					Type: meta.TypeID,
				},
				dsl.PassthroughModelArgument("param", internetUpdateSubnetParam),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.Subnet{}),
				Name: "Subnet",
			}),
			Results: dsl.Results{
				{
					SourceField: "Subnet",
					DestField:   "Subnet",
					IsPlural:    false,
					Model:       models.internetSubnetOperationResult(),
				},
			},
		},

		// DeleteSubnet
		ops.WithIDAction(internetAPIName, "DeleteSubnet", http.MethodDelete, "subnet/{{.subnetID}}",
			&dsl.Argument{
				Name: "subnetID",
				Type: meta.TypeID,
			},
		),

		// monitor
		ops.Monitor(internetAPIName, monitorParameter, monitors.routerModel()),

		// ipv6
		{
			ResourceName: internetAPIName,
			Name:         "EnableIPv6",
			PathFormat:   dsl.IDAndSuffixPathFormat("ipv6net"),
			Method:       http.MethodPost,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: ipv6netNakedType,
				Name: ipv6netAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: ipv6netAPIName,
					DestField:   ipv6netAPIName,
					IsPlural:    false,
					Model:       models.switchIPv6NetModel(),
				},
			},
		},

		ops.WithIDAction(internetAPIName, "DisableIPv6", http.MethodDelete, "ipv6net/{{.ipv6netID}}",
			&dsl.Argument{
				Name: "ipv6netID",
				Type: meta.TypeID,
			},
		),
	},
}
var (
	internetNakedType = meta.Static(naked.Internet{})

	internetView = models.internetModel()

	internetCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(internetAPIName),
		NakedType: internetNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.NetworkMaskLen(),
			fields.BandWidthMbps(),
		},
	}

	internetUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(internetAPIName),
		NakedType: internetNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	internetUpdateBandWidthParam = &dsl.Model{
		Name:      "InternetUpdateBandWidthRequest",
		NakedType: internetNakedType,
		Fields: []*dsl.FieldDesc{
			fields.BandWidthMbps(),
		},
	}

	internetAddSubnetParam = &dsl.Model{
		Name:      "InternetAddSubnetRequest",
		NakedType: meta.Static(naked.SubnetOperationRequest{}),
		Fields: []*dsl.FieldDesc{
			fields.NetworkMaskLen(),
			fields.NextHop(),
		},
	}
	internetUpdateSubnetParam = &dsl.Model{
		Name:      "InternetUpdateSubnetRequest",
		NakedType: meta.Static(naked.SubnetOperationRequest{}),
		Fields: []*dsl.FieldDesc{
			fields.NextHop(),
		},
	}
)
