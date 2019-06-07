package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.CDROM{})

	cdrom := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.DisplayOrder(),
			fields.Tags(),
			fields.Availability(),
			fields.Scope(),
			fields.Storage(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	cdromAPI := &schema.Resource{
		Name:       "CDROM",
		PathName:   "cdrom",
		PathSuffix: schema.CloudAPISuffix,
	}
	cdromAPI.Operations = []*schema.Operation{
		// find
		cdromAPI.DefineOperationFind(nakedType, findParameter, cdrom),

		// create
		cdromAPI.DefineOperationCreate(nakedType, createParam, cdrom).
			ResultFromEnvelope(models.ftpServer(), &schema.EnvelopePayloadDesc{
				PayloadName: models.ftpServer().Name,
				PayloadType: meta.Static(naked.OpeningFTPServer{}),
			}),

		// read
		cdromAPI.DefineOperationRead(nakedType, cdrom),

		// update
		cdromAPI.DefineOperationUpdate(nakedType, updateParam, cdrom),

		// delete
		cdromAPI.DefineOperationDelete(),

		// openFTP
		cdromAPI.DefineOperationOpenFTP(models.ftpServerOpenParameter(), models.ftpServer()),

		// closeFTP
		cdromAPI.DefineOperationCloseFTP(),
	}
	Resources.Def(cdromAPI)
}
