package test

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/accessor"
	"github.com/sacloud/libsacloud/sacloud/types"
	"github.com/stretchr/testify/require"
)

func isAccTest() bool {
	return os.Getenv("TESTACC") == "1"
}

// TestT テストのライフサイクルを管理するためのインターフェース.
//
// 通常は*testing.Tを実装として利用する
type TestT interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	FailNow()
	Fatal(args ...interface{})
	Skip(args ...interface{})
	Name() string
	Parallel()
}

// CRUDTestCase CRUD操作テストケース
type CRUDTestCase struct {
	// APICallerのセットアップ用Func、テストケースごとに1回呼ばれる
	SetupAPICaller func() sacloud.APICaller

	// Setup テスト前の準備(依存リソースの作成など)を行うためのFunc(省略可)
	Setup func(*CRUDTestContext, sacloud.APICaller) error

	// Create Create操作のテスト用Func(省略可)
	Create *CRUDTestFunc

	// Read Read操作のテスト用Func(必須)
	Read *CRUDTestFunc

	// Update Update操作のテスト用Func(省略可)
	Update *CRUDTestFunc

	// Shutdown Delete操作の前のシャットダウン(省略可)
	Shutdown func(*CRUDTestContext, sacloud.APICaller) error

	// Delete Delete操作のテスト用Func(省略可)
	Delete *CRUDTestDeleteFunc

	// Cleanup APIで作成/変更したリソースなどのクリーンアップ用Func(省略化)
	Cleanup func(*CRUDTestContext, sacloud.APICaller) error

	// Parallel t.Parallelを呼ぶかのフラグ
	Parallel bool

	// IgnoreStartupWait リソース作成後の起動待ちを行わない
	IgnoreStartupWait bool
}

// CRUDTestContext CRUD操作テストでのコンテキスト、一連のテスト中に共有される
type CRUDTestContext struct {
	// ID CRUDテスト対象リソースのID
	//
	// Create/Read/Updateの戻り値がidAccessorの場合に各操作の後で設定される
	ID types.ID

	// Values 一連のテスト中に共有したい値
	//
	// 依存リソースのIDの保持などで利用する
	Values map[string]interface{}
}

// CRUDTestFunc CRUD操作(DELETE以外)テストでのテスト用Func
type CRUDTestFunc struct {
	// Func API操作を行うFunc
	Func func(*CRUDTestContext, sacloud.APICaller) (interface{}, error)
	// Expect 期待値
	Expect *CRUDTestExpect
}

// CRUDTestDeleteFunc CRUD操作テストのDeleteテスト用Func
type CRUDTestDeleteFunc struct {
	// Func API操作を行うFunc
	Func func(*CRUDTestContext, sacloud.APICaller) error
}

// CRUDTestExpect CRUD操作(DELETE以外)テストでの期待値
type CRUDTestExpect struct {
	// ExpectValue CRUD操作実行後の期待値
	ExpectValue interface{}

	// IgnoreFields比較時に無視する項目
	IgnoreFields []string
}

// Prepare テスト対象値を受け取り、比較可能な状態に加工した対象値と期待値を返す
func (c *CRUDTestExpect) Prepare(actual interface{}) (interface{}, interface{}) {
	toMap := func(v interface{}) map[string]interface{} {
		data, err := json.Marshal(v)
		if err != nil {
			log.Fatalf("prepare is failed: json.Marshal returned error: %s", err)
		}
		var m map[string]interface{}
		if err := json.Unmarshal(data, &m); err != nil {
			log.Fatalf("prepare is failed: json.Unmarshal returned error: %s", err)
		}
		for _, key := range c.IgnoreFields {
			if _, ok := m[key]; ok {
				delete(m, key)
			}
		}
		return m
	}

	return toMap(actual), toMap(c.ExpectValue)
}

// PreCheckEnvs 指定の環境変数が指定されていなかった場合にテストをスキップするためのFuncを返す
func PreCheckEnvs(envs ...string) func(*testing.T) {
	return func(t *testing.T) {
		for _, env := range envs {
			v := os.Getenv(env)
			if v == "" {
				t.Skipf("environment variable %q is not set. skip", env)
			}
		}
	}
}

