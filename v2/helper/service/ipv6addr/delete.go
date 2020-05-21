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

package ipv6addr

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/service"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type DeleteRequest struct {
	Zone     string `validate:"required" mapconv:"-"`
	IPv6Addr string `validate:"required,ipv6" mapconv:"-"`

	FailIfNotFound bool
}

func (r *DeleteRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) Delete(req *DeleteRequest) error {
	return s.DeleteWithContext(context.Background(), req)
}

func (s *Service) DeleteWithContext(ctx context.Context, req *DeleteRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	client := sacloud.NewIPv6AddrOp(s.caller)
	err := client.Delete(ctx, req.Zone, req.IPv6Addr)
	if err != nil {
		return service.HandleNotFoundError(err, !req.FailIfNotFound)
	}
	return nil
}
