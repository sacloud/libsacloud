package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

type modelsDef struct{}

var models = &modelsDef{}

func (m *modelsDef) ftpServerOpenParameter() *schema.Model {
	return &schema.Model{
		Name: "OpenFTPParam",
		Fields: []*schema.FieldDesc{
			{
				Name: "ChangePassword",
				Type: meta.TypeFlag,
			},
		},
	}
}

func (m *modelsDef) ftpServer() *schema.Model {
	return &schema.Model{
		Name:      "FTPServer",
		NakedType: meta.Static(naked.OpeningFTPServer{}),
		Fields: []*schema.FieldDesc{
			fields.HostName(),
			fields.IPAddress(),
			fields.User(),
			fields.Password(),
		},
	}
}
