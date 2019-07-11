package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestSimpleMonitorOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testSimpleMonitorCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createSimpleMonitorExpected,
				IgnoreFields: ignoreSimpleMonitorFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testSimpleMonitorRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createSimpleMonitorExpected,
				IgnoreFields: ignoreSimpleMonitorFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testSimpleMonitorUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateSimpleMonitorExpected,
					IgnoreFields: ignoreSimpleMonitorFields,
				}),
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testSimpleMonitorDelete,
		},
	})
}

var (
	ignoreSimpleMonitorFields = []string{
		"ID",
		"IconID",
		"CreatedAt",
		"ModifiedAt",
		"Class",
		"SettingsHash",
	}
	createSimpleMonitorParam = &sacloud.SimpleMonitorCreateRequest{
		Target:      "libsacloud-test.usacloud.jp",
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		DelayLoop:   60,
		Enabled:     types.StringTrue,
		HealthCheck: &sacloud.SimpleMonitorHealthCheck{
			Protocol:          types.SimpleMonitorProtocols.HTTP,
			Port:              types.StringNumber(80),
			Path:              "/index.html",
			Status:            types.StringNumber(200),
			SNI:               types.StringTrue,
			Host:              "libsacloud-test.usacloud.jp",
			BasicAuthUsername: "username",
			BasicAuthPassword: "password",
		},
		NotifyEmailEnabled: types.StringTrue,
		NotifyEmailHTML:    types.StringTrue,
		NotifySlackEnabled: types.StringFalse,
		SlackWebhooksURL:   "",
	}
	createSimpleMonitorExpected = &sacloud.SimpleMonitor{
		Name:               createSimpleMonitorParam.Target,
		Description:        createSimpleMonitorParam.Description,
		Tags:               createSimpleMonitorParam.Tags,
		Target:             createSimpleMonitorParam.Target,
		DelayLoop:          createSimpleMonitorParam.DelayLoop,
		Enabled:            createSimpleMonitorParam.Enabled,
		HealthCheck:        createSimpleMonitorParam.HealthCheck,
		NotifyEmailEnabled: createSimpleMonitorParam.NotifyEmailEnabled,
		NotifyEmailHTML:    createSimpleMonitorParam.NotifyEmailHTML,
		NotifySlackEnabled: createSimpleMonitorParam.NotifySlackEnabled,
		SlackWebhooksURL:   createSimpleMonitorParam.SlackWebhooksURL,
		Availability:       types.Availabilities.Available,
	}
	updateSimpleMonitorParam = &sacloud.SimpleMonitorUpdateRequest{
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		DelayLoop:   120,
		HealthCheck: &sacloud.SimpleMonitorHealthCheck{
			Protocol:          types.SimpleMonitorProtocols.HTTPS,
			Port:              types.StringNumber(443),
			Path:              "/index2.html",
			Status:            types.StringNumber(201),
			SNI:               types.StringFalse,
			Host:              "libsacloud-test-upd.usacloud.jp",
			BasicAuthUsername: "username-upd",
			BasicAuthPassword: "password-upd",
		},
		NotifyEmailEnabled: types.StringFalse,
		NotifyEmailHTML:    types.StringFalse,
		NotifySlackEnabled: types.StringTrue,
		SlackWebhooksURL:   "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX",
	}
	updateSimpleMonitorExpected = &sacloud.SimpleMonitor{
		Name:               createSimpleMonitorParam.Target,
		Description:        updateSimpleMonitorParam.Description,
		Tags:               updateSimpleMonitorParam.Tags,
		Target:             createSimpleMonitorParam.Target,
		DelayLoop:          updateSimpleMonitorParam.DelayLoop,
		Enabled:            updateSimpleMonitorParam.Enabled,
		HealthCheck:        updateSimpleMonitorParam.HealthCheck,
		NotifyEmailEnabled: updateSimpleMonitorParam.NotifyEmailEnabled,
		NotifyEmailHTML:    updateSimpleMonitorParam.NotifyEmailHTML,
		NotifySlackEnabled: updateSimpleMonitorParam.NotifySlackEnabled,
		SlackWebhooksURL:   updateSimpleMonitorParam.SlackWebhooksURL,
		Availability:       types.Availabilities.Available,
	}
)

func testSimpleMonitorCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Create(context.Background(), sacloud.APIDefaultZone, createSimpleMonitorParam)
}

func testSimpleMonitorRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Read(context.Background(), sacloud.APIDefaultZone, testContext.ID)
}

func testSimpleMonitorUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Update(context.Background(), sacloud.APIDefaultZone, testContext.ID, updateSimpleMonitorParam)
}

func testSimpleMonitorDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Delete(context.Background(), sacloud.APIDefaultZone, testContext.ID)
}

func TestSimpleMonitorOp_StatusAndHealth(t *testing.T) {
	client := sacloud.NewSimpleMonitorOp(singletonAPICaller())
	ctx := context.Background()

	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				sm, err := client.Create(ctx, sacloud.APIDefaultZone, simpleMonitorStatusAndHealthTargetParam)
				if err != nil {
					return nil, err
				}
				if isAccTest() {
					time.Sleep(2 * time.Minute) // Statusの戻り値を確認するために数分待つ
				}
				return sm, nil
			},
		},

		Read: &CRUDTestFunc{
			Func: testSimpleMonitorRead,
		},

		Updates: []*CRUDTestFunc{
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					return client.HealthStatus(ctx, sacloud.APIDefaultZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					healthStatus := v.(*sacloud.SimpleMonitorHealthStatus)
					if !assert.NotNil(t, healthStatus) {
						return errors.New("unexpected state: SimpleMonitorHealthStatus")
					}
					return nil
				},
			},
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					return client.MonitorResponseTime(ctx, sacloud.APIDefaultZone, testContext.ID, &sacloud.MonitorCondition{})
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					monitor := v.(*sacloud.ResponseTimeSecActivity)
					if !assert.NotNil(t, monitor) {
						return errors.New("unexpected state: ResponseTimeSecActivity")
					}
					return nil
				},
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testSimpleMonitorDelete,
		},
	})
}

var simpleMonitorStatusAndHealthTargetParam = &sacloud.SimpleMonitorCreateRequest{
	Target:    "cloud.sakura.ad.jp",
	DelayLoop: 60,
	Enabled:   true,
	HealthCheck: &sacloud.SimpleMonitorHealthCheck{
		Protocol: types.SimpleMonitorProtocols.HTTPS,
		Port:     443,
		Path:     "/",
		Status:   200,
	},
	NotifySlackEnabled: true,
	SlackWebhooksURL:   "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX",
}
