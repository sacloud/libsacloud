package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	noteAPIName     = "Note"
	noteAPIPathName = "note"
)

var noteAPI = &schema.Resource{
	Name:       noteAPIName,
	PathName:   noteAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	IsGlobal:   true,
	Operations: schema.Operations{
		// find
		ops.Find(noteAPIName, noteNakedType, findParameter, noteView),

		// create
		ops.Create(noteAPIName, noteNakedType, noteCreateParam, noteView),

		// read
		ops.Read(noteAPIName, noteNakedType, noteView),

		// update
		ops.Update(noteAPIName, noteNakedType, noteUpdateParam, noteView),

		// delete
		ops.Delete(noteAPIName),
	},
}

var (
	noteNakedType = meta.Static(naked.Note{})

	noteView = &schema.Model{
		Name:      noteAPIName,
		NakedType: noteNakedType,
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
		Name:      names.CreateParameterName(noteAPIName),
		NakedType: noteNakedType,
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}

	noteUpdateParam = &schema.Model{
		Name:      names.UpdateParameterName(noteAPIName),
		NakedType: noteNakedType,
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}
)
