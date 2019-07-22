package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	zoneAPIName     = "Zone"
	zoneAPIPathName = "zone"
)

var zoneAPI = &dsl.Resource{
	Name:       zoneAPIName,
	PathName:   zoneAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		ops.Find(zoneAPIName, zoneNakedType, findParameter, zoneView),
		ops.Read(zoneAPIName, zoneNakedType, zoneView),
	},
}

var (
	zoneNakedType = meta.Static(naked.Zone{})
	zoneView      = &dsl.Model{
		Name:      zoneAPIName,
		NakedType: zoneNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.DisplayOrder(),
			fields.IsDummy(),
			fields.VNCProxy(),
			fields.FTPServer(),
			fields.Region(),
		},
	}
)
