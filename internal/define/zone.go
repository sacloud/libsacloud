package define

import (
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var zoneAPI = &schema.Resource{
	Name:       "Zone",
	PathName:   "zone",
	PathSuffix: schema.CloudAPISuffix,
	IsGlobal:   true,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			r.DefineOperationFind(zoneNakedType, findParameter, zoneView),
			r.DefineOperationRead(zoneNakedType, zoneView),
		}
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
