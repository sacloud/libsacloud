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

package query

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// ZoneIDFromArchiveShareKey アーカイブ共有キーから対応するゾーンIDを取得
func ZoneIDFromArchiveShareKey(ctx context.Context, zoneAPI sacloud.ZoneAPI, key types.ArchiveShareKey) (types.ID, error) {
	if !key.ValidFormat() {
		return types.ID(0), fmt.Errorf("archive share key is invalid format: key:%q", key.String())
	}

	searched, err := zoneAPI.Find(ctx, &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key("Name"): search.ExactMatch(key.Zone()),
		},
		Include: []string{"ID"},
	})
	if err != nil {
		return types.ID(0), err
	}
	if searched.Count == 0 {
		return types.ID(0), fmt.Errorf("zone %q is not found", key.Zone())
	}
	return searched.Zones[0].ID, nil
}
