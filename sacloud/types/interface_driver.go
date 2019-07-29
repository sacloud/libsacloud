package types

// EInterfaceDriver インターフェースドライバ
type EInterfaceDriver string

func (d EInterfaceDriver) String() string {
	return string(d)
}

var (
	// InterfaceDrivers インターフェースドライバ
	InterfaceDrivers = struct {
		VirtIO EInterfaceDriver // virtio
		E1000  EInterfaceDriver // e1000
	}{
		VirtIO: EInterfaceDriver("virtio"),
		E1000:  EInterfaceDriver("e1000"),
	}

	// InterfaceDriverMap インターフェースドライバと文字列表現のマップ
	InterfaceDriverMap = map[string]EInterfaceDriver{
		InterfaceDrivers.VirtIO.String(): InterfaceDrivers.VirtIO,
		InterfaceDrivers.E1000.String():  InterfaceDrivers.E1000,
	}

	// InterfaceDriverValues インターフェースドライバが取りうる有効値
	InterfaceDriverValues = []string{
		InterfaceDrivers.VirtIO.String(),
		InterfaceDrivers.E1000.String(),
	}
)