// Run 任意の条件でCRUD操作をテストする
func Run(t TestT, testCase *CRUDTestCase) {
	if testCase.Read == nil {
		t.Fatal("CRUDTestCase.Read is required")
	}
	if testCase.SetupAPICaller == nil {
		t.Fatal("CRUDTestCase.SetupAPICaller is required")
	}

	if testCase.Parallel {
		t.Parallel()
	}

	testContext := &CRUDTestContext{
		Values: make(map[string]interface{}),
	}
	defer func() {
		// Cleanup
		if testCase.Cleanup != nil {
			if err := testCase.Cleanup(testContext, testCase.SetupAPICaller()); err != nil {
				t.Logf("Cleanup is failed: ", err)
			}
		}
	}()

	if testCase.Setup != nil {
		if err := testCase.Setup(testContext, testCase.SetupAPICaller()); err != nil {
			t.Fatal("Setup is failed: ", err)
		}
	}

	testFunc := func(f *CRUDTestFunc) error {
		actual, err := f.Func(testContext, testCase.SetupAPICaller())
		if err != nil {
			return err
		}
		if idHolder, ok := actual.(accessor.ID); ok {
			testContext.ID = idHolder.GetID()
		}
		if f.Expect != nil {
			actual, expect := f.Expect.Prepare(actual)
			require.Equal(t, expect, actual)
		}
		return nil
	}

	// Create
	if testCase.Create != nil {
		if err := testFunc(testCase.Create); err != nil {
			t.Fatal("Create is failed: ", err)
		}

		if testCase.Create.Expect != nil && !testCase.IgnoreStartupWait {

			_, ok1 := testCase.Create.Expect.ExpectValue.(accessor.Availability)
			_, ok2 := testCase.Create.Expect.ExpectValue.(accessor.InstanceStatus)
			if ok1 || ok2 {
				waiter := sacloud.WaiterForApplianceUp(func() (interface{}, error) {
					return testCase.Read.Func(testContext, testCase.SetupAPICaller())
				}, 30)
				if _, err := waiter.WaitForState(context.TODO()); err != nil {
					t.Fatal("WaitForUp is failed: ", err)
				}
			}

		}
	}

	// Read
	if err := testFunc(testCase.Read); err != nil {
		t.Fatal("Read is failed: ", err)
	}

	// Update
	if testCase.Update != nil {
		if err := testFunc(testCase.Update); err != nil {
			t.Fatal("Update is failed: ", err)
		}
	}

	// Shutdown
	if testCase.Shutdown != nil {
		v, err := testCase.Read.Func(testContext, testCase.SetupAPICaller())
		if err != nil {
			t.Fatal("Shutdown is failed: ", err)
		}
		if v, ok := v.(accessor.InstanceStatus); ok && v.GetInstanceStatus().IsUp() {
			if err := testCase.Shutdown(testContext, testCase.SetupAPICaller()); err != nil {
				t.Fatal("Shutdown is failed: ", err)
			}

			waiter := sacloud.WaiterForDown(func() (interface{}, error) {
				return testCase.Read.Func(testContext, testCase.SetupAPICaller())
			})
			if _, err := waiter.WaitForState(context.TODO()); err != nil {
				t.Fatal("WaitForDown is failed: ", err)
			}
		}
	}

	// Delete
	if testCase.Delete != nil {
		if err := testCase.Delete.Func(testContext, testCase.SetupAPICaller()); err != nil {
			t.Fatal("Delete is failed: ", err)
		}
		// check not exists
		_, err := testCase.Read.Func(testContext, testCase.SetupAPICaller())
		if err == nil {
			t.Fatal("Resource still exists: ", testContext.ID)
		}
		if e, ok := err.(sacloud.APIError); ok {
			if e.ResponseCode() != http.StatusNotFound {
				t.Fatal("Reading after delete is failed: ", e)
			}
		}
	}

}
