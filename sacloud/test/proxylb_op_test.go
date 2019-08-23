package test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestProxyLBOp_CRUD(t *testing.T) {
	initProxyLBVariables()

	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		PreCheck: testutil.PreCheckEnvsFunc("SAKURACLOUD_PROXYLB_SERVER0", "SAKURACLOUD_PROXYLB_SERVER1", "SAKURACLOUD_PROXYLB_SERVER2"),

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: testProxyLBCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createProxyLBExpected,
				IgnoreFields: ignoreProxyLBFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testProxyLBRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createProxyLBExpected,
				IgnoreFields: ignoreProxyLBFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testProxyLBUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateProxyLBExpected,
					IgnoreFields: ignoreProxyLBFields,
				}),
			},
			{
				Func: testProxyLBUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateProxyLBToMinExpected,
					IgnoreFields: ignoreProxyLBFields,
				}),
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testProxyLBDelete,
		},
	})
}

var (
	ignoreProxyLBFields        []string
	createProxyLBParam         *sacloud.ProxyLBCreateRequest
	createProxyLBExpected      *sacloud.ProxyLB
	updateProxyLBParam         *sacloud.ProxyLBUpdateRequest
	updateProxyLBExpected      *sacloud.ProxyLB
	updateProxyLBToMinParam    *sacloud.ProxyLBUpdateRequest
	updateProxyLBToMinExpected *sacloud.ProxyLB

	createProxyLBForACMEParam *sacloud.ProxyLBCreateRequest
	updateProxyLBForACMEParam *sacloud.ProxyLBUpdateRequest
)

