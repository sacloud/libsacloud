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
