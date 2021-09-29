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
	"time"

	"github.com/sacloud/libsacloud/v2/helper/wait"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Builder マネージドPKI(CA)のビルダー
type Builder struct {
	ID types.ID // 新規登録時は空にする

	Name        string
	Description string
	Tags        types.Tags
	IconID      types.ID

	Country          string
	Organization     string
	OrganizationUnit []string
	CommonName       string
	NotAfter         time.Time

	Clients []*ClientCert // Note: ここに指定しなかったものはRevokeされる
	Servers []*ServerCert // Note: ここに指定しなかったものはRevokeされる

	Client sacloud.CertificateAuthorityAPI

	PollingTimeout  time.Duration // 証明書発行待ちのタイムアウト
	PollingInterval time.Duration // 証明書発行待ちのポーリング間隔
}

// ClientCert クライアント証明書のリクエストパラメータ
type ClientCert struct {
	ID string // 新規登録時は空にする

	Country                   string
	Organization              string
	OrganizationUnit          []string
	CommonName                string
	NotAfter                  time.Time
	IssuanceMethod            types.ECertificateAuthorityIssuanceMethod
	EMail                     string
	CertificateSigningRequest string
	PublicKey                 string

	Hold bool // 一時停止する時はTrue
}

// ServerCert サーバ証明書のリクエストパラメータ
type ServerCert struct {
	ID string // 新規作成時は空にする

	Country                   string
	Organization              string
	OrganizationUnit          []string
	CommonName                string
	NotAfter                  time.Time
	SANs                      []string
	CertificateSigningRequest string
	PublicKey                 string

	Hold bool // 一時停止する時はTrue
}

// CertificateAuthority sacloud/CertificateAuthorityのラッパー
//
// CA/クライアント/サーバの詳細情報を保持する
type CertificateAuthority struct {
	sacloud.CertificateAuthority

	Detail  *sacloud.CertificateAuthorityDetail
	Clients []*sacloud.CertificateAuthorityClient
	Servers []*sacloud.CertificateAuthorityServer
}

func (b *Builder) Build(ctx context.Context) (*CertificateAuthority, error) {
	if b.ID.IsEmpty() {
		return b.create(ctx)
	}
	return b.update(ctx)
}

func (b *Builder) create(ctx context.Context) (*CertificateAuthority, error) {
	created, err := b.Client.Create(ctx, &sacloud.CertificateAuthorityCreateRequest{
		Name:             b.Name,
		Description:      b.Description,
		Tags:             b.Tags,
		IconID:           b.IconID,
		Country:          b.Country,
		Organization:     b.Organization,
		OrganizationUnit: b.OrganizationUnit,
		CommonName:       b.CommonName,
		NotAfter:         b.NotAfter,
	})
	if err != nil {
		return nil, err
	}
	err = b.wait(ctx, func() (bool, error) {
		detail, err := b.Client.Detail(ctx, created.ID)
		if err != nil {
			return false, err
		}
		return detail.CertificateData != nil, nil // CA自体の証明書が発行されるまでCertificateDataは空のことがある
	})
	if err != nil {
		return nil, err
	}

	if err := b.reconcileClients(ctx, created.ID); err != nil {
		return nil, err
	}
	if err := b.reconcileServers(ctx, created.ID); err != nil {
		return nil, err
	}

	return read(ctx, b.Client, created.ID)
}

func (b *Builder) update(ctx context.Context) (*CertificateAuthority, error) {
	updated, err := b.Client.Update(ctx, b.ID, &sacloud.CertificateAuthorityUpdateRequest{
		Name:        b.Name,
		Description: b.Description,
		Tags:        b.Tags,
		IconID:      b.IconID,
	})
	if err != nil {
		return nil, err
	}

	if err := b.reconcileClients(ctx, updated.ID); err != nil {
		return nil, err
	}
	if err := b.reconcileServers(ctx, updated.ID); err != nil {
		return nil, err
	}

	return read(ctx, b.Client, updated.ID)
}

