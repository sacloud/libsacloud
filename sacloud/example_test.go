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

package sacloud_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func Example_basic() {
	// APIクライアントの基本的な使い方の例

	// APIキー
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	// クライアントの作成
	client := sacloud.NewClient(token, secret)

	// スイッチの作成
	swOp := sacloud.NewSwitchOp(client)
	sw, err := swOp.Create(context.Background(), "is1a", &sacloud.SwitchCreateRequest{
		Name:        "libsacloud-example",
		Description: "description",
		Tags:        types.Tags{"tag1", "tag2"},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s", sw.Name)
}

func Example_serverCRUD() {
	// ServerのCRUDを行う例

	// Note: サーバの作成を行いたい場合、通常はgithub.com/libsacloud/v2/utils/serverパッケージを利用してください。
	// この例はServer APIを直接利用したい場合向けです。

	// APIキー
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	// クライアントの作成
	client := sacloud.NewClient(token, secret)

	// サーバの作成(ディスクレス)
	ctx := context.Background()
	serverOp := sacloud.NewServerOp(client)
	server, err := serverOp.Create(ctx, "is1a", &sacloud.ServerCreateRequest{
		CPU:                  1,
		MemoryMB:             1 * size.GiB,
		ServerPlanCommitment: types.Commitments.Standard,
		ServerPlanGeneration: types.PlanGenerations.Default,
		ConnectedSwitches:    []*sacloud.ConnectedSwitch{{Scope: "shared"}},
		InterfaceDriver:      types.InterfaceDrivers.VirtIO,
		Name:                 "libsacloud-example",
		Description:          "description",
		Tags:                 types.Tags{"tag1", "tag2"},
		//IconID:               0,
		WaitDiskMigration: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 更新
	server, err = serverOp.Update(ctx, "is1a", server.ID, &sacloud.ServerUpdateRequest{
		Name:        "libsacloud-example-updated",
		Description: "description-updated",
		Tags:        types.Tags{"tag1-updated", "tag2-updated"},
		// IconID:      0,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 起動
	if err := serverOp.Boot(ctx, "is1a", server.ID); err != nil {
		log.Fatal(err)
	}

	// runningになるまで待機(同期)
	waiter := sacloud.WaiterForUp(func() (state interface{}, err error) {
		return serverOp.Read(ctx, "is1a", server.ID)
	})
	if _, err := waiter.WaitForState(ctx); err != nil {
		log.Fatal(err)
	}
	// 非同期の場合
	// completeCh, progressCh, errCh := waiter.AsyncWaitForState(ctx)

	// シャットダウン(force)
	if err := serverOp.Shutdown(ctx, "is1a", server.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
		log.Fatal(err)
	}

	// Downになるまで待機
	waiter = sacloud.WaiterForDown(func() (state interface{}, err error) {
		return serverOp.Read(ctx, "is1a", server.ID)
	})
	if _, err := waiter.WaitForState(ctx); err != nil {
		log.Fatal(err)
	}

	// 削除
	if err := serverOp.Delete(ctx, "is1a", server.ID); err != nil {
		log.Fatal(err)
	}

	fmt.Println("hoge")
}
