package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestBillOp_ByContract(t *testing.T) {
	t.Parallel()

	client := sacloud.NewBillOp(singletonAPICaller())

	// get account ID
	authStatusOp := sacloud.NewAuthStatusOp(singletonAPICaller())
	authStatus, err := authStatusOp.Read(context.Background(), sacloud.APIDefaultZone)
	if !assert.NoError(t, err) {
		return
	}

	// check current permission
	if !authStatus.ExternalPermission.PermittedBill() {
		t.Skip("current account is not permitted to viewing bills")
	}

	searched, err := client.ByContract(context.Background(), sacloud.APIDefaultZone, authStatus.AccountID)
	if !assert.NoError(t, err) {
		return
	}

	err = DoAsserts(
		AssertTrueFunc(t, len(searched.Bills) > 0, "len(Bills)"),
		AssertNotEmptyFunc(t, searched.Bills[0].ID, "Bill.ID"),
		AssertNotEmptyFunc(t, searched.Bills[0].Date, "Bill.Date"),
		AssertNotEmptyFunc(t, searched.Bills[0].MemberID, "Bill.MemberID"),
	)
	assert.NoError(t, err)

	// Details
	details, err := client.Details(context.Background(), sacloud.APIDefaultZone, authStatus.MemberCode, searched.Bills[0].ID)
	if !assert.NoError(t, err) {
		return
	}
	err = DoAsserts(
		AssertTrueFunc(t, len(details.BillDetails) > 0, "len(Bills)"),
		AssertNotEmptyFunc(t, details.BillDetails[0].ServiceClassID, "BillDetails.ServiceClassID"),
	)
	assert.NoError(t, err)
}
