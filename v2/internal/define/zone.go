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
	zoneAPIName     = "Zone"
	zoneAPIPathName = "zone"
)

var zoneAPI = &dsl.Resource{
	Name:       zoneAPIName,
	PathName:   zoneAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		ops.Find(zoneAPIName, zoneNakedType, findParameter, zoneView),
		ops.Read(zoneAPIName, zoneNakedType, zoneView),
	},
}

var (
	zoneNakedType = meta.Static(naked.Zone{})
	zoneView      = &dsl.Model{
		Name:      zoneAPIName,
		NakedType: zoneNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.DisplayOrder(),
			fields.IsDummy(),
			fields.VNCProxy(),
			fields.FTPServer(),
			fields.Region(),
		},
	}
)
