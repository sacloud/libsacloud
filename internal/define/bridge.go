package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.Bridge{})

	bridge := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.CreatedAt(),
			fields.Region(),
			fields.BridgeInfo(),
			fields.SwitchInZone(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
		},
	}

	bridgeAPI := &schema.Resource{
		Name:       "Bridge",
		PathName:   "bridge",
		PathSuffix: schema.CloudAPISuffix,
	}
	bridgeAPI.Operations = []*schema.Operation{
		// find
		bridgeAPI.DefineOperationFind(nakedType, findParameter, bridge),

		// create
		bridgeAPI.DefineOperationCreate(nakedType, createParam, bridge),

		// read
		bridgeAPI.DefineOperationRead(nakedType, bridge),

		// update
		bridgeAPI.DefineOperationUpdate(nakedType, updateParam, bridge),

		// delete
		bridgeAPI.DefineOperationDelete(),
	}
	Resources.Def(bridgeAPI)
}
