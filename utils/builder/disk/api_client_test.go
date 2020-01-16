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

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type dummyDiskPlanReader struct {
	diskPlan *sacloud.DiskPlan
	err      error
}

func (d *dummyDiskPlanReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.DiskPlan, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.diskPlan, nil
}

type dummyNoteHandler struct {
	note *sacloud.Note
	err  error
}

func (d *dummyNoteHandler) Read(ctx context.Context, id types.ID) (*sacloud.Note, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.note, nil
}

func (d *dummyNoteHandler) Create(ctx context.Context, param *sacloud.NoteCreateRequest) (*sacloud.Note, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.note, nil
}

func (d *dummyNoteHandler) Delete(ctx context.Context, id types.ID) error {
	return d.err
}
