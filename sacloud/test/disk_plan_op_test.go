package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDiskPlanOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewDiskPlanOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), sacloud.APIDefaultZone, &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = testutil.DoAsserts(
		testutil.AssertLenFunc(t, searched.DiskPlans, 1, "DiskPlans"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].ID, "DiskPlans.ID"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].Name, "DiskPlans.Name"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].Availability, "DiskPlans.Availability"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].StorageClass, "DiskPlans.StorageClass"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].Size, "DiskPlans.Size"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].Size[0].Availability, "DiskPlans.Size.Availability"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].Size[0].DisplaySize, "DiskPlans.Size.DisplaySize"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].Size[0].DisplaySuffix, "DiskPlans.Size.DisplaySuffix"),
		testutil.AssertNotEmptyFunc(t, searched.DiskPlans[0].Size[0].SizeMB, "DiskPlans.Size.SizeMB"),
	)
	assert.NoError(t, err)
}
