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

package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestInternetPlanOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewInternetPlanOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), sacloud.APIDefaultZone, &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = testutil.DoAsserts(
		testutil.AssertLenFunc(t, searched.InternetPlans, 1, "InternetPlans"),
		testutil.AssertNotEmptyFunc(t, searched.InternetPlans[0].ID, "InternetPlans.ID"),
		testutil.AssertNotEmptyFunc(t, searched.InternetPlans[0].Name, "InternetPlans.Name"),
		testutil.AssertNotEmptyFunc(t, searched.InternetPlans[0].BandWidthMbps, "InternetPlans.BandWidthMbps"),
		testutil.AssertNotEmptyFunc(t, searched.InternetPlans[0].Availability, "InternetPlans.Availability"),
	)
	assert.NoError(t, err)
}
