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

package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestPrivateHostPlanOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewPrivateHostPlanOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), "tk1a", &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = testutil.DoAsserts(
		testutil.AssertLenFunc(t, searched.PrivateHostPlans, 1, "PrivateHostPlans"),
		testutil.AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].ID, "PrivateHostPlans.ID"),
		testutil.AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].Name, "PrivateHostPlans.Name"),
		testutil.AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].Class, "PrivateHostPlans.Class"),
		testutil.AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].CPU, "PrivateHostPlans.CPU"),
		testutil.AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].MemoryMB, "PrivateHostPlans.MemoryMB"),
	)
	assert.NoError(t, err)
}
