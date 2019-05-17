package accessor

/************************************************
 MigratedMB - MigratedGB
************************************************/

//MigratedMB is accessor interface of MigratedMB field
type MigratedMB interface {
	GetMigratedMB() int
	SetMigratedMB(size int)
}

// GetMigratedGB returns GB
func GetMigratedGB(target MigratedMB) int {
	sizeMB := target.GetMigratedMB()
	if sizeMB == 0 {
		return 0
	}
	return sizeMB / 1024
}

// SetMigratedGB sets MigratedMB from GB
func SetMigratedGB(target MigratedMB, size int) {
	target.SetMigratedMB(size * 1024)
}
