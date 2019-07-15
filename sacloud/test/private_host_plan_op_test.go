package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestPrivateHostPlanOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewPrivateHostPlanOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), "tk1a", &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = DoAsserts(
		AssertLenFunc(t, searched.PrivateHostPlans, 1, "PrivateHostPlans"),
		AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].ID, "PrivateHostPlans.ID"),
		AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].Name, "PrivateHostPlans.Name"),
		AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].Class, "PrivateHostPlans.Class"),
		AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].CPU, "PrivateHostPlans.CPU"),
		AssertNotEmptyFunc(t, searched.PrivateHostPlans[0].MemoryMB, "PrivateHostPlans.MemoryMB"),
	)
	assert.NoError(t, err)
}
