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

package iface

import (
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/util"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
)

type FindRequest struct {
	Zone string `request:"-" validate:"required"`

	MACAddresses      []string `request:"-" validate:"omitempty,dive,mac"`
	ServerIDs         []string `request:"-"`
	ServerNames       []string `request:"-"`
	PacketFilterIDs   []string `request:"-"`
	PacketFilterNames []string `request:"-"`

	Sort  search.SortKeys
	Count int
	From  int
}

func (req *FindRequest) Validate() error {
	return validate.Struct(req)
}

func (req *FindRequest) ToRequestParameter() (*sacloud.FindCondition, error) {
	condition := &sacloud.FindCondition{
		Filter: map[search.FilterKey]interface{}{},
	}
	if err := service.RequestConvertTo(req, condition); err != nil {
		return nil, err
	}

	if !util.IsEmpty(req.MACAddresses) {
		condition.Filter[search.Key("MACAddress")] = search.AndEqual(req.MACAddresses...)
	}
	if !util.IsEmpty(req.ServerIDs) {
		condition.Filter[search.Key("Server.ID")] = search.AndEqual(req.ServerIDs...)
	}
	if !util.IsEmpty(req.ServerNames) {
		condition.Filter[search.Key("Server.Name")] = search.AndEqual(req.ServerNames...)
	}
	if !util.IsEmpty(req.PacketFilterIDs) {
		condition.Filter[search.Key("PacketFilter.ID")] = search.AndEqual(req.PacketFilterIDs...)
	}
	if !util.IsEmpty(req.PacketFilterNames) {
		condition.Filter[search.Key("PacketFilter.Name")] = search.AndEqual(req.PacketFilterNames...)
	}
	return condition, nil
}
