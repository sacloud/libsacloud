package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
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

	err = DoAsserts(
		AssertNotNilFunc(t, class, "ServiceClasses is not nil"),
		AssertNotEmptyFunc(t, class.ID, "ServiceClasses.ID"),
		AssertNotEmptyFunc(t, class.ServiceClassName, "ServiceClasses.ServiceClassName"),
		AssertNotEmptyFunc(t, class.ServiceClassPath, "ServiceClasses.ServiceClassPath"),
		AssertNotEmptyFunc(t, class.DisplayName, "ServiceClasses.DisplayName"),
		AssertNotEmptyFunc(t, class.Price, "ServiceClasses.Price"),
		AssertNotEmptyFunc(t, class.Price.Daily, "ServiceClasses.Price.Daily"),
		AssertNotEmptyFunc(t, class.Price.Hourly, "ServiceClasses.Price.Hourly"),
		AssertNotEmptyFunc(t, class.Price.Monthly, "ServiceClasses.Price.Monthly"),
	)
	assert.NoError(t, err)
}
