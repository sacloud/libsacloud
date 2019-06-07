package define

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.Switch{})

	sw := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.Scope(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			{
				Name: "Subnets",
				Type: models.switchSubnet(),
				Tags: &schema.FieldTags{
					MapConv: "[]Subnets,omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			fields.BridgeID(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
	swAPI := &schema.Resource{
		Name:       "Switch",
		PathName:   "switch",
		PathSuffix: schema.CloudAPISuffix,
	}
	swAPI.Operations = []*schema.Operation{
		// find
		swAPI.DefineOperationFind(nakedType, findParameter, sw),

		// create
		swAPI.DefineOperationCreate(nakedType, createParam, sw),

		// read
		swAPI.DefineOperationRead(nakedType, sw),

		// update
		swAPI.DefineOperationUpdate(nakedType, updateParam, sw),

		// delete
		swAPI.DefineOperationDelete(),

		// connect from bridge
		swAPI.DefineSimpleOperation("ConnectToBridge", http.MethodPut, "to/bridge/{{.bridgeID}}",
			&schema.Argument{
				Name: "bridgeID",
				Type: meta.TypeID,
			},
		),

		// disconnect from bridge
		swAPI.DefineSimpleOperation("DisconnectFromBridge", http.MethodDelete, "to/bridge/"),
	}
	Resources.Def(swAPI)
}
