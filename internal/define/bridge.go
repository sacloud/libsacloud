package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	bridgeAPIName     = "Bridge"
	bridgeAPIPathName = "bridge"
)

var bridgeAPI = &schema.Resource{
	Name:       bridgeAPIName,
	PathName:   bridgeAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.Find(bridgeAPIName, bridgeNakedType, findParameter, bridgeView),

		// create
		ops.Create(bridgeAPIName, bridgeNakedType, bridgeCreateParam, bridgeView),

		// read
		ops.Read(bridgeAPIName, bridgeNakedType, bridgeView),

		// update
		ops.Update(bridgeAPIName, bridgeNakedType, bridgeUpdateParam, bridgeView),

		// delete
		ops.Delete(bridgeAPIName),
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
