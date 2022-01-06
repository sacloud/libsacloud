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

package server

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) WaitBoot(req *WaitBootRequest) error {
	return s.WaitBootWithContext(context.Background(), req)
}

func (s *Service) WaitBootWithContext(ctx context.Context, req *WaitBootRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewServerOp(s.caller)
	_, err := wait.UntilServerIsUp(ctx, client, req.Zone, req.ID)
	return err
}
