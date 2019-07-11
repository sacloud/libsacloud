package test

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

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
	Skipf(format string, args ...interface{})
	Name() string
	Parallel()
}

// CRUDTestCase CRUD操作テストケース
type CRUDTestCase struct {
	// PreCheck テスト実行 or スキップを判定するためのFunc
	PreCheck func(TestT)

	// APICallerのセットアップ用Func、テストケースごとに1回呼ばれる
	SetupAPICallerFunc func() sacloud.APICaller

	// Setup テスト前の準備(依存リソースの作成など)を行うためのFunc(省略可)
	Setup func(*CRUDTestContext, sacloud.APICaller) error

	// Create Create操作のテスト用Func(省略可)
	Create *CRUDTestFunc

	// Read Read操作のテスト用Func(必須)
	Read *CRUDTestFunc

	// Updates Update操作のテスト用Func(省略可)
	Updates []*CRUDTestFunc

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

	// Expect 期待値、省略可能。省略してFunc内で自前実装することも可能。
	Expect *CRUDTestExpect

	// SkipExtractID Trueの場合Funcの戻り値からのID抽出(ioAddessor経由)を行わない
	SkipExtractID bool
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

// PreCheckEnvsFunc 指定の環境変数が指定されていなかった場合にテストをスキップするためのFuncを返す
func PreCheckEnvsFunc(envs ...string) func(TestT) {
	return func(t TestT) {
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
	if testCase.SetupAPICallerFunc == nil {
		t.Fatal("CRUDTestCase.SetupAPICallerFunc is required")
	}

	if testCase.Parallel {
		t.Parallel()
	}

	if testCase.PreCheck != nil {
		testCase.PreCheck(t)
	}

	testContext := &CRUDTestContext{
		Values: make(map[string]interface{}),
	}
	defer func() {
		// Cleanup
		if testCase.Cleanup != nil {
			if err := testCase.Cleanup(testContext, testCase.SetupAPICallerFunc()); err != nil {
				t.Logf("Cleanup is failed: ", err)
			}
		}
	}()

	if testCase.Setup != nil {
		if err := testCase.Setup(testContext, testCase.SetupAPICallerFunc()); err != nil {
			t.Fatal("Setup is failed: ", err)
		}
	}

	testFunc := func(f *CRUDTestFunc) error {
		actual, err := f.Func(testContext, testCase.SetupAPICallerFunc())
		if err != nil {
			return err
		}

		// extract ID from result of f.Func()
		if actual != nil && !f.SkipExtractID {
			if idHolder, ok := actual.(accessor.ID); ok {
				testContext.ID = idHolder.GetID()
			}
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
					return testCase.Read.Func(testContext, testCase.SetupAPICallerFunc())
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

	// Updates
	for _, updFunc := range testCase.Updates {
		if err := testFunc(updFunc); err != nil {
			t.Fatal("Update is failed: ", err)
		}
	}

	// Shutdown
	if testCase.Shutdown != nil {
		v, err := testCase.Read.Func(testContext, testCase.SetupAPICallerFunc())
		if err != nil {
			t.Fatal("Shutdown is failed: ", err)
		}
		if v, ok := v.(accessor.InstanceStatus); ok && v.GetInstanceStatus().IsUp() {
			if err := testCase.Shutdown(testContext, testCase.SetupAPICallerFunc()); err != nil {
				t.Fatal("Shutdown is failed: ", err)
			}

			waiter := sacloud.WaiterForDown(func() (interface{}, error) {
				return testCase.Read.Func(testContext, testCase.SetupAPICallerFunc())
			})
			if _, err := waiter.WaitForState(context.TODO()); err != nil {
				t.Fatal("WaitForDown is failed: ", err)
			}
		}
	}

	// Delete
	if testCase.Delete != nil {
		if err := testCase.Delete.Func(testContext, testCase.SetupAPICallerFunc()); err != nil {
			t.Fatal("Delete is failed: ", err)
		}
		// check not exists
		_, err := testCase.Read.Func(testContext, testCase.SetupAPICallerFunc())
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
