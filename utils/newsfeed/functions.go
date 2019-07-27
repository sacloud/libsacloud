package newsfeed

import (
	"encoding/json"
	"io/ioutil"
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

	data, err := ioutil.ReadAll(resp.Body)
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
