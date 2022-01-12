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
	"net/http"
	"os"
	"strconv"
	"strings"

	sacloudhttp "github.com/sacloud/go-http"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/sacloud/libsacloud/v2/sacloud/trace/otel"
)

// CallerOptions sacloud.APICallerを作成する際のオプション
type CallerOptions struct {
	AccessToken       string
	AccessTokenSecret string

	APIRootURL     string
	DefaultZone    string
	Zones          []string
	AcceptLanguage string

	HTTPClient *http.Client

	HTTPRequestTimeout   int
	HTTPRequestRateLimit int

	RetryMax     int
	RetryWaitMax int
	RetryWaitMin int

	UserAgent string

	TraceAPI             bool
	TraceHTTP            bool
	OpenTelemetry        bool
	OpenTelemetryOptions []otel.Option

	FakeMode      bool
	FakeStorePath string
}

// DefaultOption 環境変数、プロファイルからCallerOptionsを組み立てて返す
//
// プロファイルは環境変数`SAKURACLOUD_PROFILE`または`USACLOUD_PROFILE`でプロファイル名が指定されていればそちらを優先し、
// 未指定の場合は通常のプロファイル処理(~/.usacloud/currentファイルから読み込み)される。
// 同じ項目を複数箇所で指定していた場合、環境変数->プロファイルの順で上書きされたものが返される
func DefaultOption() (*CallerOptions, error) {
	return DefaultOptionWithProfile("")
}

// DefaultOptionWithProfile 環境変数、プロファイルからCallerOptionsを組み立てて返す
//
// プロファイルは引数を優先し、空の場合は環境変数`SAKURACLOUD_PROFILE`または`USACLOUD_PROFILE`が利用され、
// それも空の場合は通常のプロファイル処理(~/.usacloud/currentファイルから読み込み)される。
// 同じ項目を複数箇所で指定していた場合、環境変数->プロファイルの順で上書きされたものが返される
func DefaultOptionWithProfile(profileName string) (*CallerOptions, error) {
	if profileName == "" {
		profileName = stringFromEnvMulti([]string{"SAKURACLOUD_PROFILE", "USACLOUD_PROFILE"}, "")
	}
	fromProfile, err := OptionsFromProfile(profileName)
	if err != nil {
		return nil, err
	}
	return MergeOptions(OptionsFromEnv(), fromProfile, defaultOption), nil
}

var defaultOption = &CallerOptions{
	APIRootURL:           sacloud.SakuraCloudAPIRoot,
	DefaultZone:          sacloud.APIDefaultZone,
	HTTPRequestTimeout:   300,
	HTTPRequestRateLimit: 5,
	RetryMax:             sacloudhttp.DefaultRetryMax,
	RetryWaitMax:         int(sacloudhttp.DefaultRetryWaitMax.Seconds()),
	RetryWaitMin:         int(sacloudhttp.DefaultRetryWaitMin.Seconds()),
}

// MergeOptions 指定のCallerOptionsの非ゼロ値フィールドをoのコピーにマージして返す
func MergeOptions(opts ...*CallerOptions) *CallerOptions {
	merged := &CallerOptions{}
	for _, opt := range opts {
		if opt.AccessToken != "" {
			merged.AccessToken = opt.AccessToken
		}
		if opt.AccessTokenSecret != "" {
			merged.AccessTokenSecret = opt.AccessTokenSecret
		}
		if opt.APIRootURL != "" {
			merged.APIRootURL = opt.APIRootURL
		}
		if opt.DefaultZone != "" {
			merged.DefaultZone = opt.DefaultZone
		}
		if len(opt.Zones) > 0 {
			merged.Zones = opt.Zones
		}
		if opt.AcceptLanguage != "" {
			merged.AcceptLanguage = opt.AcceptLanguage
		}
		if opt.HTTPClient != nil {
			merged.HTTPClient = opt.HTTPClient
		}
		if opt.HTTPRequestTimeout > 0 {
			merged.HTTPRequestTimeout = opt.HTTPRequestTimeout
		}
		if opt.HTTPRequestRateLimit > 0 {
			merged.HTTPRequestRateLimit = opt.HTTPRequestRateLimit
		}
		if opt.RetryMax > 0 {
			merged.RetryMax = opt.RetryMax
		}
		if opt.RetryWaitMax > 0 {
			merged.RetryWaitMax = opt.RetryWaitMax
		}
		if opt.RetryWaitMin > 0 {
			merged.RetryWaitMin = opt.RetryWaitMin
		}
		if opt.UserAgent != "" {
			merged.UserAgent = opt.UserAgent
		}

		// Note: bool値は一度trueにしたらMergeでfalseになることがない
		if opt.TraceAPI {
			merged.TraceAPI = true
		}
		if opt.TraceHTTP {
			merged.TraceHTTP = true
		}
		if opt.OpenTelemetry {
			merged.OpenTelemetry = true
		}
		if len(opt.OpenTelemetryOptions) > 0 {
			merged.OpenTelemetryOptions = opt.OpenTelemetryOptions
		}
		if opt.FakeMode {
			merged.FakeMode = true
		}
		if opt.FakeStorePath != "" {
			merged.FakeStorePath = opt.FakeStorePath
		}
	}
	return merged
}

