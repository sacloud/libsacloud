package define

import (
	"time"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

type fieldsDef struct{}

var fields = &fieldsDef{}

func (f *fieldsDef) ID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ID",
		Type: meta.TypeID,
		ExtendAccessors: []*schema.ExtendAccessor{
			{
				Name: "StringID",
				Type: meta.TypeString,
			},
		},
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

func (f *fieldsDef) SizeMB() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SizeMB",
		Type: meta.TypeInt,
		ExtendAccessors: []*schema.ExtendAccessor{
			{Name: "SizeGB"},
		},
	}
}

func (f *fieldsDef) StorageClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "StorageClass",
		Type: meta.TypeString,
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

func (f *fieldsDef) HostName() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HostName",
		Type: meta.TypeString,
	}
}
func (f *fieldsDef) IPAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddress",
		Type: meta.TypeString,
	}
}
func (f *fieldsDef) User() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "User",
		Type: meta.TypeString,
	}
}
func (f *fieldsDef) Password() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Password",
		Type: meta.TypeString,
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

func (f *fieldsDef) Storage() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Storage",
		Type: meta.Static(naked.Storage{}),
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
