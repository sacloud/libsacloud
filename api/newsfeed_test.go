// Copyright 2016-2020 The Libsacloud Authors
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
