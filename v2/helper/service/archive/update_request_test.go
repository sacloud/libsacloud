// Copyright 2016-2021 The Libsacloud Authors
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

package archive

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/testutil"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestUpdateRequest(t *testing.T) {
	archive := &sacloud.Archive{
		ID:          1,
		Name:        "hoge",
		Description: "fuga",
		Tags:        types.Tags{"tag1", "tag2"},
	}

	updateRequest := &UpdateRequest{
		Zone: "is1a",
		ID:   1,
		Name: pointer.NewString(""),
		//Description: pointer.NewString(""), // 未指定パラメータは元の値を保持(request:,omitemptyが必要)
		Tags: pointer.NewTags(types.Tags{}),
	}

	result, err := updateRequest.ToRequestParameter(archive)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertEmpty(t, result.Name, "Name")                       // nolint
	testutil.AssertEqual(t, "fuga", result.Description, "Description") // nolint
	testutil.AssertEmpty(t, result.Tags, "Tags")                       // nolint
}