func (b *Builder) reconcileClients(ctx context.Context, id types.ID) error {
	currentCerts, err := b.Client.ListClients(ctx, id)
	if err != nil {
		return err
	}

	if currentCerts != nil {
		// delete
		for _, target := range b.deletedClients(currentCerts.CertificateAuthority) {
			switch target.IssueState {
			case "available":
				if err := b.Client.RevokeClient(ctx, id, target.ID); err != nil {
					return err
				}
			case "approved":
				if err := b.Client.DenyClient(ctx, id, target.ID); err != nil {
					return err
				}
			}
		}

		// update
		for _, target := range b.updatedClients(currentCerts.CertificateAuthority) {
			if target.Hold {
				if err := b.Client.HoldClient(ctx, id, target.ID); err != nil {
					return err
				}
			} else {
				if err := b.Client.ResumeClient(ctx, id, target.ID); err != nil {
					return err
				}
			}
		}
	}
	// create
	for _, target := range b.createdClients() {
		if err := b.addClient(ctx, id, target); err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) deletedClients(currentClients []*sacloud.CertificateAuthorityClient) []*sacloud.CertificateAuthorityClient {
	var results []*sacloud.CertificateAuthorityClient
	for _, current := range currentClients {
		exists := false
		for _, desired := range b.Clients {
			if current.ID == "" || current.ID == desired.ID {
				exists = true
				break
			}
		}
		if !exists {
			results = append(results, current)
		}
	}
	return results
}

func (b *Builder) updatedClients(currentClients []*sacloud.CertificateAuthorityClient) []*ClientCert {
	var results []*ClientCert
	for _, current := range currentClients {
		for _, desired := range b.Clients {
			if current.ID == desired.ID {
				if (desired.Hold && current.IssueState == "available") || (!desired.Hold && current.IssueState == "hold") {
					results = append(results, desired)
				}
				break
			}
		}
	}
	return results
}

func (b *Builder) createdClients() []*ClientCert {
	var results []*ClientCert
	for _, created := range b.Clients {
		if created.ID == "" {
			results = append(results, created)
		}
	}
	return results
}

func (b *Builder) addClient(ctx context.Context, id types.ID, cc *ClientCert) error {
	cert, err := b.Client.AddClient(ctx, id, &sacloud.CertificateAuthorityAddClientParam{
		Country:                   cc.Country,
		Organization:              cc.Organization,
		OrganizationUnit:          cc.OrganizationUnit,
		CommonName:                cc.CommonName,
		NotAfter:                  cc.NotAfter,
		IssuanceMethod:            cc.IssuanceMethod,
		EMail:                     cc.EMail,
		CertificateSigningRequest: cc.CertificateSigningRequest,
		PublicKey:                 cc.PublicKey,
	})
	if err != nil {
		return err
	}
	cc.ID = cert.ID // 発行されたIDをBuilderに書き戻しておく

	// 証明書発行待ち、URLまたはEMailの場合は待たなくても良い(任意のURLにアクセスしWASMで.p12を発行するため)
	switch cc.IssuanceMethod {
	case types.CertificateAuthorityIssuanceMethods.CSR, types.CertificateAuthorityIssuanceMethods.PublicKey:
		err = b.wait(ctx, func() (bool, error) {
			c, err := b.Client.ReadClient(ctx, id, cert.ID)
			if err != nil {
				return false, err
			}
			return c.CertificateData != nil, nil
		})
		if err != nil {
			return err
		}
	}

	if cc.Hold {
		if err := b.Client.HoldClient(ctx, id, cert.ID); err != nil {
			return err
		}
	}
	return nil
}

func (b *Builder) reconcileServers(ctx context.Context, id types.ID) error {
	currentCerts, err := b.Client.ListServers(ctx, id)
	if err != nil {
		return err
	}

	if currentCerts != nil {
		// delete
		for _, target := range b.deletedServers(currentCerts.CertificateAuthority) {
			switch target.IssueState {
			case "available":
				if err := b.Client.RevokeServer(ctx, id, target.ID); err != nil {
					return err
				}
			}
		}

		// update
		for _, target := range b.updatedServers(currentCerts.CertificateAuthority) {
			if target.Hold {
				if err := b.Client.HoldServer(ctx, id, target.ID); err != nil {
					return err
				}
			} else {
				if err := b.Client.ResumeServer(ctx, id, target.ID); err != nil {
					return err
				}
			}
		}
	}
	// create
	for _, target := range b.createdServers() {
		if err := b.addServer(ctx, id, target); err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) deletedServers(currentServers []*sacloud.CertificateAuthorityServer) []*sacloud.CertificateAuthorityServer {
	var results []*sacloud.CertificateAuthorityServer
	for _, current := range currentServers {
		exists := false
		for _, desired := range b.Servers {
			if current.ID == "" || current.ID == desired.ID {
				exists = true
				break
			}
		}
		if !exists {
			results = append(results, current)
		}
	}
	return results
}

func (b *Builder) updatedServers(currentServers []*sacloud.CertificateAuthorityServer) []*ServerCert {
	var results []*ServerCert
	for _, current := range currentServers {
		for _, desired := range b.Servers {
			if current.ID == desired.ID {
				if (desired.Hold && current.IssueState == "available") || (!desired.Hold && current.IssueState == "hold") {
					results = append(results, desired)
				}
				break
			}
		}
	}
	return results
}

func (b *Builder) createdServers() []*ServerCert {
	var results []*ServerCert
	for _, created := range b.Servers {
		if created.ID == "" {
			results = append(results, created)
		}
	}
	return results
}

func (b *Builder) addServer(ctx context.Context, id types.ID, sc *ServerCert) error {
	cert, err := b.Client.AddServer(ctx, id, &sacloud.CertificateAuthorityAddServerParam{
		Country:                   sc.Country,
		Organization:              sc.Organization,
		OrganizationUnit:          sc.OrganizationUnit,
		CommonName:                sc.CommonName,
		NotAfter:                  sc.NotAfter,
		SANs:                      sc.SANs,
		CertificateSigningRequest: sc.CertificateSigningRequest,
		PublicKey:                 sc.PublicKey,
	})
	if err != nil {
		return err
	}
	sc.ID = cert.ID // 発行されたIDをBuilderに書き戻しておく

	err = b.wait(ctx, func() (bool, error) {
		c, err := b.Client.ReadServer(ctx, id, cert.ID)
		if err != nil {
			return false, err
		}
		return c.CertificateData != nil, nil
	})
	if err != nil {
		return err
	}

	if sc.Hold {
		if err := b.Client.HoldServer(ctx, id, cert.ID); err != nil {
			return err
		}
	}
	return nil
}

func (b *Builder) wait(ctx context.Context, readStateFunc wait.ReadStateFunc) error {
	timeout := b.PollingTimeout
	if timeout == time.Duration(0) {
		timeout = time.Minute // デフォルト: 5分
	}
	interval := b.PollingInterval
	if interval == time.Duration(0) {
		interval = 5 * time.Second
	}

	waiter := &wait.SimpleStateWaiter{
		ReadStateFunc:   readStateFunc,
		Timeout:         timeout,
		PollingInterval: interval,
	}

	_, err := waiter.WaitForState(ctx)
	return err
}
