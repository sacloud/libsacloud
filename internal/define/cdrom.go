package define

import (
	"net/http"

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
			fields.StorageClass(),
			fields.Storage(),
			fields.Icon(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	ftpServer := &schema.Model{
		Name:      "FTPServer",
		NakedType: meta.Static(naked.OpeningFTPServer{}),
		Fields: []*schema.FieldDesc{
			fields.HostName(),
			fields.IPAddress(),
			fields.User(),
			fields.Password(),
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

	openFTPParam := &schema.Model{
		Name: "OpenFTPParam",
		Fields: []*schema.FieldDesc{
			{
				Name: "ChangePassword",
				Type: meta.TypeFlag,
			},
		},
	}

	Resources.DefineWith("CDROM", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, cdrom),

			// create
			r.DefineOperationCreate(nakedType, createParam, cdrom).
				ResultFromEnvelope(ftpServer, &schema.EnvelopePayloadDesc{
					PayloadName: ftpServer.Name,
					PayloadType: meta.Static(naked.OpeningFTPServer{}),
				}),

			// read
			r.DefineOperationRead(nakedType, cdrom),

			// update
			r.DefineOperationUpdate(nakedType, updateParam, cdrom),

			// delete
			r.DefineOperationDelete(),

			// openFTP
			r.DefineOperation("OpenFTP").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("ftp")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				PassthroughArgumentToPayload("openOption", openFTPParam).
				ResultFromEnvelope(ftpServer, &schema.EnvelopePayloadDesc{
					PayloadName: ftpServer.Name,
					PayloadType: meta.Static(naked.OpeningFTPServer{}),
				}),

			// closeFTP
			r.DefineOperation("CloseFTP").
				Method(http.MethodDelete).
				PathFormat(schema.IDAndSuffixPathFormat("ftp")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID),
		)
	})
}
