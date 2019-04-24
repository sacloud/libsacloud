package sacloud

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

var apiCaller APICaller

func TestMain(m *testing.M) {
	//環境変数にトークン/シークレットがある場合のみテスト実施
	accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	if accessToken == "" || accessTokenSecret == "" {
		log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
		os.Exit(0) // exit normal
	}
	region := os.Getenv("SAKURACLOUD_ZONE")
	if region == "" {
		region = "tk1v"
	}
	client := NewClient(accessToken, accessTokenSecret, region)
	client.DefaultTimeoutDuration = 30 * time.Minute
	client.UserAgent = fmt.Sprintf("test-libsacloud/%s", Version)
	client.AcceptLanguage = "en-US,en;q=0.9"

	apiCaller = client

	ret := m.Run()
	os.Exit(ret)
}
