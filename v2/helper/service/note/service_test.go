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

package note

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestNoteService_CRUD(t *testing.T) {
	prefix := testutil.RandomPrefix()
	name := prefix + "note-service"

	var svc *Service
	var note *sacloud.Note

	testutil.RunResource(t, &testutil.ResourceTestCase{
		SetupAPICallerFunc: testutil.SingletonAPICaller,
		Setup: func(ctx context.Context, caller sacloud.APICaller) error {
			svc = New(caller)
			return nil
		},
		Tests: []testutil.ResourceTestFunc{
			// create
			func(ctx context.Context, caller sacloud.APICaller) error {
				created, err := svc.Create(&CreateRequest{
					Name:    name,
					Tags:    types.Tags{"tag1", "tag2"},
					Class:   "shell",
					Content: "#!/bin/bash",
				})
				if err != nil {
					return err
				}
				note = created
				return nil
			},
			// update
			func(ctx context.Context, caller sacloud.APICaller) error {
				updated, err := svc.Update(&UpdateRequest{
					ID:   note.ID,
					Name: pointer.NewString(name + "-upd"),
					Tags: pointer.NewTags(types.Tags{"tag1-upd", "tag2-upd"}),
				})
				if err != nil {
					return err
				}
				note = updated
				return nil
			},
			// delete
			func(ctx context.Context, caller sacloud.APICaller) error {
				return svc.Delete(&DeleteRequest{ID: note.ID})
			},
		},
		Cleanup:  testutil.ComposeCleanupResourceFunc(prefix, testutil.CleanupTargets.Note),
		Parallel: true,
	})
}
