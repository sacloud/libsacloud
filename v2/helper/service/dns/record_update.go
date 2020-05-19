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

package dns

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRecordRequest struct {
	ID      types.ID           `validate:"required" mapconv:"-"`
	Current *sacloud.DNSRecord `validate:"required" mapconv:"-"`

	Name  *string
	Type  *types.EDNSRecordType `validate:"omitempty,dns_record_type"`
	RData *string
	TTL   *int
}

func (r *UpdateRecordRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) UpdateRecord(req *UpdateRecordRequest) error {
	return s.UpdateRecordWithContext(context.Background(), req)
}

func (s *Service) UpdateRecordWithContext(ctx context.Context, req *UpdateRecordRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewDNSOp(s.caller)
	dns, err := client.Read(ctx, req.ID)
	if err != nil {
		return err
	}

	current := dns.FindRecord(req.Current.Name, req.Current.Type, req.Current.RData)
	if current == nil {
		return fmt.Errorf("no target record: %v", current)
	}
	if err := mapconv.ConvertFrom(req, current); err != nil {
		return err
	}

	_, err = client.UpdateSettings(ctx, dns.ID, &sacloud.DNSUpdateSettingsRequest{
		Records:      dns.Records,
		SettingsHash: dns.SettingsHash,
	})
	return err
}
