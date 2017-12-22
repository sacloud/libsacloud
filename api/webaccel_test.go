package api

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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
