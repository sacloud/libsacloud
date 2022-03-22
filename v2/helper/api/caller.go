// Copyright 2016-2022 The Libsacloud Authors
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
	"fmt"
	"net/http"
	"strings"
	"time"

	sacloudhttp "github.com/sacloud/go-http"
	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/helper/defaults"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/sacloud/libsacloud/v2/sacloud/trace"
	"github.com/sacloud/libsacloud/v2/sacloud/trace/otel"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// NewCaller 指定のオプションでsacloud.APICallerを構築して返す
func NewCaller(opts ...*CallerOptions) sacloud.APICaller {
	return newCaller(MergeOptions(opts...))
}

// NewCallerWithDefaults 指定のオプション+環境変数/プロファイルを用いてsacloud.APICallerを構築して返す
//
// DefaultOption()で得られる*CallerOptionsにoptsをマージしてからNewCallerが呼ばれる
func NewCallerWithDefaults(opts *CallerOptions) (sacloud.APICaller, error) {
	defaultOpts, err := DefaultOption()
	if err != nil {
		return nil, err
	}
	return NewCaller(defaultOpts, opts), nil
}

func newCaller(opts *CallerOptions) sacloud.APICaller {
	// build http client
	httpClient := http.DefaultClient
	if opts.HTTPClient != nil {
		httpClient = opts.HTTPClient
	}
	if opts.HTTPRequestTimeout > 0 {
		httpClient.Timeout = time.Duration(opts.HTTPRequestTimeout) * time.Second
	}
	if opts.HTTPRequestTimeout == 0 {
		httpClient.Timeout = 300 * time.Second // デフォルト値
	}
	if opts.HTTPRequestRateLimit > 0 {
		httpClient.Transport = &sacloudhttp.RateLimitRoundTripper{RateLimitPerSec: opts.HTTPRequestRateLimit}
	}
	if opts.HTTPRequestRateLimit == 0 {
		httpClient.Transport = &sacloudhttp.RateLimitRoundTripper{RateLimitPerSec: 10} // デフォルト値
	}

	retryMax := 0
	if opts.RetryMax > 0 {
		retryMax = opts.RetryMax
	}

	retryWaitMax := time.Duration(0)
	if opts.RetryWaitMax > 0 {
		retryWaitMax = time.Duration(opts.RetryWaitMax) * time.Second
	}

	retryWaitMin := time.Duration(0)
	if opts.RetryWaitMin > 0 {
		retryWaitMin = time.Duration(opts.RetryWaitMin) * time.Second
	}

	ua := fmt.Sprintf("libsacloud/%s", libsacloud.Version)
	if opts.UserAgent != "" {
		ua = opts.UserAgent
	}

	caller := &sacloud.Client{
		AccessToken:       opts.AccessToken,
		AccessTokenSecret: opts.AccessTokenSecret,
		UserAgent:         ua,
		AcceptLanguage:    opts.AcceptLanguage,
		RetryMax:          retryMax,
		RetryWaitMax:      retryWaitMax,
		RetryWaitMin:      retryWaitMin,
		HTTPClient:        httpClient,
	}
	sacloud.DefaultStatePollingTimeout = 72 * time.Hour

	if opts.TraceAPI {
		// note: exact once
		trace.AddClientFactoryHooks()
	}
	if opts.TraceHTTP {
		caller.HTTPClient.Transport = &sacloudhttp.TracingRoundTripper{
			Transport: caller.HTTPClient.Transport,
		}
	}
	if opts.OpenTelemetry {
		otel.Initialize(opts.OpenTelemetryOptions...)
		transport := caller.HTTPClient.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		caller.HTTPClient.Transport = otelhttp.NewTransport(transport)
	}

	if opts.FakeMode {
		if opts.FakeStorePath != "" {
			fake.DataStore = fake.NewJSONFileStore(opts.FakeStorePath)
		}
		// note: exact once
		fake.SwitchFactoryFuncToFake()

		SetupFakeDefaults()
	}

	if opts.DefaultZone != "" {
		sacloud.APIDefaultZone = opts.DefaultZone
	}

	if opts.APIRootURL != "" {
		if strings.HasSuffix(opts.APIRootURL, "/") {
			opts.APIRootURL = strings.TrimRight(opts.APIRootURL, "/")
		}
		sacloud.SakuraCloudAPIRoot = opts.APIRootURL
	}

	if len(opts.Zones) > 0 {
		sacloud.SakuraCloudZones = opts.Zones
	}
	return caller
}

func SetupFakeDefaults() {
	defaultInterval := 10 * time.Millisecond

	// update default polling intervals: libsacloud/sacloud
	sacloud.DefaultStatePollingInterval = defaultInterval
	sacloud.DefaultDBStatusPollingInterval = defaultInterval
	// update default polling intervals: libsacloud/helper/setup
	defaults.DefaultDeleteWaitInterval = defaultInterval
	defaults.DefaultProvisioningWaitInterval = defaultInterval
	defaults.DefaultPollingInterval = defaultInterval
	// update default polling intervals: libsacloud/helper/builder
	defaults.DefaultNICUpdateWaitDuration = defaultInterval
	// update default timeouts and span: libsacloud/helper/power
	defaults.DefaultPowerHelperBootRetrySpan = defaultInterval
	defaults.DefaultPowerHelperShutdownRetrySpan = defaultInterval
	defaults.DefaultPowerHelperInitialRequestRetrySpan = defaultInterval
	defaults.DefaultPowerHelperInitialRequestTimeout = defaultInterval * 100

	fake.PowerOnDuration = time.Millisecond
	fake.PowerOffDuration = time.Millisecond
	fake.DiskCopyDuration = time.Millisecond
}
