package testutil

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
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

	// LastValue 最後の操作での戻り値
	LastValue interface{}

	ctx  context.Context
	once sync.Once
}

func (c *CRUDTestContext) initInnerContext() {
	c.once.Do(func() {
		c.ctx = context.TODO()
	})
}

// Deadline context.Context実装
func (c *CRUDTestContext) Deadline() (deadline time.Time, ok bool) {
	c.initInnerContext()
	return c.ctx.Deadline()
}

// Done context.Context実装
func (c *CRUDTestContext) Done() <-chan struct{} {
	c.initInnerContext()
	return c.ctx.Done()
}

// Err context.Context実装
func (c *CRUDTestContext) Err() error {
	c.initInnerContext()
	return c.ctx.Err()
}

// Value context.Context実装
func (c *CRUDTestContext) Value(key interface{}) interface{} {
	c.initInnerContext()
	return c.ctx.Value(key)
}

// CRUDTestFunc CRUD操作(DELETE以外)テストでのテスト用Func
type CRUDTestFunc struct {
	// Func API操作を行うFunc
	Func func(*CRUDTestContext, sacloud.APICaller) (interface{}, error)

	// CheckFunc 任意のチェックを行うためのFunc、省略可能。
	CheckFunc func(TestT, *CRUDTestContext, interface{}) error

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
			t.Error("Setup is failed: ", err)
			return
		}
	}

	testFunc := func(f *CRUDTestFunc) error {
		actual, err := f.Func(testContext, testCase.SetupAPICallerFunc())
		if err != nil {
			return err
		}
		testContext.LastValue = actual

		if actual != nil && f.CheckFunc != nil {
			if err := f.CheckFunc(t, testContext, actual); err != nil {
				return err
			}
		}

		// extract ID from result of f.Func()
		if actual != nil && !f.SkipExtractID {
			if idHolder, ok := actual.(accessor.ID); ok {
				testContext.ID = idHolder.GetID()
			}
		}

		return nil
	}

	// Create
	if testCase.Create != nil {
		if err := testFunc(testCase.Create); err != nil {
			t.Error("Create is failed: ", err)
			return
		}

		if !testCase.IgnoreStartupWait && testCase.Read != nil && testContext.LastValue != nil {
			waiter := sacloud.WaiterForApplianceUp(func() (interface{}, error) {
				return testCase.Read.Func(testContext, testCase.SetupAPICallerFunc())
			}, 100)
			if _, err := waiter.WaitForState(context.TODO()); err != nil {
				t.Error("WaitForUp is failed: ", err)
				return
			}
		}
	}

	// Read
	if testCase.Read != nil {
		if err := testFunc(testCase.Read); err != nil {
			t.Fatal("Read is failed: ", err)
		}
	}

	// Updates
	for _, updFunc := range testCase.Updates {
		if err := testFunc(updFunc); err != nil {
			t.Error("Update is failed: ", err)
			return
		}
	}

	// Shutdown
	if testCase.Shutdown != nil {
		if testCase.Read == nil {
			t.Log("CRUDTestCase.Shutdown is set, but CRUDTestCase.Read is nil. Shutdown is skipped")
		} else {
			v, err := testCase.Read.Func(testContext, testCase.SetupAPICallerFunc())
			if err != nil {
				t.Error("Shutdown is failed: ", err)
				return
			}
			if v, ok := v.(accessor.InstanceStatus); ok && v.GetInstanceStatus().IsUp() {
				if err := testCase.Shutdown(testContext, testCase.SetupAPICallerFunc()); err != nil {
					t.Error("Shutdown is failed: ", err)
					return
				}

				waiter := sacloud.WaiterForDown(func() (interface{}, error) {
					return testCase.Read.Func(testContext, testCase.SetupAPICallerFunc())
				})
				if _, err := waiter.WaitForState(context.TODO()); err != nil {
					t.Error("WaitForDown is failed: ", err)
					return
				}
			}
		}
	}

	// Delete
	if testCase.Delete != nil {
		if err := testCase.Delete.Func(testContext, testCase.SetupAPICallerFunc()); err != nil {
			t.Error("Delete is failed: ", err)
			return
		}
		if testCase.Read != nil {
			// check not exists
			_, err := testCase.Read.Func(testContext, testCase.SetupAPICallerFunc())
			if err == nil {
				t.Error("Resource still exists: ", testContext.ID)
				return
			}
			if e, ok := err.(sacloud.APIError); ok {
				if e.ResponseCode() != http.StatusNotFound {
					t.Error("Reading after delete is failed: ", e)
					return
				}
			}
		}
	}
}

