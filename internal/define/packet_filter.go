package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	packetFilterAPIName     = "PacketFilter"
	packetFilterAPIPathName = "packetfilter"
)

var packetFilterAPI = &dsl.Resource{
	Name:       packetFilterAPIName,
	PathName:   packetFilterAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
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

	packetFilterView = &dsl.Model{
		Name:      packetFilterAPIName,
		NakedType: packetFilterNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.RequiredHostVersion(),
			fields.PacketFilterExpressions(),
			fields.ExpressionHash(),
			fields.CreatedAt(),
		},
	}

	packetFilterCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(packetFilterAPIName),
		NakedType: packetFilterNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PacketFilterExpressions(),
		},
	}

	packetFilterUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(packetFilterAPIName),
		NakedType: packetFilterNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PacketFilterExpressions(),
		},
	}
)
