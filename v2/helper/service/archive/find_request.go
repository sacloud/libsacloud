// Copyright 2016-2021 The Libsacloud Authors
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

package archive

import (
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/util"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type FindRequest struct {
	Zone string `request:"-" validate:"required"`

	// OSType OS種別、NamesやTagsを指定した場合はそちらが優先される
	OSType ostype.ArchiveOSType `request:"-"`

	Names []string     `request:"-"`
	Tags  []string     `request:"-"`
	Scope types.EScope `request:"-"`

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

	filter, ok := ostype.ArchiveCriteria[req.OSType]
	if ok {
		for k, v := range filter {
			condition.Filter[k] = v
		}
	}
	if !util.IsEmpty(req.Names) {
		condition.Filter[search.Key("Name")] = search.AndEqual(req.Names...)
	}
	if !util.IsEmpty(req.Tags) {
		condition.Filter[search.Key("Tags.Name")] = search.TagsAndEqual(req.Tags...)
	}
	if !util.IsEmpty(req.Scope) {
		condition.Filter[search.Key("Scope")] = search.OrEqual(req.Scope)
	}
	return condition, nil
}
