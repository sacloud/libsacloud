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

	Resources.DefineWith("Switch", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, sw),

			// create
			r.DefineOperationCreate(nakedType, createParam, sw),

			// read
			r.DefineOperationRead(nakedType, sw),

			// update
			r.DefineOperationUpdate(nakedType, updateParam, sw),

			// delete
			r.DefineOperationDelete(),

			// connect from bridge
			r.DefineSimpleOperation("ConnectToBridge", http.MethodPut, "to/bridge/{{.bridgeID}}",
				&schema.SimpleArgument{
					Name: "bridgeID",
					Type: meta.TypeID,
				},
			),

			// disconnect from bridge
			r.DefineSimpleOperation("DisconnectFromBridge", http.MethodDelete, "to/bridge/"),
		)
	})
}
