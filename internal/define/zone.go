package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	zoneAPIName     = "Zone"
	zoneAPIPathName = "zone"
)

var zoneAPI = &schema.Resource{
	Name:       zoneAPIName,
	PathName:   zoneAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	IsGlobal:   true,
	Operations: schema.Operations{
		ops.Find(zoneAPIName, zoneNakedType, findParameter, zoneView),
		ops.Read(zoneAPIName, zoneNakedType, zoneView),
	},
}

var (
	zoneNakedType = meta.Static(naked.Zone{})
	zoneView      = &schema.Model{
		Fields: []*schema.FieldDesc{
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
