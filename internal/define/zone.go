package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {

	nakedType := meta.Static(naked.Zone{})
	zone := &schema.Model{
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

	Resources.Define("Zone").SetIsGlobal(true).
		OperationFind(nakedType, findParameter, zone).
		OperationRead(nakedType, zone)
}
