// Copyright 2016-2019 The Libsacloud Authors
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
	privateHostPlanAPIName     = "PrivateHostPlan"
	privateHostPlanAPIPathName = "product/privatehost"
)

var privateHostPlanAPI = &dsl.Resource{
	Name:       privateHostPlanAPIName,
	PathName:   privateHostPlanAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		ops.Find(privateHostPlanAPIName, privateHostPlanNakedType, findParameter, privateHostPlanView),
		ops.Read(privateHostPlanAPIName, privateHostPlanNakedType, privateHostPlanView),
	},
}

var (
	privateHostPlanNakedType = meta.Static(naked.PrivateHostPlan{})
	privateHostPlanView      = &dsl.Model{
		Name:      privateHostPlanAPIName,
		NakedType: privateHostPlanNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Class(),
			fields.CPU(),
			fields.MemoryMB(),
			fields.Availability(),
		},
	}
)
