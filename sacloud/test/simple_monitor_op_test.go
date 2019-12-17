// Copyright 2016-2019 The Libsacloud Authors
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

package test

import (
	"errors"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestSimpleMonitorOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: testSimpleMonitorCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createSimpleMonitorExpected,
				IgnoreFields: ignoreSimpleMonitorFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testSimpleMonitorRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createSimpleMonitorExpected,
				IgnoreFields: ignoreSimpleMonitorFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testSimpleMonitorUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateSimpleMonitorExpected,
					IgnoreFields: ignoreSimpleMonitorFields,
				}),
			},
			{
				Func: testSimpleMonitorUpdateSettings,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateSimpleMonitorSettingsExpected,
					IgnoreFields: ignoreSimpleMonitorFields,
				}),
			},
			{
				Func: testSimpleMonitorUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateSimpleMonitorToMinExpected,
					IgnoreFields: ignoreSimpleMonitorFields,
				}),
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testSimpleMonitorDelete,
		},
	})
}

var (
	ignoreSimpleMonitorFields = []string{
		"ID",
		"CreatedAt",
		"ModifiedAt",
		"Class",
		"SettingsHash",
	}
	createSimpleMonitorParam = &sacloud.SimpleMonitorCreateRequest{
		Target:      testutil.ResourceName("simple-monitor.usacloud.jp"),
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
		NotifyInterval:     2 * 60 * 60,
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
		NotifyInterval:     createSimpleMonitorParam.NotifyInterval,
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
		NotifyInterval:     3 * 60 * 60,
		IconID:             testIconID,
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
		NotifyInterval:     updateSimpleMonitorParam.NotifyInterval,
		SlackWebhooksURL:   updateSimpleMonitorParam.SlackWebhooksURL,
		Availability:       types.Availabilities.Available,
		IconID:             testIconID,
	}
	updateSimpleMonitorSettingsParam = &sacloud.SimpleMonitorUpdateSettingsRequest{
		DelayLoop: 120,
		HealthCheck: &sacloud.SimpleMonitorHealthCheck{
			Protocol:          types.SimpleMonitorProtocols.HTTP,
			Port:              types.StringNumber(80),
			Path:              "/index3.html",
			Status:            types.StringNumber(202),
			SNI:               types.StringFalse,
			Host:              "libsacloud-test.usacloud.jp",
			BasicAuthUsername: "username-upd2",
			BasicAuthPassword: "password-upd2",
		},
		NotifyEmailEnabled: types.StringFalse,
		NotifyEmailHTML:    types.StringFalse,
		NotifySlackEnabled: types.StringTrue,
		SlackWebhooksURL:   "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX",
		NotifyInterval:     3 * 60 * 60,
	}
	updateSimpleMonitorSettingsExpected = &sacloud.SimpleMonitor{
		Name:               createSimpleMonitorParam.Target,
		Description:        updateSimpleMonitorParam.Description,
		Tags:               updateSimpleMonitorParam.Tags,
		Target:             createSimpleMonitorParam.Target,
		DelayLoop:          updateSimpleMonitorSettingsParam.DelayLoop,
		Enabled:            updateSimpleMonitorSettingsParam.Enabled,
		HealthCheck:        updateSimpleMonitorSettingsParam.HealthCheck,
		NotifyEmailEnabled: updateSimpleMonitorSettingsParam.NotifyEmailEnabled,
		NotifyEmailHTML:    updateSimpleMonitorSettingsParam.NotifyEmailHTML,
		NotifySlackEnabled: updateSimpleMonitorSettingsParam.NotifySlackEnabled,
		NotifyInterval:     updateSimpleMonitorSettingsParam.NotifyInterval,
		SlackWebhooksURL:   updateSimpleMonitorSettingsParam.SlackWebhooksURL,
		Availability:       types.Availabilities.Available,
		IconID:             testIconID,
	}

	updateSimpleMonitorToMinParam = &sacloud.SimpleMonitorUpdateRequest{
		HealthCheck: &sacloud.SimpleMonitorHealthCheck{
			Protocol: types.SimpleMonitorProtocols.Ping,
			Host:     "libsacloud-test-upd.usacloud.jp",
		},
		NotifyEmailEnabled: types.StringTrue,
		Enabled:            true,
	}
	updateSimpleMonitorToMinExpected = &sacloud.SimpleMonitor{
		Name:               createSimpleMonitorParam.Target,
		Target:             createSimpleMonitorParam.Target,
		DelayLoop:          60, // default value
		Enabled:            updateSimpleMonitorToMinParam.Enabled,
		HealthCheck:        updateSimpleMonitorToMinParam.HealthCheck,
		NotifyEmailEnabled: updateSimpleMonitorToMinParam.NotifyEmailEnabled,
		Availability:       types.Availabilities.Available,
		NotifyInterval:     2 * 60 * 60, // default value
	}
)

func testSimpleMonitorCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Create(ctx, createSimpleMonitorParam)
}

func testSimpleMonitorRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testSimpleMonitorUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Update(ctx, ctx.ID, updateSimpleMonitorParam)
}

func testSimpleMonitorUpdateSettings(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.UpdateSettings(ctx, ctx.ID, updateSimpleMonitorSettingsParam)
}

func testSimpleMonitorUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Update(ctx, ctx.ID, updateSimpleMonitorToMinParam)
}

func testSimpleMonitorDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSimpleMonitorOp(caller)
	return client.Delete(ctx, ctx.ID)
}

func TestSimpleMonitorOp_StatusAndHealth(t *testing.T) {
	client := sacloud.NewSimpleMonitorOp(singletonAPICaller())
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				sm, err := client.Create(ctx, simpleMonitorStatusAndHealthTargetParam)
				if err != nil {
					return nil, err
				}
				if isAccTest() {
					time.Sleep(2 * time.Minute) // Statusの戻り値を確認するために数分待つ
				}
				return sm, nil
			},
		},

		Read: &testutil.CRUDTestFunc{
			Func: testSimpleMonitorRead,
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					return client.HealthStatus(ctx, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					healthStatus := v.(*sacloud.SimpleMonitorHealthStatus)
					if !assert.NotNil(t, healthStatus) {
						return errors.New("unexpected state: SimpleMonitorHealthStatus")
					}
					return nil
				},
			},
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					return client.MonitorResponseTime(ctx, ctx.ID, &sacloud.MonitorCondition{})
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					monitor := v.(*sacloud.ResponseTimeSecActivity)
					if !assert.NotNil(t, monitor) {
						return errors.New("unexpected state: ResponseTimeSecActivity")
					}
					return nil
				},
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testSimpleMonitorDelete,
		},
	})
}

var simpleMonitorStatusAndHealthTargetParam = &sacloud.SimpleMonitorCreateRequest{
	Target:    testutil.ResourceName("simple-monitor.usacloud.jp"),
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
