package define

import (
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
)

var monitorParameter = &dsl.Model{
	Name: "MonitorCondition",
	Fields: []*dsl.FieldDesc{
		{
			Name: "Start",
			Type: meta.TypeTime,
			Tags: &dsl.FieldTags{
				JSON: ",omitempty",
			},
		},
		{
			Name: "End",
			Type: meta.TypeTime,
			Tags: &dsl.FieldTags{
				JSON: ",omitempty",
			},
		},
	},
}

var findParameter = &dsl.Model{
	Name: "FindCondition",
	Fields: []*dsl.FieldDesc{
		conditions.Count(),
		conditions.From(),
		conditions.Sort(),
		conditions.Filter(),
		conditions.Include(),
		conditions.Exclude(),
	},
	//Methods: []*dsl.MethodDesc{
	//	{
	//		Name:             "ClearFilter",
	//		Description:      "フィルタのクリア",
	//		AccessorTypeName: "Filter",
	//	},
	//	{
	//		Name:             "SetANDFilterWithPartialMatch",
	//		Description:      "指定キーの値に中間一致するフィルタを設定 複数指定した場合はAND条件となる",
	//		AccessorTypeName: "Filter",
	//		Arguments: dsl.Arguments{
	//			{
	//				Name: "key",
	//				Type: meta.TypeString,
	//			},
	//			{
	//				Name: "patterns",
	//				Type: meta.TypeStringSlice,
	//			},
	//		},
	//	},
	//	{
	//		Name:             "SetORFilterWithExactMatch",
	//		Description:      "指定キーの値に完全一致するフィルタを設定 複数指定した場合はOR条件となる",
	//		AccessorTypeName: "Filter",
	//		Arguments: dsl.Arguments{
	//			{
	//				Name: "key",
	//				Type: meta.TypeString,
	//			},
	//			{
	//				Name: "patterns",
	//				Type: meta.TypeStringSlice,
	//			},
	//		},
	//	},
	//	{
	//		Name:             "SetNumericFilter",
	//		Description:      "数値型フィルタの設定",
	//		AccessorTypeName: "Filter",
	//		Arguments: dsl.Arguments{
	//			{
	//				Name: "key",
	//				Type: meta.TypeString,
	//			},
	//			{
	//				Name: "op",
	//				Type: meta.Static(accessor.FilterOperator(0)),
	//			},
	//			{
	//				Name: "value",
	//				Type: meta.TypeInt64,
	//			},
	//		},
	//	},
	//	{
	//		Name:             "SetTimeFilter",
	//		Description:      "Time型フィルタの設定",
	//		AccessorTypeName: "Filter",
	//		Arguments: dsl.Arguments{
	//			{
	//				Name: "key",
	//				Type: meta.TypeString,
	//			},
	//			{
	//				Name: "op",
	//				Type: meta.Static(accessor.FilterOperator(0)),
	//			},
	//			{
	//				Name: "value",
	//				Type: meta.TypeTime,
	//			},
	//		},
	//	},
	//},
}

type findCondtionsDef struct{}

var conditions = &findCondtionsDef{}

func (f *findCondtionsDef) From() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "From",
		Type: meta.Static(int(0)),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Count() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Count",
		Type: meta.Static(int(0)),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Sort() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Sort",
		Type: meta.Static([]string{}),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Filter() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Filter",
		Type: meta.Static(search.Filter{}),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Include() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Include",
		Type: meta.Static([]string{}),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Exclude() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Exclude",
		Type: meta.Static([]string{}),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
		},
	}
}
