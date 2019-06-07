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

	routerAPI := &schema.Resource{
		Name:       "Internet",
		PathName:   "internet",
		PathSuffix: schema.CloudAPISuffix,
	}
	routerAPI.Operations = []*schema.Operation{
		// find
		routerAPI.DefineOperationFind(nakedType, findParameter, internet),

		// create
		routerAPI.DefineOperationCreate(nakedType, createParam, internet),

		// read
		routerAPI.DefineOperationRead(nakedType, internet),

		// update
		routerAPI.DefineOperationUpdate(nakedType, updateParam, internet),

		// delete
		routerAPI.DefineOperationDelete(),

		// UpdateBandWidth
		routerAPI.DefineOperation("UpdateBandWidth").
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
		routerAPI.DefineOperation("AddSubnet").
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
		routerAPI.DefineOperation("UpdateSubnet").
			Method(http.MethodPut).
			PathFormat(schema.IDAndSuffixPathFormat("subnet/{{.subnetID}}")).
			Argument(schema.ArgumentZone).
			Argument(schema.ArgumentID).
			Argument(&schema.Argument{
				Name: "subnetID",
				Type: meta.TypeID,
			}).
			PassthroughModelArgumentWithEnvelope("param", updateSubnetParam).
			ResultFromEnvelope(models.internetSubnetOperationResult(), &schema.EnvelopePayloadDesc{
				PayloadType: meta.Static(naked.Subnet{}),
				PayloadName: "Subnet",
			}),

		// DeleteSubnet
		routerAPI.DefineSimpleOperation("DeleteSubnet", http.MethodDelete, "subnet/{{.subnetID}}",
			&schema.Argument{
				Name: "subnetID",
				Type: meta.TypeID,
			},
		),

		// monitor
		routerAPI.DefineOperationMonitor(monitorParameter, monitors.routerModel()),

		// TODO IPv6関連は後回し
	}

	Resources.Def(routerAPI)
}
