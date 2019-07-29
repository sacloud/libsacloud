package types

// ENFSSize NFSサイズ
type ENFSSize int

// NFSHDDSizes NFSのHDDプランで指定可能なサイズ
var NFSHDDSizes = struct {
	Size100GB ENFSSize
	Size500GB ENFSSize
	Size1TB   ENFSSize
	Size2TB   ENFSSize
	Size4TB   ENFSSize
	Size8TB   ENFSSize
	Size12TB  ENFSSize
}{

	Size100GB: ENFSSize(100),
	Size500GB: ENFSSize(500),
	Size1TB:   ENFSSize(1024 * 1),
	Size2TB:   ENFSSize(1024 * 2),
	Size4TB:   ENFSSize(1024 * 4),
	Size8TB:   ENFSSize(1024 * 8),
	Size12TB:  ENFSSize(1024 * 12),
}

// NFSSSDSizes NFSのSSDプランで指定可能なサイズ
var NFSSSDSizes = struct {
	Size100GB ENFSSize
	Size500GB ENFSSize
	Size1TB   ENFSSize
	Size2TB   ENFSSize
	Size4TB   ENFSSize
}{

	Size100GB: ENFSSize(100),
	Size500GB: ENFSSize(500),
	Size1TB:   ENFSSize(1024 * 1),
	Size2TB:   ENFSSize(1024 * 2),
	Size4TB:   ENFSSize(1024 * 4),
}
