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
