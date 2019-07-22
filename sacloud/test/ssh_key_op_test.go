package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func TestSSHKeyOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testSSHKeyCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createSSHKeyExpected,
				IgnoreFields: ignoreSSHKeyFields,
			}),
		},
		Read: &CRUDTestFunc{
			Func: testSSHKeyRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createSSHKeyExpected,
				IgnoreFields: ignoreSSHKeyFields,
			}),
		},
		Updates: []*CRUDTestFunc{
			{
				Func: testSSHKeyUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateSSHKeyExpected,
					IgnoreFields: ignoreSSHKeyFields,
				}),
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testSSHKeyDelete,
		},
	})
}

func TestSSHKeyOp_Generate(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := sacloud.NewSSHKeyOp(caller)
				return client.Generate(ctx, &sacloud.SSHKeyGenerateRequest{
					Name:        "libsacloud-sshKey-generate",
					Description: "libsacloud-sshKey-generate",
					PassPhrase:  "libsacloud-sshKey-passphrase",
				})
			},
			CheckFunc: func(t TestT, ctx *CRUDTestContext, v interface{}) error {
				sshKey := v.(*sacloud.SSHKeyGenerated)
				return DoAsserts(
					AssertNotNilFunc(t, sshKey, "SSHKeyGenerated"),
					AssertNotEmptyFunc(t, sshKey.PublicKey, "SSHKeyGenerated.PublicKey"),
					AssertNotEmptyFunc(t, sshKey.PrivateKey, "SSHKeyGenerated.PrivateKey"),
					AssertNotEmptyFunc(t, sshKey.Fingerprint, "SSHKeyGenerated.Fingerprint"),
				)
			},
		},
		Read: &CRUDTestFunc{
			Func: testSSHKeyRead,
		},
		Delete: &CRUDTestDeleteFunc{
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
		Name:        "libsacloud-sshKey",
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
		Name:        "libsacloud-sshKey-upd",
		Description: "libsacloud-sshKey-upd",
	}
	updateSSHKeyExpected = &sacloud.SSHKey{
		Name:        updateSSHKeyParam.Name,
		Description: updateSSHKeyParam.Description,
		PublicKey:   fakePublicKey,
		Fingerprint: fakeFingerprint,
	}
)

func testSSHKeyCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSSHKeyOp(caller)
	return client.Create(ctx, createSSHKeyParam)
}

func testSSHKeyRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSSHKeyOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testSSHKeyUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSSHKeyOp(caller)
	return client.Update(ctx, ctx.ID, updateSSHKeyParam)
}

func testSSHKeyDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSSHKeyOp(caller)
	return client.Delete(ctx, ctx.ID)
}
