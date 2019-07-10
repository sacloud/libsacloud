// Package test sacloud.xxxAPIのテスト
//
// E2Eテスト(Acceptance Test)を実行する場合は以下の環境変数を設定すること
//
// - TESTACC=1
// - SAKURACLOUD_ACCESS_TOKEN
// - SAKURACLOUD_ACCESS_TOKEN_SECRET
//
// エンハンスドロードバランサのテストを実行するには以下の環境変数も設定すること
// 値はさくらインターネットがお客様向けに提供するグローバルIPアドレス(クラウド/VPS/専用サーバなど)を指定すること
// (Unit Testの場合は任意のIPアドレスを指定可能)
//
// - SAKURACLOUD_PROXYLB_SERVER0
// - SAKURACLOUD_PROXYLB_SERVER1
// - SAKURACLOUD_PROXYLB_SERVER2
//
// エンハンスドロードバランサのLet's Encrypt設定のテストを実行するには以下の環境変数を設定すること
//
// - SAKURACLOUD_PROXYLB_SERVER0
// - SAKURACLOUD_PROXYLB_SERVER1
// - SAKURACLOUD_PROXYLB_COMMON_NAME => 証明書発行対象となるFQDN
// - SAKURACLOUD_PROXYLB_ZONE_NAME   => さくらのクラウドDNSに登録されているDNSゾーン名
//
package test
