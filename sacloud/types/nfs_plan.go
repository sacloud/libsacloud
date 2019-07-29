package types

// NFSPlans NFSプラン
//
// Note: NFS作成時のPlanIDはこの値+サイズでNFSプランを検索、そのIDを指定すること
// NFSプランの検索はutils/nfsのfunc FindPlan(plan types.ID, size int64)を利用する
var NFSPlans = struct {
	// HDD hddプラン
	HDD ID
	// SSD ssdプラン
	SSD ID
}{
	HDD: ID(1),
	SSD: ID(2),
}
