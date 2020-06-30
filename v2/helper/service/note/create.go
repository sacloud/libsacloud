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

package note

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type CreateRequest struct {
	Name    string `validate:"required"`
	Tags    types.Tags
	IconID  types.ID
	Class   string `validate:"required"`
	Content string `validate:"required"`
}

func (r *CreateRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) Create(req *CreateRequest) (*sacloud.Note, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*sacloud.Note, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	params := &sacloud.NoteCreateRequest{}
	if err := mapconv.ConvertFrom(req, params); err != nil {
		return nil, err
	}

	client := sacloud.NewNoteOp(s.caller)
	return client.Create(ctx, params)
}
