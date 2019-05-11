package sacloud

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func isAccTest() bool {
	return os.Getenv("TESTACC") == "1"
}

// TestT テストのライフサイクルを管理するためのインターフェース.
//
// 通常は*testing.Tを実装として利用する
type TestT interface {
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
	SetupAPICaller func() APICaller

	// PreCheck テスト前の前提条件チェック用Func, テストケースごとに1回呼ばれる
	PreCheck func(t *testing.T)

	// Create Create操作のテスト用Func(省略可)
	Create *CRUDTestFunc

	// Read Read操作のテスト用Func(必須)
	Read *CRUDTestFunc

	// Update Update操作のテスト用Func(省略可)
	Update *CRUDTestFunc

	// Delete Delete操作のテスト用Func(省略可)
	Delete *CRUDTestDeleteFunc

	// Cleanup APIで作成/変更したリソースなどのクリーンアップ用Func(省略化)
	Cleanup func(APICaller)

	// Parallel t.Parallelを呼ぶかのフラグ
	Parallel bool
}

// CRUDTestContext CRUD操作テストでのコンテキスト、一連のテスト中に共有される
type CRUDTestContext struct {
	ID types.ID
}

// CRUDTestIDHolder IDを保持するためのインターフェース
type CRUDTestIDHolder interface {
	GetID() types.ID
}

// CRUDTestFunc CRUD操作(DELETE以外)テストでのテスト用Func
type CRUDTestFunc struct {
	// Func API操作を行うFunc
	Func func(*CRUDTestContext, APICaller) (interface{}, error)
	// Expect 期待値
	Expect *CRUDTestExpect
}

// CRUDTestDeleteFunc CRUD操作テストのDeleteテスト用Func
type CRUDTestDeleteFunc struct {
	// Func API操作を行うFunc
	Func func(*CRUDTestContext, APICaller) error
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

// TestPreCheckEnvs 指定の環境変数が指定されていなかった場合にテストをスキップするためのFuncを返す
func TestPreCheckEnvs(envs ...string) func(*testing.T) {
	return func(t *testing.T) {
		for _, env := range envs {
			v := os.Getenv(env)
			if v == "" {
				t.Skipf("environment variable %q is not set. skip", env)
			}
		}
	}
}

// Test 任意の条件でCRUD操作をテストする
func Test(t TestT, testCase *CRUDTestCase) {
	if !isAccTest() {
		t.Skip("TESTACC is not set. skip")
	}

	if testCase.Read == nil {
		t.Fatal("CRUDTestCase.Read is required")
	}
	if testCase.SetupAPICaller == nil {
		t.Fatal("CRUDTestCase.SetupAPICaller is required")
	}

	if testCase.Parallel {
		t.Parallel()
	}

	testContext := &CRUDTestContext{}
	testFunc := func(f *CRUDTestFunc) error {
		actual, err := f.Func(testContext, testCase.SetupAPICaller())
		if err != nil {
			return err
		}
		if idHolder, ok := actual.(CRUDTestIDHolder); ok {
			testContext.ID = idHolder.GetID()
		}
		actual, expected := f.Expect.Prepare(actual)
		require.Equal(t, expected, actual)
		return nil
	}

	// Create
	if testCase.Create != nil {
		if err := testFunc(testCase.Create); err != nil {
			t.Fatal("Create is failed: ", err)
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

	// Delete
	if testCase.Delete != nil {
		if err := testCase.Delete.Func(testContext, testCase.SetupAPICaller()); err != nil {
			t.Fatal("Delete is failed: ", err)
		}
		// check not exists
		err := testFunc(testCase.Read)
		if err == nil {
			t.Fatal("Resource still exists: ", testContext.ID)
		}
		if e, ok := err.(APIError); ok {
			if e.ResponseCode() != http.StatusNotFound {
				t.Fatal("Reading after delete is failed: ", e)
			}
		}
	}

	// Cleanup
	if testCase.Cleanup != nil {
		testCase.Cleanup(testCase.SetupAPICaller())
	}
}
