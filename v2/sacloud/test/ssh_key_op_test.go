// Copyright 2016-2022 The Libsacloud Authors
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
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
)

func TestSSHKeyOpCRUD(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: testSSHKeyCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createSSHKeyExpected,
				IgnoreFields: ignoreSSHKeyFields,
			}),
		},
		Read: &testutil.CRUDTestFunc{
			Func: testSSHKeyRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createSSHKeyExpected,
				IgnoreFields: ignoreSSHKeyFields,
			}),
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testSSHKeyUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateSSHKeyExpected,
					IgnoreFields: ignoreSSHKeyFields,
				}),
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testSSHKeyDelete,
		},
	})
}

func TestSSHKeyOp_Generate(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := sacloud.NewSSHKeyOp(caller)
				return client.Generate(ctx, &sacloud.SSHKeyGenerateRequest{
					Name:        testutil.ResourceName("sshkey-generate"),
					Description: "libsacloud-sshKey-generate",
					PassPhrase:  "libsacloud-sshKey-passphrase",
				})
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
				sshKey := v.(*sacloud.SSHKeyGenerated)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, sshKey, "SSHKeyGenerated"),
					testutil.AssertNotEmptyFunc(t, sshKey.PublicKey, "SSHKeyGenerated.PublicKey"),
					testutil.AssertNotEmptyFunc(t, sshKey.PrivateKey, "SSHKeyGenerated.PrivateKey"),
					testutil.AssertNotEmptyFunc(t, sshKey.Fingerprint, "SSHKeyGenerated.Fingerprint"),
				)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: testSSHKeyRead,
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testSSHKeyDelete,
		},
	})
}

var (
	fakePublicKey   = `ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAs7YFtxjGrI49MCBnSFbUPxqz0e5HSGQPnLlPJ0u/9w4WLpoOZYmoQDTMfuFA61qv+0dp5mpMZPj3f5YEGlwUFKPy3Cmrp0ub1nYDb7n62s+Xf68TNvbVgQMLF0xdOaWxdRsQwmH8lOWan1Ubc8iwfOa3TNGwOzGLMjdW3PiJ7hcE7nFqnmbQUabHWow8G6JYDHKyjAdpz+edK8u+LY0iEP8M8VAjRJKJVg4p1/oDjHFKI0qjfjitKzoLm5FGaFv8afH2WQSpu/2To7d/RaLhfoMZsUReLSxeDnQkKGERXrAywTHnFu60cOaT3EvaAhP1H3BPj2LESm8M4ja9FaARnQ== `
	fakeFingerprint = "79:d7:ac:b8:cf:cf:01:44:b2:19:ba:d4:82:fd:c4:2d"

	ignoreSSHKeyFields = []string{
		"ID",
		"CreatedAt",
	}
	createSSHKeyParam = &sacloud.SSHKeyCreateRequest{
		Name:        testutil.ResourceName("sshkey"),
		Description: "libsacloud-sshKey",
		PublicKey:   fakePublicKey,
	}
	createSSHKeyExpected = &sacloud.SSHKey{
		Name:        createSSHKeyParam.Name,
		Description: createSSHKeyParam.Description,
		PublicKey:   fakePublicKey,
		Fingerprint: fakeFingerprint,
	}
	updateSSHKeyParam = &sacloud.SSHKeyUpdateRequest{
		Name:        testutil.ResourceName("sshkey-upd"),
		Description: "libsacloud-sshKey-upd",
	}
	updateSSHKeyExpected = &sacloud.SSHKey{
		Name:        updateSSHKeyParam.Name,
		Description: updateSSHKeyParam.Description,
		PublicKey:   fakePublicKey,
		Fingerprint: fakeFingerprint,
	}
)

func testSSHKeyCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSSHKeyOp(caller)
	return client.Create(ctx, createSSHKeyParam)
}

func testSSHKeyRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSSHKeyOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testSSHKeyUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSSHKeyOp(caller)
	return client.Update(ctx, ctx.ID, updateSSHKeyParam)
}

func testSSHKeyDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSSHKeyOp(caller)
	return client.Delete(ctx, ctx.ID)
}
