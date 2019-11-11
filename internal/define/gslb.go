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
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	gslbAPIName     = "GSLB"
	gslbAPIPathName = "commonserviceitem"
)

var gslbAPI = &dsl.Resource{
	Name:       gslbAPIName,
	PathName:   gslbAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(gslbAPIName, gslbNakedType, findParameter, gslbView),

		// create
		ops.CreateCommonServiceItem(gslbAPIName, gslbNakedType, gslbCreateParam, gslbView),

		// read
		ops.ReadCommonServiceItem(gslbAPIName, gslbNakedType, gslbView),

		// update
		ops.UpdateCommonServiceItem(gslbAPIName, gslbNakedType, gslbUpdateParam, gslbView),

		// patch
		ops.PatchCommonServiceItem(gslbAPIName, gslbNakedType, patchModel(gslbUpdateParam), gslbView),

		// delete
		ops.Delete(gslbAPIName),
	},
}

var (
	gslbNakedType = meta.Static(naked.GSLB{})

	gslbView = &dsl.Model{
		Name:      gslbAPIName,
		NakedType: gslbNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.SettingsHash(),
			fields.GSLBFQDN(),
			// settings
			fields.GSLBDelayLoop(),
			fields.GSLBWeighted(),
			fields.GSLBHealthCheck(),
			fields.GSLBSorryServer(),
			fields.GSLBDestinationServers(),
		},
	}

	gslbCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(gslbAPIName),
		NakedType: gslbNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"gslb"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			fields.GSLBHealthCheck(),
			fields.GSLBDelayLoop(),
			fields.GSLBWeighted(),
			fields.GSLBSorryServer(),
			fields.GSLBDestinationServers(),

			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	gslbUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(gslbAPIName),
		NakedType: gslbNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// settings
			fields.GSLBHealthCheck(),
			fields.GSLBDelayLoop(),
			fields.GSLBWeighted(),
			fields.GSLBSorryServer(),
			fields.GSLBDestinationServers(),
			// settings hash
			fields.SettingsHash(),
		},
	}
)
