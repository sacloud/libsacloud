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

package nfs

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/wait"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type WaitForDownRequest struct {
	Zone string   `validate:"required" mapconv:"-"`
	ID   types.ID `validate:"required" mapconv:"-"`
}

func (r *WaitForDownRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) WaitForDown(req *WaitForDownRequest) (*sacloud.NFS, error) {
	return s.WaitForDownWithContext(context.Background(), req)
}

func (s *Service) WaitForDownWithContext(ctx context.Context, req *WaitForDownRequest) (*sacloud.NFS, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	client := sacloud.NewNFSOp(s.caller)
	return wait.UntilNFSIsDown(ctx, client, req.Zone, req.ID)
}
