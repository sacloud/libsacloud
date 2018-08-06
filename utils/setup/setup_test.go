package setup

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/sacloud/libsacloud"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

var client *api.Client
var testResourceName = "retryable-setup-test"

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
	client.DefaultTimeoutDuration = 30 * time.Minute
	client.UserAgent = fmt.Sprintf("test-libsacloud/%s", libsacloud.Version)
	client.AcceptLanguage = "en-US,en;q=0.9"

	ret := m.Run()
	os.Exit(ret)
}

func TestAccRetryAbleSetUp(t *testing.T) {

	defer initResources()

	swParam := client.Switch.New()
	swParam.Name = testResourceName
	sw, err := client.Switch.Create(swParam)
	assert.NoError(t, err)
	assert.NotNil(t, sw)

	param := sacloud.NewNFS(&sacloud.CreateNFSValue{
		SwitchID:  sw.GetStrID(),
		Plan:      sacloud.NFSPlan100G,
		IPAddress: "192.2.0.1",
		MaskLen:   24,
		Name:      testResourceName,
	})

	nfsBuilder := &RetryableSetup{
		Create: func() (sacloud.ResourceIDHolder, error) {
			return client.NFS.Create(param)
		},
		AsyncWaitForCopy: func(id int64) (chan interface{}, chan interface{}, chan error) {
			c, p, e := client.NFS.AsyncSleepWhileCopying(id, client.DefaultTimeoutDuration, 5)
			return c, p, e
		},
		Delete: func(id int64) error {
			_, err := client.NFS.Delete(id)
			return err
		},
		WaitForUp: func(id int64) error {
			return client.NFS.SleepUntilUp(id, client.DefaultTimeoutDuration)
		},
		RetryCount: 3,
	}

	res, err := nfsBuilder.Setup()
	assert.NoError(t, err)
	assert.NotNil(t, res)

	nfs, ok := res.(*sacloud.NFS)
	assert.True(t, ok)
	assert.NotNil(t, nfs)

}

func initResources() func() {
	cleanupResources()
	return cleanupResources
}

func cleanupResources() {
	{
		items, _ := client.NFS.Reset().WithNameLike(testResourceName).Find()
		wg := sync.WaitGroup{}
		wg.Add(len(items.NFS))

		for _, item := range items.NFS {
			nfs := item
			go func() {
				if nfs.IsUp() {
					client.NFS.Stop(nfs.ID)
					client.NFS.SleepUntilDown(nfs.ID, client.DefaultTimeoutDuration)
				}
				client.NFS.Delete(nfs.ID)
				wg.Done()
			}()
		}
		wg.Wait()
	}
	{
		items, _ := client.Switch.Reset().WithNameLike(testResourceName).Find()
		for _, item := range items.Switches {
			client.Switch.Delete(item.ID)
		}
	}
}
