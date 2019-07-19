package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestLicenseInfoOp_Find(t *testing.T) {
	t.Parallel()

	client := sacloud.NewLicenseInfoOp(singletonAPICaller())

	searched, err := client.Find(context.Background(), &sacloud.FindCondition{Count: 1})
	assert.NoError(t, err)

	err = DoAsserts(
		AssertLenFunc(t, searched.LicenseInfo, 1, "LicenseInfos"),
		AssertNotEmptyFunc(t, searched.LicenseInfo[0].ID, "LicenseInfos.ID"),
		AssertNotEmptyFunc(t, searched.LicenseInfo[0].Name, "LicenseInfos.Name"),
		AssertNotEmptyFunc(t, searched.LicenseInfo[0].TermsOfUse, "LicenseInfos.TermsOfUse"),
	)
	assert.NoError(t, err)
}
