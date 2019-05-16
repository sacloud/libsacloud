package sacloud

import "github.com/sacloud/libsacloud-v2/sacloud/types"

/************************************************
 ID - StringID
************************************************/

type idAccessor interface {
	GetID() types.ID
	SetID(id types.ID)
}

func getStringID(target idAccessor) string {
	return target.GetID().String()
}

func setStringID(target idAccessor, id string) {
	target.SetID(types.StringID(id))
}

func getInt64ID(target idAccessor) int64 {
	return target.GetID().Int64()
}

func setInt64ID(target idAccessor, id int64) {
	target.SetID(types.ID(id))
}

/************************************************
 SizeMB - SizeGB
************************************************/

type sizeMBAccessor interface {
	GetSizeMB() int
	SetSizeMB(size int)
}

func getSizeGB(target sizeMBAccessor) int {
	sizeMB := target.GetSizeMB()
	if sizeMB == 0 {
		return 0
	}
	return sizeMB / 1024
}

func setSizeGB(target sizeMBAccessor, size int) {
	target.SetSizeMB(size * 1024)
}

//MigratedMB

/************************************************
 MigratedMB - MigratedGB
************************************************/

type migratedMBAccessor interface {
	GetMigratedMB() int
	SetMigratedMB(size int)
}

func getMigratedGB(target migratedMBAccessor) int {
	sizeMB := target.GetMigratedMB()
	if sizeMB == 0 {
		return 0
	}
	return sizeMB / 1024
}

func setMigratedGB(target migratedMBAccessor, size int) {
	target.SetMigratedMB(size * 1024)
}

/************************************************
 MemoryMB - MemoryGB
************************************************/

type memoryMBAccessor interface {
	GetMemoryMB() int
	SetMemoryMB(size int)
}

func getMemoryGB(target memoryMBAccessor) int {
	sizeMB := target.GetMemoryMB()
	if sizeMB == 0 {
		return 0
	}
	return sizeMB / 1024
}

func setMemoryGB(target memoryMBAccessor, size int) {
	target.SetMemoryMB(size * 1024)
}

/************************************************
 switchID
************************************************/

type switchIDAccessor interface {
	GetSwitchID() types.ID
	SetSwitchID(id types.ID)
}
