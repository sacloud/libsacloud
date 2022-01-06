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

package otel

import (
	"encoding/json"
	"fmt"
	"sync"
)

var initOnce sync.Once

// Initialize initialize tracer and add client factory hooks
func Initialize(opts ...Option) {
	cnf := newConfig(opts...)
	initOnce.Do(func() {
		initialize(cnf)
	})
}

func initialize(cnf *config) {
	addClientFactoryHooks(cnf)
}

func forceString(v interface{}) string {
	if v == nil {
		return "<nil>"
	}
	switch v := v.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	default:
		data, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		return string(data)
	}
}
