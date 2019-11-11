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

package builder

import (
	"log"
	"os"
	"testing"

	"github.com/sacloud/libsacloud/api"
)

var client *api.Client

func TestMain(m *testing.M) {
	//環境変数にトークン/シークレットがある場合のみテスト実施
	accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	if accessToken == "" || accessTokenSecret == "" {
		log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
		os.Exit(0) // exit normal
	}
	region := os.Getenv("SAKURACLOUD_REGION")
	if region == "" {
		region = "tk1v"
	}
	client = api.NewClient(accessToken, accessTokenSecret, region)

	ret := m.Run()
	os.Exit(ret)
}