// AssertEqualWithExpected 項目ごとに除外設定のできる期待値との比較
func AssertEqualWithExpected(testExpect *CRUDTestExpect) func(TestT, *CRUDTestContext, interface{}) error {
	return func(t TestT, testContext *CRUDTestContext, v interface{}) error {
		actual, expect := testExpect.Prepare(v)
		if !assert.Equal(t, expect, actual) {
			return errors.New("assert.Equal is failed")
		}
		return nil
	}
}

// AssertEqual 値の比較
func AssertEqual(t TestT, expected interface{}, actual interface{}, targetName string) error {
	if !assert.Equal(t, expected, actual) {
		return fmt.Errorf("assert.Equal is failed: %s", targetName)
	}
	return nil
}

// AssertLen lengthのチェック
func AssertLen(t TestT, object interface{}, length int, targetName string) error {
	if !assert.Len(t, object, length) {
		return fmt.Errorf("assert.Len is failed: %s", targetName)
	}
	return nil
}

// AssertNil nilチェック
func AssertNil(t TestT, object interface{}, targetName string) error {
	if !assert.Nil(t, object) {
		return fmt.Errorf("assert.Nil is failed: %s", targetName)
	}
	return nil
}

// AssertNotNil not nilチェック
func AssertNotNil(t TestT, object interface{}, targetName string) error {
	if !assert.NotNil(t, object) {
		return fmt.Errorf("assert.NotNil is failed: %s", targetName)
	}
	return nil
}

// AssertTrue trueチェック
func AssertTrue(t TestT, value bool, targetName string) error {
	if !assert.True(t, value) {
		return fmt.Errorf("assert.True is failed: %s", targetName)
	}
	return nil
}

// AssertFalse falseチェック
func AssertFalse(t TestT, value bool, targetName string) error {
	if !assert.False(t, value) {
		return fmt.Errorf("assert.False is failed: %s", targetName)
	}
	return nil
}

// AssertEmpty emptyチェック
func AssertEmpty(t TestT, object interface{}, targetName string) error {
	if !assert.Empty(t, object) {
		return fmt.Errorf("assert.Empty is failed: %s", targetName)
	}
	return nil
}

// AssertNotEmpty not emptyチェック
func AssertNotEmpty(t TestT, object interface{}, targetName string) error {
	if !assert.NotEmpty(t, object) {
		return fmt.Errorf("assert.NotEmpty is failed: %s", targetName)
	}
	return nil
}

// DoAsserts アサーションを複数適用、最初のエラーを返す
func DoAsserts(funcs ...func() error) error {
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

// AssertEqualFunc 値の比較
func AssertEqualFunc(t TestT, expected interface{}, actual interface{}, targetName string) func() error {
	return func() error {
		return AssertEqual(t, expected, actual, targetName)
	}
}

// AssertLenFunc lengthのチェック
func AssertLenFunc(t TestT, object interface{}, length int, targetName string) func() error {
	return func() error {
		return AssertLen(t, object, length, targetName)
	}
}

// AssertNilFunc nilチェック
func AssertNilFunc(t TestT, object interface{}, targetName string) func() error {
	return func() error {
		return AssertNil(t, object, targetName)
	}
}

// AssertNotNilFunc not nilチェック
func AssertNotNilFunc(t TestT, object interface{}, targetName string) func() error {
	return func() error {
		return AssertNotNil(t, object, targetName)
	}
}

// AssertTrueFunc trueチェック
func AssertTrueFunc(t TestT, value bool, targetName string) func() error {
	return func() error {
		return AssertTrue(t, value, targetName)
	}
}

// AssertFalseFunc falseチェック
func AssertFalseFunc(t TestT, value bool, targetName string) func() error {
	return func() error {
		return AssertFalse(t, value, targetName)
	}
}

// AssertEmptyFunc emptyチェック
func AssertEmptyFunc(t TestT, object interface{}, targetName string) func() error {
	return func() error {
		return AssertEmpty(t, object, targetName)
	}
}

// AssertNotEmptyFunc not emptyチェック
func AssertNotEmptyFunc(t TestT, object interface{}, targetName string) func() error {
	return func() error {
		return AssertNotEmpty(t, object, targetName)
	}
}
