package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var testNewsFeedJSON = `
[
    {
        "date": "1498396961",
        "desc": "[2017\u5e7406\u670825\u65e5]\u30d0\u30c3\u30af\u30dc\u30fc\u30f3\u30cd\u30c3\u30c8\u30ef\u30fc\u30af\u306e\u4e00\u90e8\u3067\u969c\u5bb3\u304c\u767a\u751f\u3057\u307e\u3057\u305f",
        "event_end": "1498394400",
        "event_start": "1498393200",
        "title": "[\u969c\u5bb3] \u30d0\u30c3\u30af\u30dc\u30fc\u30f3\u30cd\u30c3\u30c8\u30ef\u30fc\u30af\u306e\u4e00\u90e8",
        "url": "http://support.sakura.ad.jp/mainte/mainteentry.php?id=22178"
    },
    {
        "date": "1498261173",
        "desc": " [2017\u5e746\u670824\u65e5]\u3055\u304f\u3089\u306e\u30af\u30e9\u30a6\u30c9\u306e\u4e00\u90e8\u3067\u969c\u5bb3\u304c\u767a\u751f\u3057\u307e\u3057\u305f",
        "event_end": "1498262400",
        "event_start": "1498260900",
        "title": "[\u969c\u5bb3] \u3055\u304f\u3089\u306e\u30af\u30e9\u30a6\u30c9 \u4e00\u90e8\u306e\u304a\u5ba2\u69d8",
        "url": "http://support.sakura.ad.jp/mainte/mainteentry.php?id=22162"
    }
]
`

func TestMarshalNewsFeedJSON(t *testing.T) {
	var feeds = []NewsFeed{}
	err := json.Unmarshal([]byte(testNewsFeedJSON), &feeds)
	assert.NoError(t, err)
	assert.NotEmpty(t, feeds)
	assert.Equal(t, 2, len(feeds))

	var rawData = []map[string]interface{}{}
	err = json.Unmarshal([]byte(testNewsFeedJSON), &rawData)
	assert.NoError(t, err)

	feed := feeds[0]
	rawFeed := rawData[0]

	assert.EqualValues(t, rawFeed["date"], feed.StrDate)
	assert.EqualValues(t, rawFeed["desc"], feed.Description)
	assert.EqualValues(t, rawFeed["event_start"], feed.StrEventStart)
	assert.EqualValues(t, rawFeed["event_end"], feed.StrEventEnd)
	assert.EqualValues(t, rawFeed["title"], feed.Title)
	assert.EqualValues(t, rawFeed["url"], feed.Url)

	parsedDateSec, _ := strconv.ParseInt(feed.StrDate, 10, 64)
	assert.EqualValues(t, time.Unix(parsedDateSec, 0), feed.Date())

}
