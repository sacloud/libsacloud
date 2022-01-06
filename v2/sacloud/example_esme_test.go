// Copyright 2016-2022 The Libsacloud Authors
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

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func Example_sendSMSMessage() {
	// 2要素認証SMSの送信例

	// APIキー
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")
	destination := os.Getenv("SAKURACLOUD_ESME_DESTINATION") // 送信先電話番号 81 + 9012345678の形式で指定する
	if token == "" || secret == "" || destination == "" {
		log.Fatal("environment variable 'SAKURACLOUD_ACCESS_TOKEN'/'SAKURACLOUD_ACCESS_TOKEN_SECRET'/SAKURACLOUD_ESME_DESTINATION required")
	}

	// クライアントの作成
	caller := sacloud.NewClient(token, secret)
	esmeOp := sacloud.NewESMEOp(caller)

	// ESMEの作成(初回のみ必要)
	ctx := context.Background()
	esme, err := esmeOp.Create(ctx, &sacloud.ESMECreateRequest{
		Name:        "libsacloud-example",
		Description: "description",
		Tags:        types.Tags{"tag1", "tag2"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// SMS送信(OTPは自動生成の場合)
	result, err := esmeOp.SendMessageWithGeneratedOTP(ctx, esme.ID, &sacloud.ESMESendMessageWithGeneratedOTPRequest{
		Destination: destination,
		Sender:      "example-sender",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("OTP(result): %s\n", result.OTP)

	// OTPはログからも参照可能
	logs, err := esmeOp.Logs(ctx, esme.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OTP(logs): %s\n", logs[0].OTP)
}
