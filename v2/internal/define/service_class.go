// Copyright 2016-2021 The Libsacloud Authors
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
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	serviceClassAPIName     = "ServiceClass"
	serviceClassAPIPathName = "public/price"
)

var serviceClassAPI = &dsl.Resource{
	Name:       serviceClassAPIName,
	PathName:   serviceClassAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		ops.Find(serviceClassAPIName, serviceClassNakedType, findParameter, serviceClassView),
	},
}

var (
	serviceClassNakedType = meta.Static(naked.ServiceClass{})
	serviceClassView      = &dsl.Model{
		Name:      serviceClassAPIName,
		NakedType: serviceClassNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Def("ServiceClassName", meta.TypeString),
			fields.Def("ServiceClassPath", meta.TypeString),
			fields.Def("DisplayName", meta.TypeString),
			fields.Def("IsPublic", meta.TypeFlag),
			{
				Name: "Price",
				Type: priceModel,
				Tags: &dsl.FieldTags{
					MapConv: ",recursive",
				},
			},
		},
	}

	priceModel = &dsl.Model{
		Name:      "Price",
		NakedType: meta.Static(naked.Price{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Base", meta.TypeInt),
			fields.Def("Daily", meta.TypeInt),
			fields.Def("Hourly", meta.TypeInt),
			fields.Def("Monthly", meta.TypeInt),
			fields.Def("PerUse", meta.TypeInt),
			fields.Def("Basic", meta.TypeInt),
			fields.Def("Traffic", meta.TypeInt),
			fields.Def("DocomoTraffic", meta.TypeInt),
			fields.Def("KddiTraffic", meta.TypeInt),
			fields.Def("SbTraffic", meta.TypeInt),
			fields.Def("SimSheet", meta.TypeInt),
			fields.Def("Zone", meta.TypeString),
		},
	}
)
