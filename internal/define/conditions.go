package define

import (
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
)

var monitorParameter = &schema.Model{
	Name: "MonitorCondition",
	Fields: []*schema.FieldDesc{
		{
			Name: "Start",
			Type: meta.TypeTime,
			Tags: &schema.FieldTags{
				JSON: ",omitempty",
			},
		},
		{
			Name: "End",
			Type: meta.TypeTime,
			Tags: &schema.FieldTags{
				JSON: ",omitempty",
			},
		},
	},
}

var findParameter = &schema.Model{
	Name: "FindCondition",
	Fields: []*schema.FieldDesc{
		conditions.Count(),
		conditions.From(),
		conditions.Sort(),
		conditions.Filter(),
		conditions.Include(),
		conditions.Exclude(),
	},
}

type findCondtionsDef struct{}

var conditions = &findCondtionsDef{}

func (f *findCondtionsDef) From() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "From",
		Type: meta.Static(int(0)),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Count() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Count",
		Type: meta.Static(int(0)),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Sort() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Sort",
		Type: meta.Static([]string{}),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Filter() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Filter",
		Type: meta.Static(map[string]interface{}{}),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Include() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Include",
		Type: meta.Static([]string{}),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Exclude() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Exclude",
		Type: meta.Static([]string{}),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
		},
	}
}
