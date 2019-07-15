package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestInternetPlanOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewInternetPlanOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), sacloud.APIDefaultZone, &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = DoAsserts(
		AssertLenFunc(t, searched.InternetPlans, 1, "InternetPlans"),
		AssertNotEmptyFunc(t, searched.InternetPlans[0].ID, "InternetPlans.ID"),
		AssertNotEmptyFunc(t, searched.InternetPlans[0].Name, "InternetPlans.Name"),
		AssertNotEmptyFunc(t, searched.InternetPlans[0].BandWidthMbps, "InternetPlans.BandWidthMbps"),
		AssertNotEmptyFunc(t, searched.InternetPlans[0].Availability, "InternetPlans.Availability"),
	)
	assert.NoError(t, err)
}
