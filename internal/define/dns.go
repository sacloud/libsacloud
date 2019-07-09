package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	dnsAPIName     = "DNS"
	dnsAPIPathName = "commonserviceitem"
)

var dnsAPI = &dsl.Resource{
	Name:       dnsAPIName,
	PathName:   dnsAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(dnsAPIName, dnsNakedType, findParameter, dnsView),

		// create
		ops.CreateCommonServiceItem(dnsAPIName, dnsNakedType, dnsCreateParam, dnsView),

		// read
		ops.ReadCommonServiceItem(dnsAPIName, dnsNakedType, dnsView),

		// update
		ops.UpdateCommonServiceItem(dnsAPIName, dnsNakedType, dnsUpdateParam, dnsView),

		// delete
		ops.Delete(dnsAPIName),
	},
}

var (
	dnsNakedType = meta.Static(naked.DNS{})

	dnsView = &dsl.Model{
		Name:      dnsAPIName,
		NakedType: dnsNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),

			fields.DNSProviderClass(),

			// settings
			fields.DNSRecords(),
			fields.SettingsHash(),

			// status
			fields.DNSZone(),
			fields.DNSNameServers(),
		},
	}

	dnsCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(dnsAPIName),
		NakedType: dnsNakedType,
		Fields: []*dsl.FieldDesc{
			// creation time only
			fields.DNSProviderClass(),
			{
				Name: "Name",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					Validate: "required",
					MapConv:  "Name/Status.Zone", // NameとStatus.Zone2箇所に同じ値を設定
				},
			},

			// setting
			fields.DNSRecords(),

			// common fields
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	dnsUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(dnsAPIName),
		NakedType: dnsNakedType,
		Fields: []*dsl.FieldDesc{
			// setting
			fields.DNSRecords(),

			// common fields
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
