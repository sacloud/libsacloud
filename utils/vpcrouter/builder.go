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
	*sacloud.VPCRouterCreateRequest
	AdditionalSwitches []*AdditionalSwitch // eth1-7に接続するスイッチのID
	BootAfterBuild     bool
	SetupOptions       *RetryableSetupParameter
}

// AdditionalSwitch VPCルータのeth1-eth7に接続するスイッチ
type AdditionalSwitch struct {
	ID    types.ID
	Index int
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

// Build VPCルータの作成、スイッチの接続をまとめて行う
func (b *Builder) Build(ctx context.Context, client sacloud.VPCRouterAPI, zone string) (*sacloud.VPCRouter, error) {
	b.init()

	planID := b.PlanID
	originalSettings := b.Settings.Interfaces

	// API仕様: 作成時にeth1以降のNICにスイッチを接続できない。
	// このため作成時はeth0のみパラメータとして渡す
	var creationTimeInterfaceSettings []*sacloud.VPCRouterInterfaceSetting

	// sharedの場合はnil
	if planID != types.VPCRouterPlans.Standard {
		for _, s := range originalSettings {
			if s.Index == 0 {
				creationTimeInterfaceSettings = []*sacloud.VPCRouterInterfaceSetting{s}
				break
			}
		}
	}

	builder := &setup.RetryableSetup{
		Create: func(ctx context.Context, zone string) (accessor.ID, error) {
			return client.Create(ctx, zone, &sacloud.VPCRouterCreateRequest{
				Name:        b.Name,
				Description: b.Description,
				Tags:        b.Tags,
				IconID:      b.IconID,
				PlanID:      b.PlanID,
				Switch:      b.Switch,
				IPAddresses: b.IPAddresses,
				Settings: &sacloud.VPCRouterSetting{
					VRID:                      b.Settings.VRID,
					InternetConnectionEnabled: b.Settings.InternetConnectionEnabled,
					Interfaces:                creationTimeInterfaceSettings,
					SyslogHost:                b.Settings.SyslogHost,
				},
			})
		},
		ProvisionBeforeUp: func(ctx context.Context, zone string, id types.ID, target interface{}) error {
			vpcRouter := target.(*sacloud.VPCRouter)
			// スイッチの接続
			for _, sw := range b.AdditionalSwitches {
				if err := client.ConnectToSwitch(ctx, zone, id, sw.Index, sw.ID); err != nil {
					return err
				}
			}
			// [HACK] スイッチ接続直後だとエラーになることがあるため数秒待つ
			time.Sleep(3 * time.Second)

			// 残りの設定の投入
			_, err := client.UpdateSettings(ctx, zone, id, &sacloud.VPCRouterUpdateSettingsRequest{
				Settings:     b.Settings,
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
