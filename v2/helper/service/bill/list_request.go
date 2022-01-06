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

package bill

import (
	"github.com/sacloud/libsacloud/v2/helper/validate"
)

type ListRequest struct {
	Year  int `validate:"required_with=Month"`
	Month int `validate:"min=0,max=12"`
}

func (req *ListRequest) Validate() error {
	return validate.Struct(req)
}
