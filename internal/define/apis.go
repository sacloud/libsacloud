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

import "github.com/sacloud/libsacloud/v2/internal/dsl"

// APIs APIでの操作対象リソースの定義
var APIs dsl.Resources

func init() {
	APIs.Define(archiveAPI)         // アーカイブ
	APIs.Define(authStatusAPI)      // 認証情報
	APIs.Define(autoBackupAPI)      // 自動バックアップ
	APIs.Define(billAPI)            // 請求情報
	APIs.Define(bridgeAPI)          // ブリッジ
	APIs.Define(cdromAPI)           // ISOイメージ(CD-ROM)
	APIs.Define(couponAPI)          // クーポン
	APIs.Define(databaseAPI)        // データベース
	APIs.Define(diskAPI)            // ディスク
	APIs.Define(diskPlanAPI)        // ディスクプラン
	APIs.Define(dnsAPI)             // DNS
	APIs.Define(gslbAPI)            // GSLB
	APIs.Define(iconAPI)            // アイコン
	APIs.Define(interfaceAPI)       // インターフェース(NIC)
	APIs.Define(internetAPI)        // スイッチ+ルータ
	APIs.Define(internetPlanAPI)    // ルータプラン
	APIs.Define(ipAPI)              // IPアドレス
	APIs.Define(ipv6netAPI)         // IPv6ネットワーク
	APIs.Define(ipv6AddrAPI)        // IPv6アドレス
	APIs.Define(licenseAPI)         // ライセンス
	APIs.Define(licenseInfoAPI)     // ライセンスプラン
	APIs.Define(loadBalancerAPI)    // ロードバランサ
	APIs.Define(mobileGatewayAPI)   // モバイルゲートウェイ
	APIs.Define(nfsAPI)             // NFS
	APIs.Define(noteAPI)            // スタートアップスクリプト
	APIs.Define(packetFilterAPI)    // パケットフィルタ
	APIs.Define(privateHostAPI)     // 専有ホスト
	APIs.Define(privateHostPlanAPI) // 専有ホストプラン
	APIs.Define(proxyLBAPI)         // エンハンスドロードバランサ
	APIs.Define(regionAPI)          // リージョン
	APIs.Define(serverAPI)          // サーバ
	APIs.Define(serverPlanAPI)      // サーバプラン
	APIs.Define(serviceClassAPI)    // サービスクラス(価格)
	APIs.Define(simAPI)             // SIM
	APIs.Define(simpleMonitorAPI)   // シンプル監視
	APIs.Define(sshKeyAPI)          // 公開鍵
	APIs.Define(switchAPI)          // スイッチ
	APIs.Define(vpcRouterAPI)       // VPCルータ
	APIs.Define(webaccelAPI)        // ウェブアクセラレータ
	APIs.Define(zoneAPI)            // ゾーン
}
