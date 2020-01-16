// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"os"
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestWebAccelAPI(t *testing.T) {

	api := client.WebAccel
	// find
	res, err := api.Find()
	assert.NoError(t, err)
	assert.NotNil(t, res)

	// read(need SAKURACLOUD_WEBACCEL_SITE_ID)
	siteID := os.Getenv("SAKURACLOUD_WEBACCEL_SITE_ID")
	if siteID == "" {
		t.Skip("SAKURACLOUD_WEBACCEL_SITE_ID is not set. skip")
		return
	}

	site, err := api.Read(siteID)
	assert.NoError(t, err)
	assert.NotNil(t, site)
	assert.Equal(t, siteID, site.ID)

	strCert := os.Getenv("SAKURACLOUD_WEBACCEL_CERT")
	strPKey := os.Getenv("SAKURACLOUD_WEBACCEL_KEY")

	if strCert == "" || strPKey == "" {
		t.Skip("SAKURACLOUD_WEBACCEL_CERT and SAKURACLOUD_WEBACCEL_KEY is not set. skip")
		return
	}

	// check has current cert(cert API support update only)
	if !site.HasCertificate {
		t.Skip("Current certificate is empty(certificate API is supporting only update). skip")
		return
	}

	// certificate
	certRes, err := api.UpdateCertificate(site.ID, &sacloud.WebAccelCertRequest{
		CertificateChain: strCert,
		Key:              strPKey,
	})

	assert.NoError(t, err)
	assert.NotNil(t, certRes.Certificate.Current)

	cert, err := api.ReadCertificate(site.ID)
	assert.NoError(t, err)
	assert.NotNil(t, cert)
	assert.Equal(t, certRes.Certificate.Current.CertificateChain, cert.Current.CertificateChain)

}
