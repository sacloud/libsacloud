package accessor

// DiskMigratable ディスクのマイグレーション(コピー処理)が行えるリソース
type DiskMigratable interface {
	SizeMB
	MigratedMB
	Availability
}
