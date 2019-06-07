package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.Note{})

	note := &schema.Model{
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

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}

	noteAPI := &schema.Resource{
		Name:       "Note",
		PathName:   "note",
		PathSuffix: schema.CloudAPISuffix,
		IsGlobal:   true,
	}
	noteAPI.Operations = []*schema.Operation{
		// find
		noteAPI.DefineOperationFind(nakedType, findParameter, note),

		// create
		noteAPI.DefineOperationCreate(nakedType, createParam, note),

		// read
		noteAPI.DefineOperationRead(nakedType, note),

		// update
		noteAPI.DefineOperationUpdate(nakedType, updateParam, note),

		// delete
		noteAPI.DefineOperationDelete(),
	}
	Resources.Def(noteAPI)
}
