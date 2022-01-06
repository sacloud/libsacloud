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

package dsl

import "fmt"

const (
	// DefaultPathFormat デフォルトのパスフォーマット
	DefaultPathFormat = "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}"
	// CloudAPISuffix IaaSリソースでのAPIサフィックス
	CloudAPISuffix = "api/cloud/1.1"
	// BillingAPISuffix 課金関連でのAPIサフィックス
	BillingAPISuffix = "api/system/1.0"
	// WebAccelAPISuffix ウェブアクセラレータ関連でのAPIサフィックス
	WebAccelAPISuffix = "api/webaccel/1.0"
)

var (
	// DefaultPathFormatWithID デフォルトのパス+IDのパスフォーマット
	DefaultPathFormatWithID = fmt.Sprintf("%s/{{.%s}}", DefaultPathFormat, ArgumentID.ArgName())
)

// IDAndSuffixPathFormat デフォルトのパス+ID+指定のサフィックスのパスフォーマット
func IDAndSuffixPathFormat(suffix string) string {
	return fmt.Sprintf(DefaultPathFormatWithID+"/%s", suffix)
}
