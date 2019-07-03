package define

import (
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var packetFilterAPI = &schema.Resource{
	Name:       "PacketFilter",
	PathName:   "packetfilter",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(packetFilterNakedType, findParameter, packetFilterView),

			// create
			r.DefineOperationCreate(packetFilterNakedType, packetFilterCreateParam, packetFilterView),

			// read
			r.DefineOperationRead(packetFilterNakedType, packetFilterView),

			// update
			r.DefineOperationUpdate(packetFilterNakedType, packetFilterUpdateParam, packetFilterView),

			// delete
			r.DefineOperationDelete(),
		}
	},
}

var (
	packetFilterNakedType = meta.Static(naked.PacketFilter{})

	packetFilterView = &schema.Model{
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

	packetFilterCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PacketFilterExpressions(),
		},
	}

	packetFilterUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PacketFilterExpressions(),
		},
	}
)
