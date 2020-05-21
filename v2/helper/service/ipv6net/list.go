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

package ipv6net

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type ListRequest struct {
	Zone string `validate:"required" mapconv:"-"`
}

func (r *ListRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) List(req *ListRequest) ([]*sacloud.IPv6Net, error) {
	return s.ListWithContext(context.Background(), req)
}

func (s *Service) ListWithContext(ctx context.Context, req *ListRequest) ([]*sacloud.IPv6Net, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewIPv6NetOp(s.caller)
	found, err := client.List(ctx, req.Zone)
	if err != nil {
		return nil, err
	}
	return found.IPv6Nets, nil
}
