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

	"github.com/sacloud/libsacloud/v2/pkg/size"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Zone string `validate:"required" mapconv:"-"`
	Name string `validate:"required"`

	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID

	DiskPlanID      types.ID
	Connection      types.EDiskConnection
	SourceDiskID    types.ID
	SourceArchiveID types.ID
	ServerID        types.ID
	SizeGB          int

	DistantFrom []types.ID
}

func (r *CreateRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CreateRequest) toRequestParameter() (*sacloud.DiskCreateRequest, error) {
	req := &sacloud.DiskCreateRequest{}
	if err := mapconv.ConvertFrom(r, req); err != nil {
		return nil, err
	}
	req.SizeMB = r.SizeGB * size.GiB
	if req.DiskPlanID.IsEmpty() {
		req.DiskPlanID = types.DiskPlans.SSD
	}
	if req.Connection.String() == "" {
		req.Connection = types.DiskConnections.VirtIO
	}
	return req, nil
}

// Create ディスクの作成を行います。作成後は利用可能になるまでWaitForReady()などで待機する必要があります。
//
// SourceDiskID,SourceArchiveIDの両方を指定しないことでブランクディスクの作成が可能です。
func (s *Service) Create(req *CreateRequest) (*sacloud.Disk, error) {
	return s.CreateWithContext(context.Background(), req)
}

// CreateWithContext ディスクの作成を行います。作成後は利用可能になるまでWaitForReady()などで待機する必要があります。
//
// SourceDiskID,SourceArchiveIDの両方を指定しないことでブランクディスクの作成が可能です。
func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*sacloud.Disk, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	params, err := req.toRequestParameter()
	if err != nil {
		return nil, err
	}

	client := sacloud.NewDiskOp(s.caller)
	return client.Create(ctx, req.Zone, params, req.DistantFrom)
}
