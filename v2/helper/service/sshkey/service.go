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

package sshkey

import "github.com/sacloud/libsacloud/v2/sacloud"

// Service provides a high-level API of for SSHKey
type Service struct {
	caller sacloud.APICaller
}

// New returns new service instance of SSHKey
func New(caller sacloud.APICaller) *Service {
	return &Service{caller: caller}
}
