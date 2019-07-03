package define

import (
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var bridgeAPI = &schema.Resource{
	Name:       "Bridge",
	PathName:   "bridge",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(bridgeNakedType, findParameter, bridgeView),

			// create
			r.DefineOperationCreate(bridgeNakedType, bridgeCreateParam, bridgeView),

			// read
			r.DefineOperationRead(bridgeNakedType, bridgeView),

			// update
			r.DefineOperationUpdate(bridgeNakedType, bridgeUpdateParam, bridgeView),

			// delete
			r.DefineOperationDelete(),
		}
	},
}

var (
	bridgeNakedType = meta.Static(naked.Bridge{})

	bridgeView = &schema.Model{
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

	bridgeCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
		},
	}

	bridgeUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
		},
	}
)
