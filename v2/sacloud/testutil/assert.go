// Copyright 2016-2020 The Libsacloud Authors
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

package testutil

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// AssertEqualWithExpected 項目ごとに除外設定のできる期待値との比較
func AssertEqualWithExpected(testExpect *CRUDTestExpect) func(TestT, *CRUDTestContext, interface{}) error {
	return func(t TestT, testContext *CRUDTestContext, v interface{}) error {
		expect, actual := testExpect.Prepare(v)
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
