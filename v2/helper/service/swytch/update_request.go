// Copyright 2016-2022 The Libsacloud Authors
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

package swytch

import (
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `request:"-" validate:"required"`
	ID   types.ID `request:"-" validate:"required"`

	Name           *string     `request:",omitempty" validate:"omitempty,min=1"`
	Description    *string     `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags           *types.Tags `request:",omitempty"`
	IconID         *types.ID   `request:",omitempty"`
	NetworkMaskLen *int        `request:",omitempty"`
	DefaultRoute   *string     `request:",omitempty"`
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) ToRequestParameter(current *sacloud.Switch) (*sacloud.SwitchUpdateRequest, error) {
	r := &sacloud.SwitchUpdateRequest{}
	if err := service.RequestConvertTo(current, r); err != nil {
		return nil, err
	}
	if err := service.RequestConvertTo(req, r); err != nil {
		return nil, err
	}
	return r, nil
}
