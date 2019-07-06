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
				o := r.DefineOperation("UpdateBandWidth").
					ResultFromEnvelope(internetView, &schema.EnvelopePayloadDesc{
						PayloadType: internetNakedType,
						PayloadName: "Internet",
					}, "")
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: internetNakedType,
						PayloadName: "Internet",
					},
				)
				o.Arguments = schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.MappableArgument(o, "param", internetUpdateBandWidthParam),
				}
				o.PathFormat = schema.IDAndSuffixPathFormat("bandwidth")
				o.Method = http.MethodPut
				return o
			}(),

			// AddSubnet
			// TODO あとで直す
			func() *schema.Operation {
				o := r.DefineOperation("AddSubnet").
					ResultFromEnvelope(models.internetSubnetOperationResult(), &schema.EnvelopePayloadDesc{
						PayloadType: meta.Static(naked.Subnet{}),
						PayloadName: "Subnet",
					}, "Subnet")
				o.Arguments = schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.PassthroughModelArgumentWithEnvelope(o, "param", internetAddSubnetParam),
				}
				o.PathFormat = schema.IDAndSuffixPathFormat("subnet")
				o.Method = http.MethodPost
				return o
			}(),

			// UpdateSubnet
			// TODO あとで直す
			func() *schema.Operation {
				o := r.DefineOperation("UpdateSubnet").
					ResultFromEnvelope(models.internetSubnetOperationResult(), &schema.EnvelopePayloadDesc{
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
				o.PathFormat = schema.IDAndSuffixPathFormat("subnet/{{.subnetID}}")
				o.Method = http.MethodPut
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
