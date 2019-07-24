package types

// DiskPlans ディスクプランID 利用可能なサイズはDiskPlanAPIで取得すること
var DiskPlans = struct {
	// SSD ssdプラン
	SSD ID
	// HDD hddプラン
	HDD ID
}{
	SSD: ID(4),
	HDD: ID(2),
}
