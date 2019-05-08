package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	findParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			conditions.Count(),
			conditions.From(),
			conditions.Sort(),
			conditions.Filter(),
			conditions.Include(),
			conditions.Exclude(),
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

	result := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Scope(),
			fields.NoteClass(),
			fields.NoteContent(),
			fields.Icon(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	Resources.Define("Note").
		OperationCRUD(meta.Static(naked.Note{}), findParam, createParam, updateParam, result)
}
