package builder_test

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/api"
	"github.com/yamamoto-febc/libsacloud/builder"
	"github.com/yamamoto-febc/libsacloud/sacloud/ostype"
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
	sshKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDVbJLAQHDVpgsjLhauPl1dY5o5MeC1f+sPQW1W5D9Iug+qCdUI3VpWSq5oPSe4sx4n+l3eFbEsjA6Z2pDwboBwZ142P5Ry5npiIX1Bi8xbx3Cp8KylgILf+pJtFqkRMdpFEDPxN2cmqsSR4yPyMJ8R5sBPMFRtBOkBLiRrdfLBOoh4RwpS3tiNwqkLCc2YVirLL+NTz6/1shQu8++hO0xWDjzvrl/plIAHpVG8nuPuMr9zE+MPW3m+1O0jV9iFFh8/8vTnfP1kPY/CQCht05Lh+q53XKXp0a7tdKRzJ6TKV6l5VfySKIIcuKSVJ16ysbnbYMacc2mEsH5DAXxlPESl"
	// スタートアップスクリプト
	script := `#!/bin/bash
yum -y update || exit 1
exit 0`

	//---------------------------------------------------------------------

	// APIクライアントの作成
	client := api.NewClient(token, secret, zone)

	// CentOSパブリックアーカイブからサーバー作成
	result, err := builder.FromPublicArchiveUnix(client, ostype.CentOS, serverName, password).
		AddPublicNWConnectedNIC(). // NIC:共有セグメントに接続
		SetCore(core).             // スペック指定(コア数)
		SetMemory(memory).         // スペック指定(メモリ)
		SetDiskSize(diskSize).     // スペック指定(ディスクサイズ)
		AddSSHKey(sshKey).         // SSH公開鍵を登録
		SetDisablePWAuth(true).    // パスワード認証を無効化
		AddNote(script).           // スタートアップスクリプトを登録
		Build()                    // 構築実施

	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", result.Server)

}
