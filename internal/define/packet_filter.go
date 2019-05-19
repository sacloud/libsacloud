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

	Resources.DefineWith("PacketFilter", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, pf),

			// create
			r.DefineOperationCreate(nakedType, createParam, pf),

			// read
			r.DefineOperationRead(nakedType, pf),

			// update
			r.DefineOperationUpdate(nakedType, updateParam, pf),

			// delete
			r.DefineOperationDelete(),
		)
	}).PathName("packetfilter")
}
