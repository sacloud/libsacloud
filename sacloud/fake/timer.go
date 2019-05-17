package fake

import (
	"fmt"
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/accessor"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

var (
	// DiskCopyDuration ディスクコピー処理のtickerで利用するduration
	DiskCopyDuration = 10 * time.Millisecond
	// PowerOnDuration 電源On処理のtickerで利用するduration
	PowerOnDuration = 10 * time.Millisecond
	// PowerOffDuration 電源Off処理のtickerで利用するduration
	PowerOffDuration = 10 * time.Millisecond
)

func startDiskCopy(resourceKey, zone string, id types.ID) {
	counter := 0
	ticker := time.NewTicker(DiskCopyDuration)
	go func() {
		for {
			<-ticker.C

			raw := s.getByID(resourceKey, zone, id)
			if raw == nil {
				return
			}
			target, ok := raw.(accessor.DiskMigratable)
			if !ok {
				return
			}

			if counter < 3 {
				target.SetAvailability(types.Availabilities.Migrating)
				if counter == 0 {
					target.SetMigratedMB(0)
				} else {
					target.SetMigratedMB(int(target.GetSizeMB() / counter))
				}
			} else {
				target.SetAvailability(types.Availabilities.Available)
				target.SetMigratedMB(target.GetSizeMB())
			}
			s.set(resourceKey, zone, target)
			counter++
		}
	}()
}

func startPowerOn(resourceKey, zone string, id types.ID) {
	counter := 0
	ticker := time.NewTicker(PowerOnDuration)
	go func() {
		for {
			<-ticker.C

			raw := s.getByID(resourceKey, zone, id)
			if raw == nil {
				return
			}
			target, ok := raw.(accessor.InstanceStatus)
			if !ok {
				return
			}

			if counter < 3 {
				target.SetInstanceStatus(types.ServerInstanceStatuses.Down)
			} else {
				target.SetInstanceStatus(types.ServerInstanceStatuses.Up)
				if status, ok := target.(accessor.Instance); ok {
					now := time.Now()
					status.SetInstanceHostName(fmt.Sprintf("sac-%s-svXXX", zone))
					status.SetInstanceHostInfoURL("")
					status.SetInstanceStatusChangedAt(&now)
				}
				if available, ok := target.(accessor.Availability); ok {
					available.SetAvailability(types.Availabilities.Available)
				}
			}
			s.set(resourceKey, zone, target)
			counter++
		}
	}()
}

func startPowerOff(resourceKey, zone string, id types.ID) {
	counter := 0
	ticker := time.NewTicker(PowerOnDuration)
	go func() {
		for {
			<-ticker.C

			raw := s.getByID(resourceKey, zone, id)
			if raw == nil {
				return
			}
			target, ok := raw.(accessor.InstanceStatus)
			if !ok {
				return
			}

			if counter < 3 {
				target.SetInstanceStatus(types.ServerInstanceStatuses.Cleaning)
			} else {
				target.SetInstanceStatus(types.ServerInstanceStatuses.Down)
			}
			if status, ok := target.(accessor.Instance); ok {
				now := time.Now()
				status.SetInstanceHostName(fmt.Sprintf("sac-%s-svXXX", zone))
				status.SetInstanceHostInfoURL("")
				status.SetInstanceStatusChangedAt(&now)
			}
			s.set(resourceKey, zone, target)
			counter++
		}
	}()
}
