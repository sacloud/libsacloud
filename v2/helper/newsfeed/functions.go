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

package newsfeed

import (
	"encoding/json"
	"io"
	"net/http"
)

// NewsFeedURL フィード取得URL
var NewsFeedURL = "https://secure.sakura.ad.jp/rss/sakuranews/getfeeds.php?service=cloud&format=json"

// Get ニュースフィード(障害/メンテナンス情報)を取得
func Get() ([]*FeedItem, error) {
	resp, err := http.Get(NewsFeedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var items []*FeedItem
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// GetByURL 指定のURLを持つフィードを取得
func GetByURL(url string) (*FeedItem, error) {
	items, err := Get()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.URL == url {
			return item, nil
		}
	}
	return nil, nil
}
