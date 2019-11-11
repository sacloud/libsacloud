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
	autoBackupAPIName     = "AutoBackup"
	autoBackupAPIPathName = "commonserviceitem"
)

var autoBackupAPI = &dsl.Resource{
	Name:       autoBackupAPIName,
	PathName:   autoBackupAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(autoBackupAPIName, autoBackupNakedType, findParameter, autoBackupView),

		// create
		ops.CreateCommonServiceItem(autoBackupAPIName, autoBackupNakedType, autoBackupCreateParam, autoBackupView),

		// read
		ops.ReadCommonServiceItem(autoBackupAPIName, autoBackupNakedType, autoBackupView),

		// update
		ops.UpdateCommonServiceItem(autoBackupAPIName, autoBackupNakedType, autoBackupUpdateParam, autoBackupView),

		// patch
		ops.PatchCommonServiceItem(autoBackupAPIName, autoBackupNakedType, patchModel(autoBackupUpdateParam), autoBackupView),

		// delete
		ops.Delete(autoBackupAPIName),
	},
}

var (
	autoBackupNakedType = meta.Static(naked.AutoBackup{})

	autoBackupView = &dsl.Model{
		Name:      autoBackupAPIName,
		NakedType: autoBackupNakedType,
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
			fields.AutoBackupBackupSpanWeekDays(),
			fields.AutoBackupMaximumNumberOfArchives(),
			fields.SettingsHash(),

			// status
			fields.AutoBackupDiskID(),
			fields.AutoBackupAccountID(),
			fields.AutoBackupZoneID(),
			fields.AutoBackupZoneName(),
		},
	}

	autoBackupCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(autoBackupAPIName),
		NakedType: autoBackupNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"autobackup"`,
			},
			{
				Name: "BackupSpanType",
				Type: meta.TypeBackupSpanType,
				Tags: &dsl.FieldTags{
					MapConv: "Settings.Autobackup.BackupSpanType",
				},
				Value: `types.BackupSpanTypes.Weekdays`,
			},
		},

		Fields: []*dsl.FieldDesc{
			// creation time only
			fields.AutoBackupDiskID(),

			// backup setting
			fields.AutoBackupBackupSpanWeekDays(),
			fields.AutoBackupMaximumNumberOfArchives(),

			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	autoBackupUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(autoBackupAPIName),
		NakedType: autoBackupNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "BackupSpanType",
				Type: meta.TypeBackupSpanType,
				Tags: &dsl.FieldTags{
					MapConv: "Settings.Autobackup.BackupSpanType",
				},
				Value: `types.BackupSpanTypes.Weekdays`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// backup setting
			fields.AutoBackupBackupSpanWeekDays(),
			fields.AutoBackupMaximumNumberOfArchives(),
			// settings hash
			fields.SettingsHash(),
		},
	}
)
