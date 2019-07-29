package accessor

/************************************************
 SizeMB - SizeGB
************************************************/

// SizeMB is accessor interface of SizeMB field
type SizeMB interface {
	GetSizeMB() int
	SetSizeMB(size int)
}

// GetSizeGB returns GB
func GetSizeGB(target SizeMB) int {
	sizeMB := target.GetSizeMB()
	if sizeMB == 0 {
		return 0
	}
	return sizeMB / 1024
}

// SetSizeGB sets SizeMB from GB
func SetSizeGB(target SizeMB, size int) {
	target.SetSizeMB(size * 1024)
}
