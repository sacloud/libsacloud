package sacloud

import (
	"fmt"
	"strconv"
)

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

type idAccessor interface {
	GetID() int64
	SetID(id int64)
}

func getStringID(target idAccessor) string {
	return fmt.Sprintf("%d", target.GetID())
}

func setStringID(target idAccessor, id string) {
	intID, _ := strconv.ParseInt(id, 10, 64) // nolint ignore error
	target.SetID(intID)
}
