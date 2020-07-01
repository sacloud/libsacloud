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

package privatehostplan

import (
	"context"

	"github.com/sacloud/libsacloud/v2/pkg/size"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type FindRequest struct {
	Zone string `validate:"required" mapconv:"-"`

	CPU      int
	MemoryGB int

	*sacloud.FindCondition
}

func (r *FindRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) Find(req *FindRequest) ([]*sacloud.PrivateHostPlan, error) {
	return s.FindWithContext(context.Background(), req)
}

func (s *Service) FindWithContext(ctx context.Context, req *FindRequest) ([]*sacloud.PrivateHostPlan, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.FindCondition == nil {
		req.FindCondition = &sacloud.FindCondition{}
	}

	client := sacloud.NewPrivateHostPlanOp(s.caller)
	found, err := client.Find(ctx, req.Zone, req.FindCondition)
	if err != nil {
		return nil, err
	}

	var res []*sacloud.PrivateHostPlan
	for _, v := range found.PrivateHostPlans {
		if req.CPU > 0 && v.CPU != req.CPU {
			continue
		}
		if req.MemoryGB > 0 && v.MemoryMB != req.MemoryGB*size.GiB {
			continue
		}
		res = append(res, v)
	}

	return res, nil
}
