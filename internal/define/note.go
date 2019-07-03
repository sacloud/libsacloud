package define

import (
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var noteAPI = &schema.Resource{
	Name:       "Note",
	PathName:   "note",
	PathSuffix: schema.CloudAPISuffix,
	IsGlobal:   true,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(noteNakedType, findParameter, noteView),

			// create
			r.DefineOperationCreate(noteNakedType, noteCreateParam, noteView),

			// read
			r.DefineOperationRead(noteNakedType, noteView),

			// update
			r.DefineOperationUpdate(noteNakedType, noteUpdateParam, noteView),

			// delete
			r.DefineOperationDelete(),
		}
	},
}

var (
	noteNakedType = meta.Static(naked.Note{})

	noteView = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Scope(),
			fields.NoteClass(),
			fields.NoteContent(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	noteCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}

	noteUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}
)
