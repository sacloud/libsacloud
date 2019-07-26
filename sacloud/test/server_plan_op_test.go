package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestServerPlanOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewServerPlanOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), sacloud.APIDefaultZone, &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = testutil.DoAsserts(
		testutil.AssertLenFunc(t, searched.ServerPlans, 1, "ServerPlans"),
		testutil.AssertNotEmptyFunc(t, searched.ServerPlans[0].ID, "ServerPlans.ID"),
		testutil.AssertNotEmptyFunc(t, searched.ServerPlans[0].Name, "ServerPlans.Name"),
		testutil.AssertNotEmptyFunc(t, searched.ServerPlans[0].CPU, "ServerPlans.CPU"),
		testutil.AssertNotEmptyFunc(t, searched.ServerPlans[0].Commitment, "ServerPlans.Commitment"),
		testutil.AssertNotEmptyFunc(t, searched.ServerPlans[0].Generation, "ServerPlans.Generation"),
		testutil.AssertNotEmptyFunc(t, searched.ServerPlans[0].Availability, "ServerPlans.Availability"),
	)
	assert.NoError(t, err)
}
