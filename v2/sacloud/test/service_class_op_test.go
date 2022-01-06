// Copyright 2016-2022 The Libsacloud Authors
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

func TestServiceClassOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewServiceClassOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), sacloud.APIDefaultZone, nil)
	assert.NoError(t, err)

	var class *sacloud.ServiceClass
	for _, c := range searched.ServiceClasses {
		if c.IsPublic {
			class = c
			break
		}
	}

	err = testutil.DoAsserts(
		testutil.AssertNotNilFunc(t, class, "ServiceClasses is not nil"),
		testutil.AssertNotEmptyFunc(t, class.ID, "ServiceClasses.ID"),
		testutil.AssertNotEmptyFunc(t, class.ServiceClassName, "ServiceClasses.ServiceClassName"),
		testutil.AssertNotEmptyFunc(t, class.ServiceClassPath, "ServiceClasses.ServiceClassPath"),
		testutil.AssertNotEmptyFunc(t, class.DisplayName, "ServiceClasses.DisplayName"),
		testutil.AssertNotEmptyFunc(t, class.Price, "ServiceClasses.Price"),
		testutil.AssertNotEmptyFunc(t, class.Price.Daily, "ServiceClasses.Price.Daily"),
		testutil.AssertNotEmptyFunc(t, class.Price.Hourly, "ServiceClasses.Price.Hourly"),
		testutil.AssertNotEmptyFunc(t, class.Price.Monthly, "ServiceClasses.Price.Monthly"),
	)
	assert.NoError(t, err)
}
