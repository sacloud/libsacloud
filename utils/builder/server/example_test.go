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

package server_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sacloud/libsacloud/v2/utils/builder/disk"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/builder/server"
)

func Example_builder() {
	// APIキー
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	// クライアントの作成
	client := sacloud.NewClient(token, secret)

	// ビルダーの準備
	builder := &server.Builder{
		Client:   server.NewBuildersAPIClient(client),
		Name:     "libsacloud-example",
		CPU:      2,
		MemoryGB: 4,
		// Commitment:      types.Commitments.Standard,
		// Generation:      types.PlanGenerations.Default,
		// InterfaceDriver: types.InterfaceDrivers.VirtIO,
		Description:     "description",
		Tags:            types.Tags{"tag1", "tag2"},
		BootAfterCreate: true,
		NIC:             &server.SharedNICSetting{
			// PacketFilterID: types.ID(123456789012),
		},
		//AdditionalNICs: []server.AdditionalNICSettingHolder{
		//	&server.ConnectedNICSetting{
		//		SwitchID:        types.ID(123456789012),
		//		DisplayIPAddress: "192.168.0.1",
		//		PacketFilterID:  types.ID(123456789012),
		//	},
		//},

		// IconID:          types.ID(123456789012),
		// CDROMID:         types.ID(123456789012),
		// PrivateHostID:   types.ID(123456789012),

		DiskBuilders: []disk.Builder{
			&disk.FromUnixBuilder{
				Client:     disk.NewBuildersAPIClient(client),
				OSType:     ostype.CentOS,
				Name:       "libsacloud-example",
				SizeGB:     20,
				PlanID:     types.DiskPlans.SSD,
				Connection: types.DiskConnections.VirtIO,
				//DistantFrom:     []types.ID{types.ID(123456789012)},
				Description: "description",
				Tags:        types.Tags{"tag1", "tag2"},
				EditParameter: &disk.UnixEditRequest{
					HostName:            "libsacloud-example", // ホスト名
					Password:            "P@ssW0rd",           // パスワード
					DisablePWAuth:       false,                // パスワード認証の無効化
					ChangePartitionUUID: true,                 // UUIDの変更
					EnableDHCP:          false,                // DHCPの有効化

					//IPAddress:                 "192.168.11.11",                    // IPアドレス(スイッチ or スイッチ+ルータに接続する場合)
					//NetworkMaskLen:            24,                                 // ネットワークマスク長(スイッチ or スイッチ+ルータに接続する場合)
					//DefaultRoute:              "192.168.11.1",                     // デフォルトルート(スイッチ or スイッチ+ルータに接続する場合)
					//SSHKeys:                   []string{sshKey1, sshKey2},         // 公開鍵(文字列で指定)
					//SSHKeyIDs:                 []types.ID{types.ID(123456789012)}, // 公開鍵(IDで指定)
					//IsSSHKeysEphemeral:        false,                              // 公開鍵をさくらのクラウド側で生成した場合にサーバ作成後に該当鍵の削除を行うか
					//GenerateSSHKeyName:        "",                                 // 生成する公開鍵の名称
					//GenerateSSHKeyPassPhrase:  "",                                 // 生成する公開鍵のパスフレーズ
					//GenerateSSHKeyDescription: "",                                 // 生成する公開鍵の説明
					//IsNotesEphemeral:          false,                              // Notesで指定したスタートアップスクリプトをサーバ作成後に削除するか
					//Notes:                     []string{note1, note2},             // スタートアップスクリプト(文字列で指定)
					//NoteIDs:                   []types.ID{types.ID(123456789012)}, // スタートアップスクリプト(IDで指定)
				},
				// IconID:          0,
			},
		},
	}

	result, err := builder.Build(context.Background(), "is1a")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ServerID: %s", result.ServerID)
}
