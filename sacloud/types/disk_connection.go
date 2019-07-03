package types

// EDiskConnection ディスク接続方法
type EDiskConnection string

// String EDiskConnectionの文字列表現
func (c EDiskConnection) String() string {
	return string(c)
}

// DiskConnections ディスク接続方法
var (
	DiskConnections = struct {
		// VirtIO virtio
		VirtIO EDiskConnection
		// IDE ide
		IDE EDiskConnection
	}{
		VirtIO: EDiskConnection("virtio"),
		IDE:    EDiskConnection("ide"),
	}

	DiskConnectionMap = map[string]EDiskConnection{
		DiskConnections.VirtIO.String(): DiskConnections.VirtIO,
		DiskConnections.IDE.String():    DiskConnections.IDE,
	}

	DiskConnectionValues = []string{
		DiskConnections.VirtIO.String(),
		DiskConnections.IDE.String(),
	}
)
