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

package builder_test

import (
	"fmt"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/builder"
	"github.com/sacloud/libsacloud/sacloud/ostype"
)

func Example() {

	/**********************************************************************
	 * CentOSベースのサーバー作成例:
	 * NIC:共有セグメントに接続 / コア数:2 / メモリ: 4GB / ディスク: 100GB /
	 * 公開鍵を登録+パスワード認証無効化 / スタートアップスクリプトでyum update実施
	 *********************************************************************/

	token := "PUT-YOUR-TOKEN"       // APIトークン
	secret := "PUT-YOUR-SECRET"     // APIシークレット
	zone := "tk1a"                  // 対象ゾーン[ tk1a or is1b ]
	serverName := "example-server"  // サーバー名
	password := "PUT-YOUR-PASSWORD" // パスワード
	core := 2                       // コア数
	memory := 4                     // メモリ(GB)
	diskSize := 100                 // ディスクサイズ(GB)

	// SSH公開鍵
	sshKey := "ssh-rsa AAAA..."
	// スタートアップスクリプト
	script := `#!/bin/bash
yum -y update || exit 1
exit 0`

	//---------------------------------------------------------------------

	// APIクライアントの作成
	client := api.NewClient(token, secret, zone)

	// CentOSパブリックアーカイブからサーバー作成
	builder := builder.ServerPublicArchiveUnix(builder.NewAPIClient(client), ostype.CentOS, serverName, password)
	builder.AddPublicNWConnectedNIC() // NIC:共有セグメントに接続
	builder.SetCore(core)             // スペック指定(コア数)
	builder.SetMemory(memory)         // スペック指定(メモリ)
	builder.SetDiskSize(diskSize)     // スペック指定(ディスクサイズ)
	builder.AddSSHKey(sshKey)         // SSH公開鍵を登録
	builder.SetDisablePWAuth(true)    // パスワード認証を無効化
	builder.AddNote(script)           // スタートアップスクリプトを登録
	result, err := builder.Build()

	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", result.Server)

}
