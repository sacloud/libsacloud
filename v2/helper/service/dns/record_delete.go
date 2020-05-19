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

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type DeleteRecordRequest struct {
	ID      types.ID           `validate:"required" mapconv:"-"`
	Current *sacloud.DNSRecord `validate:"required" mapconv:"-"`
}

func (r *DeleteRecordRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) DeleteRecord(req *DeleteRecordRequest) error {
	return s.DeleteRecordWithContext(context.Background(), req)
}

func (s *Service) DeleteRecordWithContext(ctx context.Context, req *DeleteRecordRequest) error {
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
		return nil // noop if not exists
	}

	dns.RemoveRecords(req.Current)
	_, err = client.UpdateSettings(ctx, dns.ID, &sacloud.DNSUpdateSettingsRequest{
		Records:      dns.Records,
		SettingsHash: dns.SettingsHash,
	})
	return err
}
