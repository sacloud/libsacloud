package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var switchAPI = &schema.Resource{
	Name:       "Switch",
	PathName:   "switch",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(switchNakedType, findParameter, switchView),

			// create
			r.DefineOperationCreate(switchNakedType, switchCreateParam, switchView),

			// read
			r.DefineOperationRead(switchNakedType, switchView),

			// update
			r.DefineOperationUpdate(switchNakedType, switchUpdateParam, switchView),

			// delete
			r.DefineOperationDelete(),

			// connect from bridge
			r.DefineSimpleOperation("ConnectToBridge", http.MethodPut, "to/bridge/{{.bridgeID}}",
				&schema.Argument{
					Name: "bridgeID",
					Type: meta.TypeID,
				},
			),

			// disconnect from bridge
			r.DefineSimpleOperation("DisconnectFromBridge", http.MethodDelete, "to/bridge/"),
		}
	},
}

var (
	switchNakedType = meta.Static(naked.Switch{})

	switchView = &schema.Model{
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

	switchCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	switchUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.UserSubnetNetworkMaskLen(),
			fields.UserSubnetDefaultRoute(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
