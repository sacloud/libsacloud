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

package search_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
)

func Example() {

	// API Keys
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")
	zone := os.Getenv("SAKURACLOUD_ZONE")

	// API Client
	caller := sacloud.NewClient(token, secret)
	serverOp := sacloud.NewArchiveOp(caller)

	// ******************************************
	// Find
	// ******************************************

	// 名称に"Example"を含むサーバを検索
	condition := &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key("Name"): search.PartialMatch("Example"),
		},
	}
	searched, err := serverOp.Find(context.Background(), zone, condition)
	if err != nil {
		panic(err)
	}

	fmt.Printf("searched: %#v", searched)

	// 以下の条件で検索
	//   - 名称に"test"と"example"を含む
	//   - ゾーンが"is1a"または"is1b"
	//   - 作成日時が1週間以上前
	condition = &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key("Name"):                               search.AndEqual("test", "example"),
			search.Key("Zone.Name"):                          search.OrEqual("is1a", "is1b"),
			search.KeyWithOp("CreatedAt", search.OpLessThan): time.Now().Add(-7 * 24 * time.Hour),
		},
	}
	searched, err = serverOp.Find(context.Background(), zone, condition)
	if err != nil {
		panic(err)
	}

	fmt.Printf("searched: %#v", searched)
}
