package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLicenseJSON = `
{
	"ID": 123456789012,
	"Name": "test_license (1)",
	"CreatedAt": "2016-05-02T13:58:27+09:00",
	"ModifiedAt": "2016-05-02T13:58:27+09:00",
	"LicenseInfo": {
		"ID": 10011,
		"Name": "Windows RDS SAL + Office SAL",
		"ServiceClass": "cloud\/os\/windows\/rds-sal+office-sal"
	}
}
`

func TestMarshalLicenseJSON(t *testing.T) {
	var license License
	err := json.Unmarshal([]byte(testLicenseJSON), &license)

	assert.NoError(t, err)
	assert.NotEmpty(t, license)

	assert.NotEmpty(t, license.ID)
	assert.NotEmpty(t, license.LicenseInfo)
	assert.NotEmpty(t, license.LicenseInfo.ID)
}