func initProxyLBVariables() {
	ignoreProxyLBFields = []string{
		"ID",
		"CreatedAt",
		"ModifiedAt",
		"Class",
		"SettingsHash",
		"Region",
		"ProxyNetworks",
		"FQDN",
	}

	createProxyLBParam = &sacloud.ProxyLBCreateRequest{
		Name:        testutil.ResourceName("proxylb"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		Plan:        types.ProxyLBPlans.CPS500,
		HealthCheck: &sacloud.ProxyLBHealthCheck{
			Protocol:  types.ProxyLBProtocols.HTTP,
			Path:      "/",
			DelayLoop: 10,
		},
		SorryServer: &sacloud.ProxyLBSorryServer{
			IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER0"),
			Port:      80,
		},
		BindPorts: []*sacloud.ProxyLBBindPort{
			{
				ProxyMode:       types.ProxyLBProxyModes.HTTP,
				Port:            80,
				RedirectToHTTPS: true,
				AddResponseHeader: []*sacloud.ProxyLBResponseHeader{
					{
						Header: "Cache-Control",
						Value:  "public, max-age=60",
					},
				},
			},
			{
				ProxyMode:    types.ProxyLBProxyModes.HTTPS,
				Port:         443,
				SupportHTTP2: true,
			},
		},
		Servers: []*sacloud.ProxyLBServer{
			{
				IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER1"),
				Port:      80,
				Enabled:   true,
			},
			{
				IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER2"),
				Port:      80,
				Enabled:   true,
			},
		},
		LetsEncrypt: &sacloud.ProxyLBACMESetting{
			Enabled: false,
		},
		StickySession: &sacloud.ProxyLBStickySession{
			Method:  "cookie",
			Enabled: true,
		},
		UseVIPFailover: true,
		Region:         types.ProxyLBRegions.IS1,
	}
	createProxyLBExpected = &sacloud.ProxyLB{
		Name:         createProxyLBParam.Name,
		Description:  createProxyLBParam.Description,
		Tags:         createProxyLBParam.Tags,
		Availability: types.Availabilities.Available,

		Plan:           createProxyLBParam.Plan,
		HealthCheck:    createProxyLBParam.HealthCheck,
		SorryServer:    createProxyLBParam.SorryServer,
		BindPorts:      createProxyLBParam.BindPorts,
		Servers:        createProxyLBParam.Servers,
		LetsEncrypt:    createProxyLBParam.LetsEncrypt,
		StickySession:  createProxyLBParam.StickySession,
		UseVIPFailover: createProxyLBParam.UseVIPFailover,
		Region:         createProxyLBParam.Region,
	}
	updateProxyLBParam = &sacloud.ProxyLBUpdateRequest{
		Name:        testutil.ResourceName("proxylb-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
		HealthCheck: &sacloud.ProxyLBHealthCheck{
			Protocol:  types.ProxyLBProtocols.HTTP,
			Path:      "/index.html",
			DelayLoop: 20,
		},
		SorryServer: &sacloud.ProxyLBSorryServer{
			IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER0"),
			Port:      8080,
		},
		BindPorts: []*sacloud.ProxyLBBindPort{
			{
				ProxyMode:       types.ProxyLBProxyModes.HTTP,
				Port:            8080,
				RedirectToHTTPS: true,
			},
			{
				ProxyMode:    types.ProxyLBProxyModes.HTTPS,
				Port:         8443,
				SupportHTTP2: true,
			},
		},
		Servers: []*sacloud.ProxyLBServer{
			{
				IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER1"),
				Port:      8080,
				Enabled:   true,
			},
			{
				IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER2"),
				Port:      8080,
				Enabled:   true,
			},
		},
		// LetsEncryptのテストはA or CNAMEレコードの登録が必要なため別ケースで行う
		LetsEncrypt: &sacloud.ProxyLBACMESetting{
			Enabled: false,
		},
		StickySession: &sacloud.ProxyLBStickySession{
			Enabled: false,
		},
	}
	updateProxyLBExpected = &sacloud.ProxyLB{
		Name:           updateProxyLBParam.Name,
		Description:    updateProxyLBParam.Description,
		Tags:           updateProxyLBParam.Tags,
		IconID:         testIconID,
		Availability:   types.Availabilities.Available,
		Plan:           createProxyLBParam.Plan,
		HealthCheck:    updateProxyLBParam.HealthCheck,
		SorryServer:    updateProxyLBParam.SorryServer,
		BindPorts:      updateProxyLBParam.BindPorts,
		Servers:        updateProxyLBParam.Servers,
		LetsEncrypt:    updateProxyLBParam.LetsEncrypt,
		StickySession:  updateProxyLBParam.StickySession,
		UseVIPFailover: createProxyLBParam.UseVIPFailover,
		Region:         createProxyLBParam.Region,
	}

	updateProxyLBToMinParam = &sacloud.ProxyLBUpdateRequest{
		Name: testutil.ResourceName("proxylb-to-min"),
		HealthCheck: &sacloud.ProxyLBHealthCheck{
			Protocol:  types.ProxyLBProtocols.TCP,
			DelayLoop: 10,
		},
		LetsEncrypt: &sacloud.ProxyLBACMESetting{
			Enabled: false,
		},
		StickySession: &sacloud.ProxyLBStickySession{
			Enabled: false,
		},
	}
	updateProxyLBToMinExpected = &sacloud.ProxyLB{
		Name:         updateProxyLBToMinParam.Name,
		Availability: types.Availabilities.Available,
		Plan:         createProxyLBParam.Plan,
		HealthCheck:  updateProxyLBToMinParam.HealthCheck,
		SorryServer:  &sacloud.ProxyLBSorryServer{},
		LetsEncrypt: &sacloud.ProxyLBACMESetting{
			Enabled: false,
		},
		StickySession: &sacloud.ProxyLBStickySession{
			Enabled: false,
		},
		UseVIPFailover: createProxyLBParam.UseVIPFailover,
		Region:         createProxyLBParam.Region,
	}

	createProxyLBForACMEParam = &sacloud.ProxyLBCreateRequest{
		Name: testutil.ResourceName("proxylb-acme"),
		Plan: types.ProxyLBPlans.CPS100,
		HealthCheck: &sacloud.ProxyLBHealthCheck{
			Protocol:  types.ProxyLBProtocols.HTTP,
			Path:      "/",
			DelayLoop: 20,
		},
		BindPorts: []*sacloud.ProxyLBBindPort{
			{
				ProxyMode:       types.ProxyLBProxyModes.HTTP,
				Port:            80,
				RedirectToHTTPS: true,
			},
			{
				ProxyMode:    types.ProxyLBProxyModes.HTTPS,
				Port:         443,
				SupportHTTP2: true,
			},
		},
		Servers: []*sacloud.ProxyLBServer{
			{
				IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER0"),
				Port:      80,
				Enabled:   true,
			},
			{
				IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER1"),
				Port:      80,
				Enabled:   true,
			},
		},
		LetsEncrypt: &sacloud.ProxyLBACMESetting{
			Enabled: false,
		},
		UseVIPFailover: true,
	}

	updateProxyLBForACMEParam = &sacloud.ProxyLBUpdateRequest{
		Name: testutil.ResourceName("proxylb-acme"),
		HealthCheck: &sacloud.ProxyLBHealthCheck{
			Protocol:  types.ProxyLBProtocols.HTTP,
			Path:      "/",
			DelayLoop: 20,
		},
		BindPorts: []*sacloud.ProxyLBBindPort{
			{
				ProxyMode:       types.ProxyLBProxyModes.HTTP,
				Port:            80,
				RedirectToHTTPS: true,
			},
			{
				ProxyMode:    types.ProxyLBProxyModes.HTTPS,
				Port:         443,
				SupportHTTP2: true,
			},
		},
		Servers: []*sacloud.ProxyLBServer{
			{
				IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER0"),
				Port:      80,
				Enabled:   true,
			},
			{
				IPAddress: os.Getenv("SAKURACLOUD_PROXYLB_SERVER1"),
				Port:      80,
				Enabled:   true,
			},
		},
		LetsEncrypt: &sacloud.ProxyLBACMESetting{
			CommonName: os.Getenv("SAKURACLOUD_PROXYLB_COMMON_NAME"),
			Enabled:    true,
		},
	}
}

func testProxyLBCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewProxyLBOp(caller)
	return client.Create(ctx, createProxyLBParam)
}

func testProxyLBRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewProxyLBOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testProxyLBUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewProxyLBOp(caller)
	return client.Update(ctx, ctx.ID, updateProxyLBParam)
}

func testProxyLBUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewProxyLBOp(caller)
	return client.Update(ctx, ctx.ID, updateProxyLBToMinParam)
}

func testProxyLBDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewProxyLBOp(caller)
	return client.Delete(ctx, ctx.ID)
}

func TestProxyLBOpLetsEncryptAndHealth(t *testing.T) {
	if !isAccTest() {
		t.Skip("TestProxyLBOpLetsEncrypt only exec at Acceptance Test")
	}

	t.Parallel()
	initProxyLBVariables()
	testutil.PreCheckEnvsFunc(
		"SAKURACLOUD_PROXYLB_SERVER0",
		"SAKURACLOUD_PROXYLB_SERVER1",
		"SAKURACLOUD_PROXYLB_COMMON_NAME",
		"SAKURACLOUD_PROXYLB_ZONE_NAME",
	)(t)

	// prepare variables
	commonName := os.Getenv("SAKURACLOUD_PROXYLB_COMMON_NAME")
	zoneName := os.Getenv("SAKURACLOUD_PROXYLB_ZONE_NAME")
	if !strings.HasSuffix(commonName, zoneName) {
		t.Fatal("$SAKURACLOUD_PROXYLB_COMMON_NAME does not have suffix $SAKURACLOUD_PROXYLB_ZONE_NAME")
	}
	recordName := strings.Replace(commonName, "."+zoneName, "", -1)

	ctx := context.Background()
	proxyLBOp := sacloud.NewProxyLBOp(singletonAPICaller())

	// create proxyLB
	proxyLB, err := proxyLBOp.Create(ctx, createProxyLBForACMEParam)
	if !assert.NoError(t, err) {
		return
	}
	defer func() {
		proxyLBOp.Delete(ctx, proxyLB.ID) // nolint - ignore error
	}()

	// read DNS
	dns, err := lookupDNSByName(singletonAPICaller(), os.Getenv("SAKURACLOUD_PROXYLB_ZONE_NAME"))
	if !assert.NoError(t, err) {
		return
	}

	dns.Records = append(dns.Records, &sacloud.DNSRecord{
		Name:  recordName,
		Type:  types.DNSRecordTypes.CNAME,
		RData: fmt.Sprintf("%s.", proxyLB.FQDN),
		TTL:   10,
	})

	// update DNS record
	dnsOp := sacloud.NewDNSOp(singletonAPICaller())
	dns, err = dnsOp.Update(ctx, dns.ID, &sacloud.DNSUpdateRequest{
		Records: dns.Records,
	})
	if !assert.NoError(t, err) {
		return
	}
	defer func() {
		var records []*sacloud.DNSRecord
		for i, r := range dns.Records {
			if r.Name != recordName {
				records = append(records, dns.Records[i])
			}
		}
		dnsOp.Update(ctx, dns.ID, &sacloud.DNSUpdateRequest{
			Records: records,
		}) // nolint - ignore error
	}()

	time.Sleep(time.Minute)

	// update proxyLB
	retryMax := 10
	done := false

	for retryMax >= 0 {
		proxyLB, err = proxyLBOp.Update(ctx, proxyLB.ID, updateProxyLBForACMEParam)
		if err != nil {
			t.Log("Update Let's encrypt setting is failed. retry after 10 sec.")
			time.Sleep(10 * time.Second)
			retryMax--
			continue
		}
		done = true
		break
	}
	if !done {
		t.Error("Update Let's encrypt settings was failed: given up after 10 retries")
		return
	}

	// renew certs
	err = proxyLBOp.RenewLetsEncryptCert(ctx, proxyLB.ID)
	if !assert.NoError(t, err) {
		return
	}

	time.Sleep(time.Minute)

	// get cert
	certs, err := proxyLBOp.GetCertificates(ctx, proxyLB.ID)

	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, certs.ServerCertificate)
	assert.NotEmpty(t, certs.IntermediateCertificate)
	assert.NotEmpty(t, certs.PrivateKey)
	assert.NotEmpty(t, certs.CertificateCommonName)
	assert.NotEmpty(t, certs.CertificateEndDate)

	// check health status
	status, err := proxyLBOp.HealthStatus(ctx, proxyLB.ID)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, status.CurrentVIP)
}
