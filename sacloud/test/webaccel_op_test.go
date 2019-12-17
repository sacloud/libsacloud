// Copyright 2016-2019 The Libsacloud Authors
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

package test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
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
	searched, err := client.List(context.Background())
	assert.NoError(t, err)

	if searched.Count == 0 {
		t.Skip("webaccel doesn't have any sites")
	}
	site := searched.WebAccels[0]
	err = testutil.DoAsserts(
		testutil.AssertNotEmptyFunc(t, site.ID, "WebAccel.ID"),
		testutil.AssertNotEmptyFunc(t, site.Name, "WebAccel.Name"),
		testutil.AssertNotEmptyFunc(t, site.DomainType, "WebAccel.DomainType"),
		testutil.AssertNotEmptyFunc(t, site.Domain, "WebAccel.Domain"),
		testutil.AssertNotEmptyFunc(t, site.Subdomain, "WebAccel.Subdomain"),
		testutil.AssertNotEmptyFunc(t, site.ASCIIDomain, "WebAccel.ASCIIDomain"),
		testutil.AssertNotEmptyFunc(t, site.Origin, "WebAccel.Origin"),
		//testutil.AssertNotEmptyFunc(t, site.HostHeader, "WebAccel.HostHeader"),
		testutil.AssertNotEmptyFunc(t, site.Status, "WebAccel.Status"),
		//testutil.AssertNotEmptyFunc(t, site.HasCertificate, "WebAccel.HasCertificate"),
		//testutil.AssertNotEmptyFunc(t, site.HasOldCertificate, "WebAccel.HasOldCertificate"),
		//testutil.AssertNotEmptyFunc(t, site.GibSentInLastWeek, "WebAccel.GibSentInLastWeek"),
		//testutil.AssertNotEmptyFunc(t, site.CertValidNotBefore, "WebAccel.CertValidNotBefore"),
		//testutil.AssertNotEmptyFunc(t, site.CertValidNotAfter, "WebAccel.CertValidNotAfter"),
		testutil.AssertNotEmptyFunc(t, site.CreatedAt, "WebAccel.CreatedAt"),
	)
	assert.NoError(t, err)

	// read
	read, err := client.Read(context.Background(), site.ID)
	assert.NoError(t, err)
	assert.Equal(t, site, read)
}

func TestWebAccelOp_Cert(t *testing.T) {

	if !isAccTest() {
		t.Skip("TestWebAccelOp_Cert only exec at Acceptance Test")
	}

	t.Parallel()

	envKeys := []string{
		"SAKURACLOUD_WEBACCEL_SITE_ID",
		"SAKURACLOUD_WEBACCEL_CERT",
		"SAKURACLOUD_WEBACCEL_KEY",
		"SAKURACLOUD_WEBACCEL_CERT_UPD",
		"SAKURACLOUD_WEBACCEL_KEY_UPD",
	}
	testutil.PreCheckEnvsFunc(envKeys...)(t)

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
	crtUpd := os.Getenv("SAKURACLOUD_WEBACCEL_CERT_UPD")
	keyUpd := os.Getenv("SAKURACLOUD_WEBACCEL_KEY_UPD")

	// create certs
	_, err = client.CreateCertificate(ctx, id, &sacloud.WebAccelCertRequest{
		CertificateChain: crt,
		Key:              key,
	})
	if !assert.NoError(t, err) {
		return
	}

	// update certs
	certs, err := client.UpdateCertificate(ctx, id, &sacloud.WebAccelCertRequest{
		CertificateChain: crtUpd,
		Key:              keyUpd,
	})
	if !assert.NoError(t, err) {
		return
	}

	// read cert
	read, err := client.ReadCertificate(ctx, id)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, certs, read)
}

func TestWebAccelOp_DeleteAllCache(t *testing.T) {
	t.Parallel()

	testutil.PreCheckEnvsFunc("SAKURACLOUD_WEBACCEL_DOMAIN")(t)

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
	err = client.DeleteAllCache(ctx, &sacloud.WebAccelDeleteAllCacheRequest{
		Domain: domain,
	})
	if !assert.NoError(t, err) {
		return
	}
}

func TestWebAccelOp_DeleteCache(t *testing.T) {
	t.Parallel()

	testutil.PreCheckEnvsFunc("SAKURACLOUD_WEBACCEL_URLS")(t)

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
	result, err := client.DeleteCache(ctx, &sacloud.WebAccelDeleteCacheRequest{
		URL: urls,
	})
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, result)
}

func hasWebAccelPermission() (bool, error) {
	authStatusOp := sacloud.NewAuthStatusOp(singletonAPICaller())
	authStatus, err := authStatusOp.Read(context.Background())
	if err != nil {
		return false, err
	}

	// check current permission
	return authStatus.ExternalPermission.PermittedWebAccel(), nil
}
