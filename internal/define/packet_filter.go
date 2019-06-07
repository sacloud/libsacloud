package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.PacketFilter{})

	pf := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.RequiredHostVersion(),
			fields.PacketFilterExpressions(),
			fields.ExpressionHash(),
			fields.CreatedAt(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PacketFilterExpressions(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PacketFilterExpressions(),
		},
	}

	pfAPI := &schema.Resource{
		Name:       "PacketFilter",
		PathName:   "packetfilter",
		PathSuffix: schema.CloudAPISuffix,
	}

	pfAPI.Operations = []*schema.Operation{
		// find
		pfAPI.DefineOperationFind(nakedType, findParameter, pf),

		// create
		pfAPI.DefineOperationCreate(nakedType, createParam, pf),

		// read
		pfAPI.DefineOperationRead(nakedType, pf),

		// update
		pfAPI.DefineOperationUpdate(nakedType, updateParam, pf),

		// delete
		pfAPI.DefineOperationDelete(),
	}
	Resources.Def(pfAPI)
}
