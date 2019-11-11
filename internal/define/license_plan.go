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
	licenseInfoAPIName     = "LicenseInfo"
	licenseInfoAPIPathName = "product/license"
)

var licenseInfoAPI = &dsl.Resource{
	Name:       licenseInfoAPIName,
	PathName:   licenseInfoAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		ops.Find(licenseInfoAPIName, licenseInfoNakedType, findParameter, licenseInfoView),
		ops.Read(licenseInfoAPIName, licenseInfoNakedType, licenseInfoView),
	},
}

var (
	licenseInfoNakedType = meta.Static(naked.LicenseInfo{})
	licenseInfoView      = &dsl.Model{
		Name:      licenseInfoAPIName,
		NakedType: licenseInfoNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.Def("TermsOfUse", meta.TypeString),
		},
	}
)
