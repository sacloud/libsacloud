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

		// updateSettings
		ops.UpdateCommonServiceItemSettings(dnsAPIName, dnsUpdateSettingsNakedType, dnsUpdateSettingsParam, dnsView),

		// delete
		ops.Delete(dnsAPIName),
	},
}

var (
	dnsNakedType               = meta.Static(naked.DNS{})
	dnsUpdateSettingsNakedType = meta.Static(naked.DNSSettingsUpdate{})

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
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"dns"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// creation time only
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
			// common fields
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// setting
			fields.DNSRecords(),
			// settings hash
			fields.SettingsHash(),
		},
	}

	dnsUpdateSettingsParam = &dsl.Model{
		Name:      names.UpdateSettingsParameterName(dnsAPIName),
		NakedType: dnsNakedType,
		Fields: []*dsl.FieldDesc{
			// setting
			fields.DNSRecords(),
			// settings hash
			fields.SettingsHash(),
		},
	}
)
