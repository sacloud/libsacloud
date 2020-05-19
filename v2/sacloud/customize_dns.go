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

package sacloud

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// AddRecords レコードを追加します。名前/タイプ/値が同じレコードが存在する場合は何もしません
func (o *DNS) AddRecords(rs ...*DNSRecord) {
	for _, r := range rs {
		if o.IsExistRecord(r) {
			continue
		}
		o.Records = append(o.Records, r)
	}
}

// RemoveRecords 名前/タイプ/値が同じレコードを削除します
func (o *DNS) RemoveRecords(rs ...*DNSRecord) {
	var res []*DNSRecord
	for _, cur := range o.Records {
		remove := false
		for _, r := range rs {
			if cur.Equal(r) {
				remove = true
				break
			}
		}
		if !remove {
			res = append(res, cur)
		}
	}
	o.Records = res
}

// FindRecord 名前/タイプ/値が同じレコードを返す
func (o *DNS) FindRecord(name string, tp types.EDNSRecordType, rdata string) *DNSRecord {
	for _, r := range o.Records {
		if r.Equal(&DNSRecord{Name: name, Type: tp, RData: rdata}) {
			return r
		}
	}
	return nil
}

// IsExistRecord 名前/タイプ/値が同じレコードが存在する場合にtrueを返す
func (o *DNS) IsExistRecord(record *DNSRecord) bool {
	if record == nil {
		return false
	}
	for _, r := range o.Records {
		if r.Equal(record) {
			return true
		}
	}
	return false
}
