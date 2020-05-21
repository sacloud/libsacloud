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

package iface

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type ConnectToPacketFilterRequest struct {
	Zone           string   `validate:"required" mapconv:"-"`
	ID             types.ID `validate:"required" mapconv:"-"`
	PacketFilterID types.ID `validate:"required"`
}

func (r *ConnectToPacketFilterRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) ConnectToPacketFilter(req *ConnectToPacketFilterRequest) error {
	return s.ConnectToPacketFilterWithContext(context.Background(), req)
}

func (s *Service) ConnectToPacketFilterWithContext(ctx context.Context, req *ConnectToPacketFilterRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewInterfaceOp(s.caller)
	return client.ConnectToPacketFilter(ctx, req.Zone, req.ID, req.PacketFilterID)
}
