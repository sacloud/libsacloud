package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	bridgeAPIName     = "Bridge"
	bridgeAPIPathName = "bridge"
)

var bridgeAPI = &dsl.Resource{
	Name:       bridgeAPIName,
	PathName:   bridgeAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(bridgeAPIName, bridgeNakedType, findParameter, bridgeView, nil),

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

	bridgeView = &dsl.Model{
		Name:      bridgeAPIName,
		NakedType: bridgeNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.CreatedAt(),
			fields.Region(),
			fields.BridgeInfo(),
			fields.SwitchInZone(),
		},
	}

	bridgeCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(bridgeAPIName),
		NakedType: bridgeNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
		},
	}

	bridgeUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(bridgeAPIName),
		NakedType: bridgeNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
		},
	}
)
