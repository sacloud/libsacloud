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

// APIs APIでの操作対象リソースの定義
var APIs schema.Resources

func init() {
	APIs.Def(archiveAPI)      // アーカイブ
	APIs.Def(bridgeAPI)       // ブリッジ
	APIs.Def(cdromAPI)        // ISOイメージ(CD-ROM)
	APIs.Def(diskAPI)         // ディスク
	APIs.Def(gslbAPI)         // GSLB
	APIs.Def(interfaceAPI)    // インターフェース(NIC)
	APIs.Def(internetAPI)     // スイッチ+ルータ
	APIs.Def(loadBalancerAPI) // ロードバランサ
	APIs.Def(nfsAPI)          // NFS
	APIs.Def(noteAPI)         // スタートアップスクリプト
	APIs.Def(packetFilterAPI) // パケットフィルタ
	APIs.Def(serverAPI)       // サーバ
	APIs.Def(simAPI)          // SIM
	APIs.Def(switchAPI)       // スイッチ
	APIs.Def(vpcRouterAPI)    // VPCルータ
	APIs.Def(zoneAPI)         // ゾーン
}
