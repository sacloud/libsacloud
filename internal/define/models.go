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
		Name: "OpenFTPRequest",
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

func (m *modelsDef) diskEdit() *schema.Model {

	sshKeyFields := []*schema.FieldDesc{
		{
			Name: "ID",
			Type: meta.TypeID,
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "PublicKey",
			Type: meta.TypeString,
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
	}

	noteFields := []*schema.FieldDesc{
		{
			Name: "ID",
			Type: meta.TypeID,
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "Variables",
			Type: meta.Static(map[string]interface{}{}),
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
	}

	userSubnetFdields := []*schema.FieldDesc{
		{
			Name: "DefaultRoute",
			Type: meta.TypeString,
			Tags: &schema.FieldTags{
				MapConv: ",omitempty",
				JSON:    ",omitempty",
			},
		},
		{
			Name: "NetworkMaskLen",
			Type: meta.TypeInt,
			Tags: &schema.FieldTags{
				MapConv:  ",omitempty",
				Validate: "min=0,max=32",
				JSON:     ",omitempty",
			},
		},
	}

	return &schema.Model{
		Name:      "DiskEditRequest",
		NakedType: meta.Static(naked.DiskEdit{}),
		Fields: []*schema.FieldDesc{
			{
				Name: "Password",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "SSHKey",
				Type: &schema.Model{
					Name:   "DiskEditSSHKey",
					Fields: sshKeyFields,
				},
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "SSHKeys",
				Type: &schema.Model{
					Name:    "DiskEditSSHKey",
					IsArray: true,
					Fields:  sshKeyFields,
				},
				Tags: &schema.FieldTags{
					MapConv: "[]SSHKeys,omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "DisablePWAuth",
				Type: meta.TypeFlag,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "EnableDHCP",
				Type: meta.TypeFlag,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "ChangePartitionUUID",
				Type: meta.TypeFlag,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "HostName",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "Notes",
				Type: &schema.Model{
					Name:    "DiskEditNote",
					IsArray: true,
					Fields:  noteFields,
				},
				Tags: &schema.FieldTags{
					MapConv: ",omitempty,recursive",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "UserIPAddress",
				Type: meta.TypeString,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
			{
				Name: "UserSubnet",
				Type: &schema.Model{
					Name:   "DiskEditUserSubnet",
					Fields: userSubnetFdields,
				},
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}
}
