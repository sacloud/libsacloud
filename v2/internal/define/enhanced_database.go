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
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	enhancedDatabaseAPIName     = "EnhancedDB"
	enhancedDatabaseAPIPathName = "commonserviceitem"
)

var enhancedDatabaseAPI = &dsl.Resource{
	Name:       enhancedDatabaseAPIName,
	PathName:   enhancedDatabaseAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(enhancedDatabaseAPIName, enhancedDatabaseNakedType, findParameter, enhancedDatabaseView),

		// create
		ops.CreateCommonServiceItem(enhancedDatabaseAPIName, enhancedDatabaseNakedType, enhancedDatabaseCreateParam, enhancedDatabaseView),

		// read
		ops.ReadCommonServiceItem(enhancedDatabaseAPIName, enhancedDatabaseNakedType, enhancedDatabaseView),

		// update
		ops.UpdateCommonServiceItem(enhancedDatabaseAPIName, enhancedDatabaseNakedType, enhancedDatabaseUpdateParam, enhancedDatabaseView),

		// delete
		ops.Delete(enhancedDatabaseAPIName),

		// Set Password
		{
			ResourceName: enhancedDatabaseAPIName,
			Name:         "SetPassword",
			PathFormat:   dsl.IDAndSuffixPathFormat("enhanceddb/set-password"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: enhancedDatabaseUserNakedType,
				Name: "CommonServiceItem",
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", enhancedDatabaseSetPasswordParam, "CommonServiceItem"),
			},
		},
	},
}

var (
	enhancedDatabaseNakedType     = meta.Static(naked.EnhancedDB{})
	enhancedDatabaseUserNakedType = meta.Static(naked.EnhancedDBPasswordSettings{})

	enhancedDatabaseView = &dsl.Model{
		Name:      enhancedDatabaseAPIName,
		NakedType: enhancedDatabaseNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),

			// settings
			fields.SettingsHash(),
			fields.EnhancedDBMaxConnections(),

			// status
			fields.EnhancedDBDatabaseName(),
			fields.EnhancedDBDatabaseType(),
			fields.EnhancedDBDatabaseRegion(),
			fields.EnhancedDBDatabaseHostName(),
			fields.EnhancedDBDatabasePort(),
		},
	}

	enhancedDatabaseCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(enhancedDatabaseAPIName),
		NakedType: enhancedDatabaseNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"enhanceddb"`,
			},
			{
				Name: "Region",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Status.Region",
				},
				Value: `"is1"`,
			},
			{
				Name: "DatabaseType",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Status.DatabaseType",
				},
				Value: `"tidb"`,
			},
			{
				Name: "MaxConnections",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: "Settings.EnhancedDB.MaxConnections",
				},
				Value: `50`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// settings
			fields.EnhancedDBDatabaseName(),
		},
	}

	enhancedDatabaseUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(enhancedDatabaseAPIName),
		NakedType: enhancedDatabaseNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "MaxConnections",
				Type: meta.TypeInt,
				Tags: &dsl.FieldTags{
					MapConv: "Settings.EnhancedDB.MaxConnections",
				},
				Value: `50`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// settings hash
			fields.SettingsHash(),
		},
	}

	enhancedDatabaseSetPasswordParam = &dsl.Model{
		Name:      enhancedDatabaseAPIName + "SetPasswordRequest",
		NakedType: meta.Static(naked.EnhancedDBPasswordSetting{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "Password",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "EnhancedDB.Password",
				},
			},
		},
	}
)
