// Package define .
//go:generate go run ../tools/gen-api-models/main.go
//go:generate go run ../tools/gen-api-interfaces/main.go
//go:generate go run ../tools/gen-api-envelope/main.go
//go:generate go run ../tools/gen-api-result/main.go
//go:generate go run ../tools/gen-api-op/main.go
//go:generate go run ../tools/gen-api-tracer/main.go
//go:generate go run ../tools/gen-api-stub/main.go
//go:generate go run ../tools/gen-api-meta/main.go
//go:generate go run ../tools/gen-api-fake-store/main.go
//go:generate go run ../tools/gen-api-fake-op/main.go
package define

import "github.com/sacloud/libsacloud/v2/internal/schema"

// Resources APIでの操作対象リソースの定義
var Resources schema.Resources

func init() {
	Resources.Def(archiveAPI)      // アーカイブ
	Resources.Def(bridgeAPI)       // ブリッジ
	Resources.Def(cdromAPI)        // ISOイメージ(CD-ROM)
	Resources.Def(diskAPI)         // ディスク
	Resources.Def(gslbAPI)         // GSLB
	Resources.Def(interfaceAPI)    // インターフェース(NIC)
	Resources.Def(internetAPI)     // スイッチ+ルータ
	Resources.Def(loadBalancerAPI) // ロードバランサ
	Resources.Def(nfsAPI)          // NFS
	Resources.Def(noteAPI)         // スタートアップスクリプト
	Resources.Def(packetFilterAPI) // パケットフィルタ
	Resources.Def(serverAPI)       // サーバ
	Resources.Def(simAPI)          // SIM
	Resources.Def(switchAPI)       // スイッチ
	Resources.Def(vpcRouterAPI)    // VPCルータ
	Resources.Def(zoneAPI)         // ゾーン
}
