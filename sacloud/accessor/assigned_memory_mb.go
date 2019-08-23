package accessor

/************************************************
 AssignedMemoryMB - MemoryGB
************************************************/

// AssignedMemoryMB is accessor interface of MemoryMB field
type AssignedMemoryMB interface {
	GetAssignedMemoryMB() int
	SetAssignedMemoryMB(size int)
}

// GetAssignedMemoryGB returns GB
func GetAssignedMemoryGB(target AssignedMemoryMB) int {
	sizeMB := target.GetAssignedMemoryMB()
	if sizeMB == 0 {
		return 0
	}
	return sizeMB / 1024
}

// SetAssignedMemoryGB sets MemoryMB from GB
func SetAssignedMemoryGB(target AssignedMemoryMB, size int) {
	target.SetAssignedMemoryMB(size * 1024)
}
