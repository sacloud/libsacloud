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

	Resources.DefineWith("Bridge", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, bridge),

			// create
			r.DefineOperationCreate(nakedType, createParam, bridge),

			// read
			r.DefineOperationRead(nakedType, bridge),

			// update
			r.DefineOperationUpdate(nakedType, updateParam, bridge),

			// delete
			r.DefineOperationDelete(),
		)
	})
}
