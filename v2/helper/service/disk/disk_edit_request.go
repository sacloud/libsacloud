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
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type EditRequest struct {
	HostName string
	Password string

	DisablePWAuth       bool
	EnableDHCP          bool
	ChangePartitionUUID bool

	IPAddress      string `validate:"omitempty,ipv4"`
	NetworkMaskLen int    `validate:"omitempty,min=1,max=32"`
	DefaultRoute   string `validate:"omitempty,ipv4"`

	SSHKeys   []string
	SSHKeyIDs []types.ID

	Notes []*sacloud.DiskEditNote
}

func (r *EditRequest) Validate() error {
	return validate.Struct(r)
}

func (r *EditRequest) toRequestParameter() *sacloud.DiskEditRequest {
	editReq := &sacloud.DiskEditRequest{
		Background:          true,
		Password:            r.Password,
		DisablePWAuth:       r.DisablePWAuth,
		EnableDHCP:          r.EnableDHCP,
		ChangePartitionUUID: r.ChangePartitionUUID,
		HostName:            r.HostName,
	}

	if r.IPAddress != "" {
		editReq.UserIPAddress = r.IPAddress
	}
	if r.NetworkMaskLen > 0 || r.DefaultRoute != "" {
		editReq.UserSubnet = &sacloud.DiskEditUserSubnet{
			NetworkMaskLen: r.NetworkMaskLen,
			DefaultRoute:   r.DefaultRoute,
		}
	}

	// ssh key
	var sshKeys []*sacloud.DiskEditSSHKey
	for _, key := range r.SSHKeys {
		sshKeys = append(sshKeys, &sacloud.DiskEditSSHKey{
			PublicKey: key,
		})
	}
	for _, id := range r.SSHKeyIDs {
		sshKeys = append(sshKeys, &sacloud.DiskEditSSHKey{
			ID: id,
		})
	}
	editReq.SSHKeys = sshKeys

	// startup script
	var notes []*sacloud.DiskEditNote
	for _, note := range r.Notes {
		notes = append(notes, &sacloud.DiskEditNote{
			ID:        note.ID,
			APIKeyID:  note.APIKeyID,
			Variables: note.Variables,
		})
	}
	editReq.Notes = notes

	return editReq
}
