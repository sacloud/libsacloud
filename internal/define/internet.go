package define

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.Internet{})

	internet := models.internetModel()

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.NetworkMaskLen(),
			fields.BandWidthMbps(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	updateBandWidthParam := &schema.Model{
		Name:      "InternetUpdateBandWidthRequest",
		NakedType: nakedType,
		Fields: []*schema.FieldDesc{
			fields.BandWidthMbps(),
		},
	}

	addSubnetParam := &schema.Model{
		Name:      "InternetAddSubnetRequest",
		NakedType: meta.Static(naked.SubnetOperationRequest{}),
		Fields: []*schema.FieldDesc{
			fields.NetworkMaskLen(),
			fields.NextHop(),
		},
	}
	updateSubnetParam := &schema.Model{
		Name:      "InternetUpdateSubnetRequest",
		NakedType: meta.Static(naked.SubnetOperationRequest{}),
		Fields: []*schema.FieldDesc{
			fields.NextHop(),
		},
	}

	Resources.DefineWith("Internet", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, internet),

			// create
			r.DefineOperationCreate(nakedType, createParam, internet),

			// read
			r.DefineOperationRead(nakedType, internet),

			// update
			r.DefineOperationUpdate(nakedType, updateParam, internet),

			// delete
			r.DefineOperationDelete(),

			// UpdateBandWidth
			r.DefineOperation("UpdateBandWidth").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("bandwidth")).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: nakedType,
					PayloadName: "Internet",
				}).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				MappableArgument("param", updateBandWidthParam).
				ResultFromEnvelope(internet, &schema.EnvelopePayloadDesc{
					PayloadType: nakedType,
					PayloadName: "Internet",
				}),

			// AddSubnet
			r.DefineOperation("AddSubnet").
				Method(http.MethodPost).
				PathFormat(schema.IDAndSuffixPathFormat("subnet")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				PassthroughModelArgumentWithEnvelope("param", addSubnetParam).
				ResultFromEnvelope(models.internetSubnetOperationResult(), &schema.EnvelopePayloadDesc{
					PayloadType: meta.Static(naked.Subnet{}),
					PayloadName: "Subnet",
				}),

			// UpdateSubnet
			r.DefineOperation("UpdateSubnet").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("subnet/{{.subnetID}}")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.SimpleArgument{
					Name: "subnetID",
					Type: meta.TypeID,
				}).
				PassthroughModelArgumentWithEnvelope("param", updateSubnetParam).
				ResultFromEnvelope(models.internetSubnetOperationResult(), &schema.EnvelopePayloadDesc{
					PayloadType: meta.Static(naked.Subnet{}),
					PayloadName: "Subnet",
				}),

			// DeleteSubnet
			r.DefineSimpleOperation("DeleteSubnet", http.MethodDelete, "subnet/{{.subnetID}}",
				&schema.SimpleArgument{
					Name: "subnetID",
					Type: meta.TypeID,
				},
			),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.routerModel()),
		)

		// TODO IPv6関連は後回し
	})
}
