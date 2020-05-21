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

package internet

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type ListIPv6Request struct {
	Zone string   `validate:"required" mapconv:"-"`
	ID   types.ID `validate:"required" mapconv:"-"`
}

func (r *ListIPv6Request) Validate() error {
	return validate.Struct(r)
}

func (s *Service) ListIPv6(req *ListIPv6Request) ([]*sacloud.IPv6Net, error) {
	return s.ListIPv6WithContext(context.Background(), req)
}

func (s *Service) ListIPv6WithContext(ctx context.Context, req *ListIPv6Request) ([]*sacloud.IPv6Net, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	internetOp := sacloud.NewInternetOp(s.caller)
	current, err := internetOp.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, fmt.Errorf("reading the internet resource[%s] failed: %s", req.ID, err)
	}

	ipv6Op := sacloud.NewIPv6NetOp(s.caller)
	var results []*sacloud.IPv6Net
	for _, net := range current.Switch.IPv6Nets {
		sn, err := ipv6Op.Read(ctx, req.Zone, net.ID)
		if err != nil {
			return nil, err
		}
		results = append(results, sn)
	}
	return results, nil
}
