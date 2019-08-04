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
