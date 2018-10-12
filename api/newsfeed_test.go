package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsFeedAPI(t *testing.T) {
	defer initNewsFeedAPI()()

	feedAPI := client.NewsFeed

	res, err := feedAPI.GetFeed()
	assert.NoError(t, err)
	assert.True(t, len(res) > 0)

	feed := res[0]
	assert.NotEmpty(t, feed.Description)
	assert.NotEmpty(t, feed.StrDate)
	assert.NotEmpty(t, feed.StrEventStart)
	assert.NotEmpty(t, feed.StrEventEnd)
	assert.NotEmpty(t, feed.Title)
	assert.NotEmpty(t, feed.URL)

	// by URL
	single, err := feedAPI.GetFeedByURL(feed.URL)
	assert.NoError(t, err)
	assert.NotNil(t, single)
	assert.EqualValues(t, &feed, single)
}

func initNewsFeedAPI() func() {
	cleanupNewsFeedAPI()
	return cleanupNewsFeedAPI
}

func cleanupNewsFeedAPI() {
}
