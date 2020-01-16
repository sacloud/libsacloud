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

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZone_Find(t *testing.T) {
	api := client.Facility.Zone
	res, err := api.Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Zones)
	assert.NotEmpty(t, res.Zones[0].ID)

	id := res.Zones[0].ID

	zone, err := api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, zone)
	assert.NotEmpty(t, zone.ID)
}
