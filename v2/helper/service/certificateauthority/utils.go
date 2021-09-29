// Copyright 2016-2021 The Libsacloud Authors
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

package certificateauthority

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func read(ctx context.Context, apiClient sacloud.CertificateAuthorityAPI, id types.ID) (*CertificateAuthority, error) {
	current, err := apiClient.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	ca := &CertificateAuthority{CertificateAuthority: *current}

	// detail
	detail, err := apiClient.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	ca.Detail = detail

	// clients
	clients, err := apiClient.ListClients(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, c := range clients.CertificateAuthority {
		if c.IssueState != "revoked" && c.IssueState != "deny" {
			ca.Clients = append(ca.Clients, c)
		}
	}

	// servers
	servers, err := apiClient.ListServers(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, s := range servers.CertificateAuthority {
		if s.IssueState != "revoked" && s.IssueState != "deny" {
			ca.Servers = append(ca.Servers, s)
		}
	}

	return ca, nil
}
