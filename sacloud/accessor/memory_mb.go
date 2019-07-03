package accessor

/************************************************
 MemoryMB - MemoryGB
************************************************/

// MemoryMB is accessor interface of MemoryMB field
type MemoryMB interface {
	GetMemoryMB() int
	SetMemoryMB(size int)
}

// GetMemoryGB returns GB
func GetMemoryGB(target MemoryMB) int {
	sizeMB := target.GetMemoryMB()
	if sizeMB == 0 {
		return 0
	}
	return sizeMB / 1024
}

// SetMemoryGB sets MemoryMB from GB
func SetMemoryGB(target MemoryMB, size int) {
	target.SetMemoryMB(size * 1024)
}
