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

package disk

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

func TestCreateRequest_ToRequestParameter(t *testing.T) {
	in := &CreateRequest{SizeGB: 2}
	expect := &sacloud.DiskCreateRequest{
		SizeMB: size.GiBToMiB(2),
	}

	req, err := in.ToRequestParameter(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if req.SizeMB != expect.SizeMB {
		t.Error("ToRequestParameter returns unexpected value")
	}
}
