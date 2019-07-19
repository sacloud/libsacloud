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
			fields.AutoBackupProviderClass(),

			// settings
			fields.AutoBackupBackupSpanType(),
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
		Fields: []*dsl.FieldDesc{
			// creation time only
			fields.AutoBackupProviderClass(),
			fields.AutoBackupDiskID(),

			// backup setting
			fields.AutoBackupBackupSpanType(),
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
		Fields: []*dsl.FieldDesc{
			// backup setting
			fields.AutoBackupBackupSpanType(),
			fields.AutoBackupBackupSpanWeekDays(),
			fields.AutoBackupMaximumNumberOfArchives(),

			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)