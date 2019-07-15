package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestServerPlanOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewServerPlanOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), sacloud.APIDefaultZone, &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = DoAsserts(
		AssertLenFunc(t, searched.ServerPlans, 1, "ServerPlans"),
		AssertNotEmptyFunc(t, searched.ServerPlans[0].ID, "ServerPlans.ID"),
		AssertNotEmptyFunc(t, searched.ServerPlans[0].Name, "ServerPlans.Name"),
		AssertNotEmptyFunc(t, searched.ServerPlans[0].CPU, "ServerPlans.CPU"),
		AssertNotEmptyFunc(t, searched.ServerPlans[0].Commitment, "ServerPlans.Commitment"),
		AssertNotEmptyFunc(t, searched.ServerPlans[0].Generation, "ServerPlans.Generation"),
		AssertNotEmptyFunc(t, searched.ServerPlans[0].Availability, "ServerPlans.Availability"),
	)
	assert.NoError(t, err)
}
