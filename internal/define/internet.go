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
			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource:   r,
					Name:       "UpdateBandWidth",
					PathFormat: schema.IDAndSuffixPathFormat("bandwidth"),
					Method:     http.MethodPut,
					Arguments: schema.Arguments{
						schema.ArgumentZone,
						schema.ArgumentID,
						schema.MappableArgument("param", internetUpdateBandWidthParam, "Internet"),
					},
				}
				o.ResponseEnvelope = schema.ResultFromEnvelope(o, internetView, &schema.EnvelopePayloadDesc{
					PayloadType: internetNakedType,
					PayloadName: "Internet",
				}, "")
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: internetNakedType,
						PayloadName: "Internet",
					},
				)
				return o
			}(),

			// AddSubnet
			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource:   r,
					Name:       "AddSubnet",
					PathFormat: schema.IDAndSuffixPathFormat("subnet"),
					Method:     http.MethodPost,
				}
				o.ResponseEnvelope = schema.ResultFromEnvelope(o, models.internetSubnetOperationResult(), &schema.EnvelopePayloadDesc{
					PayloadType: meta.Static(naked.Subnet{}),
					PayloadName: "Subnet",
				}, "Subnet")
				o.Arguments = schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.PassthroughModelArgumentWithEnvelope(o, "param", internetAddSubnetParam),
				}
				return o
			}(),

			// UpdateSubnet
			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource:   r,
					Name:       "UpdateSubnet",
					PathFormat: schema.IDAndSuffixPathFormat("subnet/{{.subnetID}}"),
					Method:     http.MethodPut,
				}
				o.ResponseEnvelope = schema.ResultFromEnvelope(o, models.internetSubnetOperationResult(), &schema.EnvelopePayloadDesc{
					PayloadType: meta.Static(naked.Subnet{}),
					PayloadName: "Subnet",
				}, "Subnet")
				o.Arguments = schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					{
						Name: "subnetID",
						Type: meta.TypeID,
					},
					schema.PassthroughModelArgumentWithEnvelope(o, "param", internetUpdateSubnetParam),
				}
				return o
			}(),

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
