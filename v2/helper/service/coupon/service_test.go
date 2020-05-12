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

package coupon

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
)

func TestService_List(t *testing.T) {
	service := New(testutil.SingletonAPICaller())
	coupon, err := service.List()
	if err != nil {
		t.Fatal(err)
	}
	if coupon == nil {
		t.Fatal("no results")
	}
}
