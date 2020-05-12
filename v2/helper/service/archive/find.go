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

package archive

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type FindRequest struct {
	Zone string `validate:"required" mapconv:"-"`

	Names  []string
	Tags   []string
	Scope  types.EScope
	OSType ostype.ArchiveOSType
	*sacloud.FindCondition
}

func (r *FindRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) Find(req *FindRequest) ([]*sacloud.Archive, error) {
	return s.FindWithContext(context.Background(), req)
}

func (s *Service) FindWithContext(ctx context.Context, req *FindRequest) ([]*sacloud.Archive, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.FindCondition == nil {
		req.FindCondition = &sacloud.FindCondition{}
	}
	filter, ok := ostype.ArchiveCriteria[req.OSType]
	if ok {
		for k, v := range filter {
			req.Filter[k] = v
		}
	}
	if len(req.Names) > 0 {
		req.Filter[search.Key("Name")] = search.PartialMatch(req.Names...)
	}
	if len(req.Tags) > 0 {
		req.Filter[search.Key("Tags.Name")] = search.TagsAndEqual(req.Tags...)
	}
	if req.Scope.String() != "" {
		req.Filter[search.Key("Scope")] = search.ExactMatch(req.Scope.String())
	}

	client := sacloud.NewArchiveOp(s.caller)
	found, err := client.Find(ctx, req.Zone, req.FindCondition)
	if err != nil {
		return nil, err
	}
	return found.Archives, nil
}
