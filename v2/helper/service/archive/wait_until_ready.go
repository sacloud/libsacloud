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
	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type WaitUntilReadyRequest struct {
	Zone string   `validate:"required" mapconv:"-"`
	ID   types.ID `validate:"required" mapconv:"-"`
}

func (r *WaitUntilReadyRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) WaitUntilReady(req *WaitUntilReadyRequest) error {
	return s.WaitUntilReadyWithContext(context.Background(), req)
}

func (s *Service) WaitUntilReadyWithContext(ctx context.Context, req *WaitUntilReadyRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewArchiveOp(s.caller)
	_, err := wait.UntilArchiveIsReady(ctx, client, req.Zone, req.ID)
	return err
}