// OptionsFromEnv 環境変数からCallerOptionsを組み立てて返す
func OptionsFromEnv() *CallerOptions {
	return &CallerOptions{
		AccessToken:       stringFromEnv("SAKURACLOUD_ACCESS_TOKEN", ""),
		AccessTokenSecret: stringFromEnv("SAKURACLOUD_ACCESS_TOKEN_SECRET", ""),

		APIRootURL:     stringFromEnv("SAKURACLOUD_API_ROOT_URL", ""),
		DefaultZone:    stringFromEnv("SAKURACLOUD_DEFAULT_ZONE", ""),
		Zones:          stringSliceFromEnv("SAKURACLOUD_ZONES", []string{}),
		AcceptLanguage: stringFromEnv("SAKURACLOUD_ACCEPT_LANGUAGE", ""),

		HTTPRequestTimeout:   intFromEnv("SAKURACLOUD_API_REQUEST_TIMEOUT", 0),
		HTTPRequestRateLimit: intFromEnv("SAKURACLOUD_API_REQUEST_RATE_LIMIT", 0),

		RetryMax:     intFromEnv("SAKURACLOUD_RETRY_MAX", 0),
		RetryWaitMax: intFromEnv("SAKURACLOUD_RETRY_WAIT_MAX", 0),
		RetryWaitMin: intFromEnv("SAKURACLOUD_RETRY_WAIT_MIN", 0),

		TraceAPI:  profile.EnableAPITrace(stringFromEnv("SAKURACLOUD_TRACE", "")),
		TraceHTTP: profile.EnableHTTPTrace(stringFromEnv("SAKURACLOUD_TRACE", "")),

		FakeMode:      os.Getenv("SAKURACLOUD_FAKE_MODE") != "",
		FakeStorePath: stringFromEnv("SAKURACLOUD_FAKE_STORE_PATH", ""),
	}
}

// OptionsFromProfile 指定のプロファイルからCallerOptionsを組み立てて返す
// プロファイル名に空文字が指定された場合はカレントプロファイルが利用される
func OptionsFromProfile(profileName string) (*CallerOptions, error) {
	if profileName == "" {
		current, err := profile.CurrentName()
		if err != nil {
			return nil, err
		}
		profileName = current
	}

	config := profile.ConfigValue{}
	if err := profile.Load(profileName, &config); err != nil {
		return nil, err
	}

	return &CallerOptions{
		AccessToken:          config.AccessToken,
		AccessTokenSecret:    config.AccessTokenSecret,
		APIRootURL:           config.APIRootURL,
		DefaultZone:          config.DefaultZone,
		Zones:                config.Zones,
		AcceptLanguage:       config.AcceptLanguage,
		HTTPRequestTimeout:   config.HTTPRequestTimeout,
		HTTPRequestRateLimit: config.HTTPRequestRateLimit,
		RetryMax:             config.RetryMax,
		RetryWaitMax:         config.RetryWaitMax,
		RetryWaitMin:         config.RetryWaitMin,
		TraceAPI:             config.EnableAPITrace(),
		TraceHTTP:            config.EnableHTTPTrace(),
		FakeMode:             config.FakeMode,
		FakeStorePath:        config.FakeStorePath,
	}, nil
}

func stringFromEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func stringFromEnvMulti(keys []string, defaultValue string) string {
	for _, key := range keys {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return defaultValue
}

func stringSliceFromEnv(key string, defaultValue []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	values := strings.Split(v, ",")
	for i := range values {
		values[i] = strings.Trim(values[i], " ")
	}
	return values
}

func intFromEnv(key string, defaultValue int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(i)
}
