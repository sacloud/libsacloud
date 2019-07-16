package test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestWebAccelOp_Find(t *testing.T) {
	t.Parallel()

	hasPermission, err := hasWebAccelPermission()
	if !assert.NoError(t, err) {
		return
	}
	// check current permission
	if !hasPermission {
		t.Skip("current account is not permitted to using webaccel APIs")
	}

	client := sacloud.NewWebAccelOp(singletonAPICaller())
	searched, err := client.List(context.Background(), sacloud.APIDefaultZone)
	assert.NoError(t, err)

	if searched.Count == 0 {
		t.Skip("webaccel doesn't have any sites")
	}
	site := searched.WebAccels[0]
	err = DoAsserts(
		AssertNotEmptyFunc(t, site.ID, "WebAccel.ID"),
		AssertNotEmptyFunc(t, site.Name, "WebAccel.Name"),
		AssertNotEmptyFunc(t, site.DomainType, "WebAccel.DomainType"),
		AssertNotEmptyFunc(t, site.Domain, "WebAccel.Domain"),
		AssertNotEmptyFunc(t, site.Subdomain, "WebAccel.Subdomain"),
		AssertNotEmptyFunc(t, site.ASCIIDomain, "WebAccel.ASCIIDomain"),
		AssertNotEmptyFunc(t, site.Origin, "WebAccel.Origin"),
		//AssertNotEmptyFunc(t, site.HostHeader, "WebAccel.HostHeader"),
		AssertNotEmptyFunc(t, site.Status, "WebAccel.Status"),
		//AssertNotEmptyFunc(t, site.HasCertificate, "WebAccel.HasCertificate"),
		//AssertNotEmptyFunc(t, site.HasOldCertificate, "WebAccel.HasOldCertificate"),
		//AssertNotEmptyFunc(t, site.GibSentInLastWeek, "WebAccel.GibSentInLastWeek"),
		//AssertNotEmptyFunc(t, site.CertValidNotBefore, "WebAccel.CertValidNotBefore"),
		//AssertNotEmptyFunc(t, site.CertValidNotAfter, "WebAccel.CertValidNotAfter"),
		AssertNotEmptyFunc(t, site.CreatedAt, "WebAccel.CreatedAt"),
	)
	assert.NoError(t, err)

	// read
	read, err := client.Read(context.Background(), sacloud.APIDefaultZone, site.ID)
	assert.NoError(t, err)
	assert.Equal(t, site, read)
}

func TestWebAccelOp_Cert(t *testing.T) {
	t.Parallel()

	PreCheckEnvsFunc("SAKURACLOUD_WEBACCEL_SITE_ID", "SAKURACLOUD_WEBACCEL_CERT", "SAKURACLOUD_WEBACCEL_KEY")(t)

	hasPermission, err := hasWebAccelPermission()
	if !assert.NoError(t, err) {
		return
	}
	// check current permission
	if !hasPermission {
		t.Skip("current account is not permitted to using webaccel APIs")
	}

	client := sacloud.NewWebAccelOp(singletonAPICaller())
	ctx := context.Background()
	id := types.StringID(os.Getenv("SAKURACLOUD_WEBACCEL_SITE_ID"))
	crt := os.Getenv("SAKURACLOUD_WEBACCEL_CERT")
	key := os.Getenv("SAKURACLOUD_WEBACCEL_KEY")

	// update certs
	certs, err := client.UpdateCertificate(ctx, sacloud.APIDefaultZone, id, &sacloud.WebAccelCertUpdateRequest{
		CertificateChain: crt,
		Key:              key,
	})
	if !assert.NoError(t, err) {
		return
	}

	// read cert
	read, err := client.ReadCertificate(ctx, sacloud.APIDefaultZone, id)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, certs, read)
}

func TestWebAccelOp_DeleteAllCache(t *testing.T) {
	t.Parallel()

	PreCheckEnvsFunc("SAKURACLOUD_WEBACCEL_DOMAIN")(t)

	hasPermission, err := hasWebAccelPermission()
	if !assert.NoError(t, err) {
		return
	}
	// check current permission
	if !hasPermission {
		t.Skip("current account is not permitted to using webaccel APIs")
	}

	client := sacloud.NewWebAccelOp(singletonAPICaller())
	ctx := context.Background()
	domain := os.Getenv("SAKURACLOUD_WEBACCEL_DOMAIN")

	// delete cache
	err = client.DeleteAllCache(ctx, sacloud.APIDefaultZone, &sacloud.WebAccelDeleteAllCacheRequest{
		Domain: domain,
	})
	if !assert.NoError(t, err) {
		return
	}
}

func TestWebAccelOp_DeleteCache(t *testing.T) {
	t.Parallel()

	PreCheckEnvsFunc("SAKURACLOUD_WEBACCEL_URLS")(t)

	hasPermission, err := hasWebAccelPermission()
	if !assert.NoError(t, err) {
		return
	}
	// check current permission
	if !hasPermission {
		t.Skip("current account is not permitted to using webaccel APIs")
	}

	client := sacloud.NewWebAccelOp(singletonAPICaller())
	ctx := context.Background()
	strURLs := os.Getenv("SAKURACLOUD_WEBACCEL_URLS")

	urls := strings.Split(strURLs, ",")

	// delete cache
	result, err := client.DeleteCache(ctx, sacloud.APIDefaultZone, &sacloud.WebAccelDeleteCacheRequest{
		URL: urls,
	})
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, result)
}

func hasWebAccelPermission() (bool, error) {
	authStatusOp := sacloud.NewAuthStatusOp(singletonAPICaller())
	authStatus, err := authStatusOp.Read(context.Background(), sacloud.APIDefaultZone)
	if err != nil {
		return false, err
	}

	// check current permission
	return authStatus.ExternalPermission.PermittedWebAccel(), nil
}
