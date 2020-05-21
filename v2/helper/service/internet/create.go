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

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type CreateRequest struct {
	Zone        string `validate:"required" mapconv:"-"`
	Name        string `validate:"required"`
	Description string `validate:"min=0,max=512"`

	Tags           types.Tags
	IconID         types.ID
	NetworkMaskLen int `validate:"omitempty,min=24,max=28"`
	BandWidthMbps  int `validate:"omitempty,oneof=100 250 500 1000 1500 2000 2500 3000 5000"` // TODO 将来的にコード生成で対応
}

func (r *CreateRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) Create(req *CreateRequest) (*sacloud.Internet, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*sacloud.Internet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if req.NetworkMaskLen == 0 {
		req.NetworkMaskLen = 28 // TODO デフォルト値の切り出し
	}
	if req.BandWidthMbps == 0 {
		req.BandWidthMbps = 100 // TODO デフォルト値の切り出し
	}

	params := &sacloud.InternetCreateRequest{}
	if err := mapconv.ConvertFrom(req, params); err != nil {
		return nil, err
	}

	client := sacloud.NewInternetOp(s.caller)
	return client.Create(ctx, req.Zone, params)
}
