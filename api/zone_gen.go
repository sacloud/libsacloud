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

package api

/************************************************
  generated by IDE. for [ZoneAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *ZoneAPI) Reset() *ZoneAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *ZoneAPI) Offset(offset int) *ZoneAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *ZoneAPI) Limit(limit int) *ZoneAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *ZoneAPI) Include(key string) *ZoneAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *ZoneAPI) Exclude(key string) *ZoneAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *ZoneAPI) FilterBy(key string, value interface{}) *ZoneAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *ZoneAPI) FilterMultiBy(key string, value interface{}) *ZoneAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *ZoneAPI) WithNameLike(name string) *ZoneAPI {
	return api.FilterBy("Name", name)
}

//// WithTag タグ条件
//func (api *ZoneAPI) WithTag(tag string) *ZoneAPI {
//	return api.FilterBy("Tags.Name", tag)
//}
//
//// WithTags タグ(複数)条件
//func (api *ZoneAPI) WithTags(tags []string) *ZoneAPI {
//	return api.FilterBy("Tags.Name", []interface{}{tags})
//}

// func (api *ZoneAPI) WithSizeGib(size int) *ZoneAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *ZoneAPI) WithSharedScope() *ZoneAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *ZoneAPI) WithUserScope() *ZoneAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *ZoneAPI) SortBy(key string, reverse bool) *ZoneAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *ZoneAPI) SortByName(reverse bool) *ZoneAPI {
	api.sortByName(reverse)
	return api
}

// func (api *ZoneAPI) SortBySize(reverse bool) *ZoneAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
   To support Setxxx interface for Find()
************************************************/

// SetEmpty 検索条件のリセット
func (api *ZoneAPI) SetEmpty() {
	api.reset()
}

// SetOffset オフセット
func (api *ZoneAPI) SetOffset(offset int) {
	api.offset(offset)
}

// SetLimit リミット
func (api *ZoneAPI) SetLimit(limit int) {
	api.limit(limit)
}

// SetInclude 取得する項目
func (api *ZoneAPI) SetInclude(key string) {
	api.include(key)
}

// SetExclude 除外する項目
func (api *ZoneAPI) SetExclude(key string) {
	api.exclude(key)
}

// SetFilterBy 指定キーでのフィルター
func (api *ZoneAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetFilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *ZoneAPI) SetFilterMultiBy(key string, value interface{}) {
	api.filterBy(key, value, true)
}

// SetNameLike 名称条件
func (api *ZoneAPI) SetNameLike(name string) {
	api.FilterBy("Name", name)
}

//// SetTag タグ条件
//func (api *ZoneAPI) SetTag(tag string) {
//	api.FilterBy("Tags.Name", tag)
//}
//
//// SetTags タグ(複数)条件
//func (api *ZoneAPI) SetTags(tags []string) {
//	api.FilterBy("Tags.Name", []interface{}{tags})
//}

// func (api *ZoneAPI) SetSizeGib(size int) {
// 	api.FilterBy("SizeMB", size*1024)
// }

// func (api *ZoneAPI) SetSharedScope() {
// 	api.FilterBy("Scope", "shared")
// }

// func (api *ZoneAPI) SetUserScope() {
// 	api.FilterBy("Scope", "user")
// }

// SetSortBy 指定キーでのソート
func (api *ZoneAPI) SetSortBy(key string, reverse bool) {
	api.sortBy(key, reverse)
}

// SetSortByName 名称でのソート
func (api *ZoneAPI) SetSortByName(reverse bool) {
	api.sortByName(reverse)
}

// func (api *ZoneAPI) SetSortBySize(reverse bool) {
// 	api.sortBy("SizeMB", reverse)
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// func (api *ZoneAPI) New() *sacloud.Zone {
// 	return &sacloud.Zone{}
// }

// func (api *ZoneAPI) Create(value *sacloud.Zone) (*sacloud.Zone, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.create(api.createRequest(value), res)
// 	})
// }

// Read 読み取り
func (api *ZoneAPI) Read(id int64) (*sacloud.Zone, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

// func (api *ZoneAPI) Update(id string, value *sacloud.Zone) (*sacloud.Zone, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.update(id, api.createRequest(value), res)
// 	})
// }

// func (api *ZoneAPI) Delete(id string) (*sacloud.Zone, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.delete(id, nil, res)
// 	})
// }

/************************************************
  Inner functions
************************************************/

func (api *ZoneAPI) setStateValue(setFunc func(*sacloud.Request)) *ZoneAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *ZoneAPI) request(f func(*sacloud.Response) error) (*sacloud.Zone, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Zone, nil
}

func (api *ZoneAPI) createRequest(value *sacloud.Zone) *sacloud.Request {
	req := &sacloud.Request{}
	req.Zone = value
	return req
}
