// Copyright 2016-2019 The Libsacloud Authors
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

package vpcrouter

import (
	"context"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/setup"
)

// Builder VPCルータの構築を行う
type Builder struct {
	Name                  string
	Description           string
	Tags                  types.Tags
	IconID                types.ID
	PlanID                types.ID
	NICSetting            NICSettingHolder
	AdditionalNICSettings []AdditionalNICSettingHolder
	RouterSetting         *RouterSetting

	BootAfterBuild bool
	SetupOptions   *RetryableSetupParameter
}

// RouterSetting VPCルータの設定
type RouterSetting struct {
	VRID                      int
	InternetConnectionEnabled types.StringFlag
	StaticNAT                 []*sacloud.VPCRouterStaticNAT
	PortForwarding            []*sacloud.VPCRouterPortForwarding
	Firewall                  []*sacloud.VPCRouterFirewall
	DHCPServer                []*sacloud.VPCRouterDHCPServer
	DHCPStaticMapping         []*sacloud.VPCRouterDHCPStaticMapping
	PPTPServer                *sacloud.VPCRouterPPTPServer
	L2TPIPsecServer           *sacloud.VPCRouterL2TPIPsecServer
	RemoteAccessUsers         []*sacloud.VPCRouterRemoteAccessUser
	SiteToSiteIPsecVPN        []*sacloud.VPCRouterSiteToSiteIPsecVPN
	StaticRoute               []*sacloud.VPCRouterStaticRoute
	SyslogHost                string
}

// RetryableSetupParameter VPCルータ作成時に利用するsetup.RetryableSetupのパラメータ
type RetryableSetupParameter struct {
	// RetryCount リトライ回数
	RetryCount int
	// DeleteRetryCount 削除リトライ回数
	DeleteRetryCount int
	// DeleteRetryInterval 削除リトライ間隔
	DeleteRetryInterval time.Duration
	// sacloud.StateWaiterによるステート待ちの間隔
	PollInterval time.Duration
}

func (b *Builder) init() {
	if b.SetupOptions == nil {
		b.SetupOptions = &RetryableSetupParameter{}
	}
}

func (b *Builder) getInitInterfaceSettings() []*sacloud.VPCRouterInterfaceSetting {
	s := b.NICSetting.getInterfaceSetting()
	if s != nil {
		return []*sacloud.VPCRouterInterfaceSetting{s}
	}
	return nil
}

func (b *Builder) getInterfaceSettings() []*sacloud.VPCRouterInterfaceSetting {
	var settings []*sacloud.VPCRouterInterfaceSetting
	if s := b.NICSetting.getInterfaceSetting(); s != nil {
		settings = append(settings, s)
	}
	for _, additionalNIC := range b.AdditionalNICSettings {
		settings = append(settings, additionalNIC.getInterfaceSetting())
	}
	return settings
}

// Build VPCルータの作成、スイッチの接続をまとめて行う
func (b *Builder) Build(ctx context.Context, client sacloud.VPCRouterAPI, zone string) (*sacloud.VPCRouter, error) {
	b.init()

	builder := &setup.RetryableSetup{
		Create: func(ctx context.Context, zone string) (accessor.ID, error) {
			return client.Create(ctx, zone, &sacloud.VPCRouterCreateRequest{
				Name:        b.Name,
				Description: b.Description,
				Tags:        b.Tags,
				IconID:      b.IconID,
				PlanID:      b.PlanID,
				Switch:      b.NICSetting.getConnectedSwitch(),
				IPAddresses: b.NICSetting.getIPAddresses(),
				Settings: &sacloud.VPCRouterSetting{
					VRID:                      b.RouterSetting.VRID,
					InternetConnectionEnabled: b.RouterSetting.InternetConnectionEnabled,
					Interfaces:                b.getInitInterfaceSettings(),
					SyslogHost:                b.RouterSetting.SyslogHost,
				},
			})
		},
		ProvisionBeforeUp: func(ctx context.Context, zone string, id types.ID, target interface{}) error {
			vpcRouter := target.(*sacloud.VPCRouter)

			// スイッチの接続
			for _, additionalNIC := range b.AdditionalNICSettings {
				switchID, index := additionalNIC.getSwitchInfo()
				if err := client.ConnectToSwitch(ctx, zone, id, index, switchID); err != nil {
					return err
				}
			}

			// [HACK] スイッチ接続直後だとエラーになることがあるため数秒待つ
			time.Sleep(3 * time.Second)

			// 残りの設定の投入
			_, err := client.UpdateSettings(ctx, zone, id, &sacloud.VPCRouterUpdateSettingsRequest{
				Settings: &sacloud.VPCRouterSetting{
					VRID:                      b.RouterSetting.VRID,
					InternetConnectionEnabled: b.RouterSetting.InternetConnectionEnabled,
					Interfaces:                b.getInterfaceSettings(),
					StaticNAT:                 b.RouterSetting.StaticNAT,
					PortForwarding:            b.RouterSetting.PortForwarding,
					Firewall:                  b.RouterSetting.Firewall,
					DHCPServer:                b.RouterSetting.DHCPServer,
					DHCPStaticMapping:         b.RouterSetting.DHCPStaticMapping,
					PPTPServer:                b.RouterSetting.PPTPServer,
					PPTPServerEnabled:         b.RouterSetting.PPTPServer != nil,
					L2TPIPsecServer:           b.RouterSetting.L2TPIPsecServer,
					L2TPIPsecServerEnabled:    b.RouterSetting.L2TPIPsecServer != nil,
					RemoteAccessUsers:         b.RouterSetting.RemoteAccessUsers,
					SiteToSiteIPsecVPN:        b.RouterSetting.SiteToSiteIPsecVPN,
					StaticRoute:               b.RouterSetting.StaticRoute,
					SyslogHost:                b.RouterSetting.SyslogHost,
				},
				SettingsHash: vpcRouter.SettingsHash,
			})
			if err != nil {
				return err
			}

			if b.BootAfterBuild {
				return client.Boot(ctx, zone, id)
			}
			return nil
		},
		Delete: func(ctx context.Context, zone string, id types.ID) error {
			return client.Delete(ctx, zone, id)
		},
		Read: func(ctx context.Context, zone string, id types.ID) (interface{}, error) {
			return client.Read(ctx, zone, id)
		},
		IsWaitForCopy:          true,
		IsWaitForUp:            b.BootAfterBuild,
		RetryCount:             b.SetupOptions.RetryCount,
		ProvisioningRetryCount: 1,
		DeleteRetryCount:       b.SetupOptions.DeleteRetryCount,
		DeleteRetryInterval:    b.SetupOptions.DeleteRetryInterval,
		PollInterval:           b.SetupOptions.PollInterval,
	}

	result, err := builder.Setup(ctx, zone)
	if err != nil {
		return nil, err
	}
	return result.(*sacloud.VPCRouter), nil
}
