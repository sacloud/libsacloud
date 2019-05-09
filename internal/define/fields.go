package define

import (
	"time"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

type fieldsDef struct{}
type findCondtionsDef struct{}

var fields = &fieldsDef{}
var conditions = &findCondtionsDef{}

func (f *fieldsDef) ID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ID",
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) Name() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Name",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "required",
		},
	}
}

func (f *fieldsDef) IconID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IconID",
		Tags: &schema.FieldTags{
			MapConv: "Icon.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) Tags() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Tags",
		Type: meta.TypeStringSlice,
	}
}

func (f *fieldsDef) NoteClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) NoteContent() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Content",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) Description() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Description",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) Availability() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Availability",
		Type: meta.TypeAvailability,
	}
}

func (f *fieldsDef) Scope() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Scope",
		Type: meta.TypeScope,
	}
}

func (f *fieldsDef) DisplayOrder() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DisplayOrder",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) IsDummy() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IsDummy",
		Type: meta.TypeFlag,
	}
}

func (f *fieldsDef) VNCProxy() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "VNCProxy",
		Type: meta.Static(naked.VNCProxy{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) FTPServer() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "FTPServer",
		Type: meta.Static(naked.FTPServer{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}
func (f *fieldsDef) Region() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Region",
		Type: meta.Static(naked.Region{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Icon() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Icon",
		Type: meta.Static(naked.Icon{}),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) CreatedAt() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CreatedAt",
		Type: meta.Static(time.Time{}),
	}
}

func (f *fieldsDef) ModifiedAt() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ModifiedAt",
		Type: meta.Static(time.Time{}),
	}
}

/************************************************
 Find Conditions
************************************************/

func (f *findCondtionsDef) From() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "From",
		Type: meta.Static(int(0)),
	}
}

func (f *findCondtionsDef) Count() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Count",
		Type: meta.Static(int(0)),
	}
}

func (f *findCondtionsDef) Sort() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Sort",
		Type: meta.Static([]string{}),
	}
}

func (f *findCondtionsDef) Filter() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Filter",
		Type: meta.Static(map[string]interface{}{}),
	}
}

func (f *findCondtionsDef) Include() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Include",
		Type: meta.Static([]string{}),
	}
}

func (f *findCondtionsDef) Exclude() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Exclude",
		Type: meta.Static([]string{}),
	}
}
