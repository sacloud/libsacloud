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

	Clients []*ClientCert // Note: API的に証明書の削除はできないため、指定した以上の証明書が存在する可能性がある
	Servers []*ServerCert // Note: API的に証明書の削除はできないため、指定した以上の証明書が存在する可能性がある

	Client       sacloud.CertificateAuthorityAPI
	WaitDuration time.Duration // 証明書発行待ち時間、省略した場合10秒
}

// ClientCert クライアント証明書のリクエストパラメータ
type ClientCert struct {
	ID string // 新規登録時は空にする

	Country                   string
	Organization              string
	OrganizationUnit          []string
	CommonName                string
	NotAfter                  time.Time
	IssuanceMethod            string
	EMail                     string
	CertificateSigningRequest string
	PublicKey                 string
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

	shouldWait := false
	for _, cc := range b.Clients {
		// URLまたはeメールの場合は待つ必要なし
		if cc.IssuanceMethod == "public_key" || cc.IssuanceMethod == "csr" {
			shouldWait = true
		}
		_, err := b.Client.AddClient(ctx, created.ID, &sacloud.CertificateAuthorityAddClientParam{
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
			return nil, err
		}
	}
	for _, sc := range b.Servers {
		shouldWait = true

		_, err := b.Client.AddServer(ctx, created.ID, &sacloud.CertificateAuthorityAddServerParam{
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
			return nil, err
		}
	}

	if shouldWait {
		b.wait(ctx)
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

	shouldWait := false
	for _, cc := range b.Clients {
		if cc.ID != "" {
			continue
		}
		// URLまたはeメールの場合は待つ必要なし
		if cc.IssuanceMethod == "public_key" || cc.IssuanceMethod == "csr" {
			shouldWait = true
		}
		_, err := b.Client.AddClient(ctx, updated.ID, &sacloud.CertificateAuthorityAddClientParam{
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
			return nil, err
		}
	}
	for _, sc := range b.Servers {
		if sc.ID != "" {
			continue
		}

		shouldWait = true

		_, err := b.Client.AddServer(ctx, updated.ID, &sacloud.CertificateAuthorityAddServerParam{
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
			return nil, err
		}
	}

	if shouldWait {
		b.wait(ctx)
	}

	return read(ctx, b.Client, updated.ID)
}

func (b *Builder) wait(ctx context.Context) {
	// Note: 本来はクライアント/サーバ証明書ごとに状態を確認し証明書の発行が行われたか待つべきだが
	// 実装が煩雑となる & 現状だと数秒程度で発行されるため、数秒スリープすることで代わりとする
	wait := 10 * time.Second
	if b.WaitDuration > 0 {
		wait = b.WaitDuration
	}
	time.Sleep(wait)
}

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
	ca.Clients = clients.CertificateAuthority

	// servers
	servers, err := apiClient.ListServers(ctx, id)
	if err != nil {
		return nil, err
	}
	ca.Servers = servers.CertificateAuthority

	return ca, nil
}
