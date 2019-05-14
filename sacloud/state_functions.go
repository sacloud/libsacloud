package sacloud

import "github.com/sacloud/libsacloud-v2/sacloud/types"

// WaiterForUp 起動完了まで待つためのStateWaiterを返す
func WaiterForUp(readFunc StateReadFunc) StateWaiter {
	return &StatePollWaiter{
		ReadFunc: readFunc,
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
			types.Availabilities.Migrating,
			types.Availabilities.Uploading,
			types.Availabilities.Transfering,
			types.Availabilities.Discontinued,
		},
		TargetInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Up,
		},
		PendingInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Unknown,
			types.ServerInstanceStatuses.Cleaning,
			types.ServerInstanceStatuses.Down,
		},
	}
}

// WaiterForApplianceUp 起動完了まで待つためのStateWaiterを返す
//
// アプライアンス向けに404発生時のリトライを設定可能
func WaiterForApplianceUp(readFunc StateReadFunc, notFoundRetry int) StateWaiter {
	return &StatePollWaiter{
		ReadFunc: readFunc,
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
			types.Availabilities.Migrating,
			types.Availabilities.Uploading,
			types.Availabilities.Transfering,
			types.Availabilities.Discontinued,
		},
		TargetInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Up,
		},
		PendingInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Unknown,
			types.ServerInstanceStatuses.Cleaning,
			types.ServerInstanceStatuses.Down,
		},
		NotFoundRetry: notFoundRetry,
	}
}

// WaiterForDown シャットダウン完了まで待つためのStateWaiterを返す
func WaiterForDown(readFunc StateReadFunc) StateWaiter {
	return &StatePollWaiter{
		ReadFunc: readFunc,
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
		},
		TargetInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Down,
		},
		PendingInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Up,
			types.ServerInstanceStatuses.Cleaning,
			types.ServerInstanceStatuses.Unknown,
		},
	}
}

// WaiterForReady リソースの利用準備完了まで待つためのStateWaiterを返す
func WaiterForReady(readFunc StateReadFunc) StateWaiter {
	return &StatePollWaiter{
		ReadFunc: readFunc,
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
			types.Availabilities.Migrating,
			types.Availabilities.Uploading,
			types.Availabilities.Transfering,
			types.Availabilities.Discontinued,
		},
	}
}
