// Copyright 2016-2019 The Libsacloud Authors
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

package query

import (
	"context"
	"net/http"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type dummyArchiveFinder struct {
	archive *sacloud.ArchiveFindResult
	err     error
}

func (f *dummyArchiveFinder) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ArchiveFindResult, error) {
	return f.archive, f.err
}

type dummyNoteFinder struct {
	notes []*sacloud.Note
	err   error
}

func (f *dummyNoteFinder) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.NoteFindResult, error) {
	if f.err != nil {
		return nil, f.err
	}

	return &sacloud.NoteFindResult{
		Total: len(f.notes),
		Count: len(f.notes),
		Notes: f.notes,
	}, nil
}

type dummyServerPlanFinder struct {
	plans []*sacloud.ServerPlan
	err   error
}

func (f *dummyServerPlanFinder) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServerPlanFindResult, error) {
	if f.err != nil {
		return nil, f.err
	}

	return &sacloud.ServerPlanFindResult{
		Total:       len(f.plans),
		Count:       len(f.plans),
		ServerPlans: f.plans,
	}, nil
}

type dummyServerReader struct {
	servers []*sacloud.Server
	err     error
}

func (r *dummyServerReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Server, error) {
	if r.err != nil {
		return nil, r.err
	}
	for _, s := range r.servers {
		if s.ID == id {
			return s, nil
		}
	}
	return nil, sacloud.NewAPIError(http.MethodGet, nil, "", http.StatusNotFound, nil)
}

type dummyArchiveReader struct {
	archives []*sacloud.Archive
	err      error
}

func (r *dummyArchiveReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Archive, error) {
	if r.err != nil {
		return nil, r.err
	}
	for _, a := range r.archives {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, sacloud.NewAPIError(http.MethodGet, nil, "", http.StatusNotFound, nil)
}

type dummyDiskReader struct {
	disks []*sacloud.Disk
	err   error
}

func (r *dummyDiskReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Disk, error) {
	if r.err != nil {
		return nil, r.err
	}
	for _, d := range r.disks {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, sacloud.NewAPIError(http.MethodGet, nil, "", http.StatusNotFound, nil)
}
