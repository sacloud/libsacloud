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

func TestBillOp_ByContract(t *testing.T) {
	t.Parallel()

	client := sacloud.NewBillOp(singletonAPICaller())

	// get account ID
	authStatusOp := sacloud.NewAuthStatusOp(singletonAPICaller())
	authStatus, err := authStatusOp.Read(context.Background())
	if !assert.NoError(t, err) {
		return
	}

	// check current permission
	if !authStatus.ExternalPermission.PermittedBill() {
		t.Skip("current account is not permitted to viewing bills")
	}

	searched, err := client.ByContract(context.Background(), authStatus.AccountID)
	if !assert.NoError(t, err) {
		return
	}

	err = testutil.DoAsserts(
		testutil.AssertTrueFunc(t, len(searched.Bills) > 0, "len(Bills)"),
		testutil.AssertNotEmptyFunc(t, searched.Bills[0].ID, "Bill.ID"),
		testutil.AssertNotEmptyFunc(t, searched.Bills[0].Date, "Bill.Date"),
		testutil.AssertNotEmptyFunc(t, searched.Bills[0].MemberID, "Bill.MemberID"),
	)
	assert.NoError(t, err)

	// Details
	details, err := client.Details(context.Background(), authStatus.MemberCode, searched.Bills[0].ID)
	if !assert.NoError(t, err) {
		return
	}
	err = testutil.DoAsserts(
		testutil.AssertTrueFunc(t, len(details.BillDetails) > 0, "len(Bills)"),
		testutil.AssertNotEmptyFunc(t, details.BillDetails[0].ServiceClassID, "BillDetails.ServiceClassID"),
	)
	assert.NoError(t, err)
}
