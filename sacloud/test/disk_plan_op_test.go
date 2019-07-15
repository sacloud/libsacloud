package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestDiskPlanOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewDiskPlanOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), sacloud.APIDefaultZone, &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = DoAsserts(
		AssertLenFunc(t, searched.DiskPlans, 1, "DiskPlans"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].ID, "DiskPlans.ID"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].Name, "DiskPlans.Name"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].Availability, "DiskPlans.Availability"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].StorageClass, "DiskPlans.StorageClass"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].Size, "DiskPlans.Size"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].Size[0].Availability, "DiskPlans.Size.Availability"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].Size[0].DisplaySize, "DiskPlans.Size.DisplaySize"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].Size[0].DisplaySuffix, "DiskPlans.Size.DisplaySuffix"),
		AssertNotEmptyFunc(t, searched.DiskPlans[0].Size[0].SizeMB, "DiskPlans.Size.SizeMB"),
	)
	assert.NoError(t, err)
}
