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

package fake

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var (
	accountID    = types.ID(123456789012)
	accountName  = "fakeアカウント"
	accountCode  = "fake"
	accountClass = "member"
	memberCode   = "fake-member"
	memberClass  = "member"
)

var (
	sharedSegmentSwitch *sacloud.Switch

	zones      = types.ZoneNames
	zoneIDs    = types.ZoneIDs
	authStatus = &sacloud.AuthStatus{
		AccountID:          accountID,
		AccountName:        accountName,
		AccountCode:        accountCode,
		AccountClass:       accountClass,
		MemberCode:         memberCode,
		MemberClass:        memberClass,
		AuthClass:          types.AuthClasses.Account,
		AuthMethod:         types.AuthMethods.APIKey,
		IsAPIKey:           true,
		ExternalPermission: types.ExternalPermission("bill+eventlog+cdn"),
		OperationPenalty:   types.OperationPenalties.None,
		Permission:         types.Permissions.Create,
	}
)
