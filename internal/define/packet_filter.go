package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	packetFilterAPIName     = "PacketFilter"
	packetFilterAPIPathName = "packetfilter"
)

var packetFilterAPI = &schema.Resource{
	Name:       packetFilterAPIName,
	PathName:   packetFilterAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.Find(packetFilterAPIName, packetFilterNakedType, findParameter, packetFilterView),

		// create
		ops.Create(packetFilterAPIName, packetFilterNakedType, packetFilterCreateParam, packetFilterView),

		// read
		ops.Read(packetFilterAPIName, packetFilterNakedType, packetFilterView),

		// update
		ops.Update(packetFilterAPIName, packetFilterNakedType, packetFilterUpdateParam, packetFilterView),

		// delete
		ops.Delete(packetFilterAPIName),
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
