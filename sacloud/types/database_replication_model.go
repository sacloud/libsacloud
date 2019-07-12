package types

// EDatabaseReplicationModel データベースのレプリケーションモデル
type EDatabaseReplicationModel string

// DatabaseReplicationModels データベースのレプリケーションモデル
var DatabaseReplicationModels = struct {
	// MasterSlave マスター側
	MasterSlave EDatabaseReplicationModel
	// AsyncReplica スレーブ側
	AsyncReplica EDatabaseReplicationModel
}{
	MasterSlave:  EDatabaseReplicationModel("Master-Slave"),
	AsyncReplica: EDatabaseReplicationModel("Async-Replica"),
}
