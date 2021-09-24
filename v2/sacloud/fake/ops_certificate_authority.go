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

package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *CertificateAuthorityOp) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.CertificateAuthorityFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.CertificateAuthority
	for _, res := range results {
		dest := &sacloud.CertificateAuthority{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.CertificateAuthorityFindResult{
		Total:                  len(results),
		Count:                  len(results),
		From:                   0,
		CertificateAuthorities: values,
	}, nil
}

// Create is fake implementation
func (o *CertificateAuthorityOp) Create(ctx context.Context, param *sacloud.CertificateAuthorityCreateRequest) (*sacloud.CertificateAuthority, error) {
	result := &sacloud.CertificateAuthority{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)
	result.Availability = types.Availabilities.Available

	putCertificateAuthority(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *CertificateAuthorityOp) Read(ctx context.Context, id types.ID) (*sacloud.CertificateAuthority, error) {
	value := getCertificateAuthorityByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.CertificateAuthority{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *CertificateAuthorityOp) Update(ctx context.Context, id types.ID, param *sacloud.CertificateAuthorityUpdateRequest) (*sacloud.CertificateAuthority, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	return value, nil
}

// Delete is fake implementation
func (o *CertificateAuthorityOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}

func (o *CertificateAuthorityOp) Detail(ctx context.Context, id types.ID) (*sacloud.CertificateAuthorityDetail, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return &sacloud.CertificateAuthorityDetail{
		Subject: "fake",
		CertificateData: &sacloud.CertificateData{
			CertificatePEM: "...",
			Subject:        "fake",
			SerialNumber:   "...",
			NotBefore:      time.Time{},
			NotAfter:       time.Time{},
		},
	}, nil
}

func (o *CertificateAuthorityOp) ListClients(ctx context.Context, id types.ID) (*sacloud.CertificateAuthorityListClientsResult, error) {
	return nil, fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) ReadClient(ctx context.Context, id types.ID, clientID string) (*sacloud.CertificateAuthorityClient, error) {
	return nil, fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) HoldClient(ctx context.Context, id types.ID, clientID string) error {
	return fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) ResumeClient(ctx context.Context, id types.ID, clientID string) error {
	return fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) RevokeClient(ctx context.Context, id types.ID, clientID string) error {
	return fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) DenyClient(ctx context.Context, id types.ID, clientID string) error {
	return fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) ListServers(ctx context.Context, id types.ID) (*sacloud.CertificateAuthorityListServersResult, error) {
	return nil, fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) ReadServer(ctx context.Context, id types.ID, ServerID string) (*sacloud.CertificateAuthorityServer, error) {
	return nil, fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) HoldServer(ctx context.Context, id types.ID, ServerID string) error {
	return fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) ResumeServer(ctx context.Context, id types.ID, ServerID string) error {
	return fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) RevokeServer(ctx context.Context, id types.ID, ServerID string) error {
	return fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) AddClient(ctx context.Context, id types.ID, param *sacloud.CertificateAuthorityAddClientParam) (*sacloud.CertificateAuthorityAddClientOrServerResult, error) {
	return nil, fmt.Errorf("not supported")
}

func (o *CertificateAuthorityOp) AddServer(ctx context.Context, id types.ID, param *sacloud.CertificateAuthorityAddServerParam) (*sacloud.CertificateAuthorityAddClientOrServerResult, error) {
	return nil, fmt.Errorf("not supported")
}
