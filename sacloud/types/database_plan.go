package types

// DatabasePlans データベースプラン
var DatabasePlans = struct {
	// DB10GB 10GB
	DB10GB ID
	// DB30GB 30GB
	DB30GB ID
	// DB90GB 90GB
	DB90GB ID
	// DB240GB 240GB
	DB240GB ID
	// DB500GB 500GB
	DB500GB ID
	// DB1TB 1TB
	DB1TB ID
}{
	DB10GB:  ID(10),
	DB30GB:  ID(30),
	DB90GB:  ID(90),
	DB240GB: ID(240),
	DB500GB: ID(500),
	DB1TB:   ID(1000),
}

// SlaveDatabasePlanID マスター側のプランIDからスレーブのプランIDを算出
func SlaveDatabasePlanID(masterPlanID ID) ID {
	return ID(int64(masterPlanID) + 1)
}
