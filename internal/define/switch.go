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
			r.DefineOperation("ConnectToBridge").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("to/bridge/{{.bridgeID}}")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.SimpleArgument{
					Name: "bridgeID",
					Type: meta.TypeID,
				}),

			// disconnect from bridge
			r.DefineOperation("DisconnectFromBridge").
				Method(http.MethodDelete).
				PathFormat(schema.IDAndSuffixPathFormat("to/bridge")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID),
		)
	})
}
