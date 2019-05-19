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

	Resources.DefineWith("CDROM", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, cdrom),

			// create
			r.DefineOperationCreate(nakedType, createParam, cdrom).
				ResultFromEnvelope(models.ftpServer(), &schema.EnvelopePayloadDesc{
					PayloadName: models.ftpServer().Name,
					PayloadType: meta.Static(naked.OpeningFTPServer{}),
				}),

			// read
			r.DefineOperationRead(nakedType, cdrom),

			// update
			r.DefineOperationUpdate(nakedType, updateParam, cdrom),

			// delete
			r.DefineOperationDelete(),

			// openFTP
			r.DefineOperationOpenFTP(models.ftpServerOpenParameter(), models.ftpServer()),

			// closeFTP
			r.DefineOperationCloseFTP(),
		)
	})
}
