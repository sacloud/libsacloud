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
	"strconv"
	"time"
)

// FeedItem メンテナンス/障害情報お知らせ
type FeedItem struct {
	StrDate       string `json:"date,omitempty"`
	Description   string `json:"desc,omitempty"`
	StrEventStart string `json:"event_start,omitempty"`
	StrEventEnd   string `json:"event_end,omitempty"`
	Title         string `json:"title,omitempty"`
	URL           string `json:"url,omitempty"`
}

// Date 対象日時
func (f *FeedItem) Date() time.Time {
	return f.parseTime(f.StrDate)
}

// EventStart 掲載開始日時
func (f *FeedItem) EventStart() time.Time {
	return f.parseTime(f.StrEventStart)
}

// EventEnd 掲載終了日時
func (f *FeedItem) EventEnd() time.Time {
	return f.parseTime(f.StrEventEnd)
}

func (f *FeedItem) parseTime(sec string) time.Time {
	s, _ := strconv.ParseInt(sec, 10, 64)
	return time.Unix(s, 0)
}
