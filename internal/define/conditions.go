// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
			DefaultValue: `time.Now().Truncate(time.Second).Add(-time.Hour)`,
		},
		{
			Name: "End",
			Type: meta.TypeTime,
			Tags: &dsl.FieldTags{
				JSON: ",omitempty",
			},
			DefaultValue: `time.Now().Truncate(time.Second)`,
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
	Methods: []*dsl.MethodDesc{
		{
			Name:        "ClearFilter",
			Description: "フィルタのクリア",
		},
	},
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
		Type: meta.Static(search.SortKeys{}),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Filter() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Filter",
		Type: meta.Static(search.Filter{}),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Include() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Include",
		Type: meta.Static([]string{}),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *findCondtionsDef) Exclude() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Exclude",
		Type: meta.Static([]string{}),
		Tags: &dsl.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}
