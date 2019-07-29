package types

// RDBMSType データベースアプライアンスでのRDBMS種別
type RDBMSType string

const (
	// RDBMSTypesMariaDB MariaDB
	RDBMSTypesMariaDB = RDBMSType("MariaDB")
	// RDBMSTypesPostgreSQL PostgreSQL
	RDBMSTypesPostgreSQL = RDBMSType("postgres")
)

// RDBMSVersion RDBMSごとの名称やリビジョンなどのバージョン指定時のパラメータ情報
type RDBMSVersion struct {
	Name     string
	Version  string
	Revision string
}

// RDBMSVersions RDBMSごとの名称やリビジョンなどのバージョン指定時のパラメータ情報
var RDBMSVersions = map[RDBMSType]*RDBMSVersion{
	RDBMSTypesMariaDB: {
		Name:     "MariaDB",
		Version:  "10.3",
		Revision: "10.3.15",
	},
	RDBMSTypesPostgreSQL: {
		Name:     "postgres",
		Version:  "11",
		Revision: "",
	},
}
