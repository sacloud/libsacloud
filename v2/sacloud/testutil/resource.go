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

package testutil

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

type ResourceTestCase struct {
	// PreCheck テスト実行 or スキップを判定するためのFunc
	PreCheck func(TestT)

	// APICallerのセットアップ用Func、テストケースごとに1回呼ばれる
	SetupAPICallerFunc func() sacloud.APICaller

	// Setup テスト前の準備(依存リソースの作成など)を行うためのFunc(省略可)
	Setup func(context.Context, sacloud.APICaller) error

	Tests []ResourceTestFunc

	// Cleanup APIで作成/変更したリソースなどのクリーンアップ用Func(省略化)
	Cleanup func(context.Context, sacloud.APICaller) error

	// Parallel t.Parallelを呼ぶかのフラグ
	Parallel bool

	Timeout time.Duration
}

type ResourceTestFunc func(ctx context.Context, caller sacloud.APICaller) error

func RunResource(t TestT, testCase *ResourceTestCase) {
	if testCase.SetupAPICallerFunc == nil {
		t.Fatal("CRUDTestCase.SetupAPICallerFunc is required")
	}

	if testCase.Parallel {
		t.Parallel()
	}

	if testCase.PreCheck != nil {
		testCase.PreCheck(t)
	}

	ctx := context.Background()
	if testCase.Timeout > 0 {
		withTimeout, cancel := context.WithTimeout(ctx, testCase.Timeout)
		defer cancel()
		ctx = withTimeout
	}

	defer func() {
		// Cleanup
		if testCase.Cleanup != nil {
			if err := testCase.Cleanup(ctx, testCase.SetupAPICallerFunc()); err != nil {
				t.Logf("cleanup failed: ", err)
			}
		}
		if err := recover(); err != nil {
			t.Logf("unexpected error has happen: %v, trace: %s", err, string(debug.Stack()))
		}
	}()

	if testCase.Setup != nil {
		if err := testCase.Setup(ctx, testCase.SetupAPICallerFunc()); err != nil {
			t.Error("setup failed: ", err)
			return
		}
	}

	for _, testFunc := range testCase.Tests {
		if err := testFunc(ctx, testCase.SetupAPICallerFunc()); err != nil {
			t.Error("test func returns error: ", err)
			return
		}
	}
}
