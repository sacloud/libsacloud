package api

import (
	"log"
	"os"
	"testing"
	"time"
)

var (
	client               *Client
	testSetupHandlers    []func()
	testTearDownHandlers []func()
)

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
	client = NewClient(accessToken, accessTokenSecret, region)
	client.DefaultTimeoutDuration = 1 * time.Hour

	for _, f := range testSetupHandlers {
		f()
	}

	ret := m.Run()

	for _, f := range testTearDownHandlers {
		f()
	}

	os.Exit(ret)
}
