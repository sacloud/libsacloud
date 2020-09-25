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

package containerregistry

import (
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type ApplyRequest struct {
	ID types.ID `request:"-"`

	Name           string `validate:"required"`
	Description    string `validate:"min=0,max=512"`
	Tags           types.Tags
	IconID         types.ID
	AccessLevel    types.EContainerRegistryAccessLevel
	VirtualDomain  string
	SubDomainLabel string
	Users          []*User

	SettingsHash string
}

func (req *ApplyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ApplyRequest) Builder(caller sacloud.APICaller) (*Builder, error) {
	return &Builder{
		Name:           req.Name,
		Description:    req.Description,
		Tags:           req.Tags,
		IconID:         req.IconID,
		AccessLevel:    req.AccessLevel,
		VirtualDomain:  req.VirtualDomain,
		SubDomainLabel: req.SubDomainLabel,
		Users:          req.Users,
		SettingsHash:   req.SettingsHash,
		Client:         sacloud.NewContainerRegistryOp(caller),
	}, nil
}
