package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	iconAPIName     = "Icon"
	iconAPIPathName = "icon"
)

// iconAPI アイコンAPI
//
// Note: libsacloudでは画像データ取得(GET /icon/:id?Size=[small|medium|large])はサポートしない。
var iconAPI = &dsl.Resource{
	Name:       iconAPIName,
	PathName:   iconAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.Find(iconAPIName, iconNakedType, findParameter, iconView),

		// create
		ops.Create(iconAPIName, iconNakedType, iconCreateParam, iconView),

		// read
		ops.Read(iconAPIName, iconNakedType, iconView),

		// update
		ops.Update(iconAPIName, iconNakedType, iconUpdateParam, iconView),

		// delete
		ops.Delete(iconAPIName),
	},
}

var (
	iconNakedType = meta.Static(naked.Icon{})

	iconView = &dsl.Model{
		Name:      iconAPIName,
		NakedType: iconNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Tags(),
			fields.Availability(),
			fields.Scope(),
			fields.IconURL(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}
	iconCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(iconAPIName),
		NakedType: iconNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconImage(),
		},
	}

	iconUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(iconAPIName),
		NakedType: iconNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Tags(),
		},
	}
)
