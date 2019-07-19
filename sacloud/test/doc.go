/*
APIのテストのためのパッケージ

テスト実行時は実施したいテストに応じて以下の環境変数を定義すること

- E2Eテスト

     TESTACC=1
     SAKURACLOUD_ACCESS_TOKEN
     SAKURACLOUD_ACCESS_TOKEN_SECRET

- エンハンスドロードバランサ

     値はさくらインターネットがお客様向けに提供するグローバルIPアドレス(クラウド/VPS/専用サーバなど)を指定すること
     (Unit Testの場合は任意のIPアドレスを指定可能)

     SAKURACLOUD_PROXYLB_SERVER0
     SAKURACLOUD_PROXYLB_SERVER1
     SAKURACLOUD_PROXYLB_SERVER2

- エンハンスドロードバランサのLet's Encrypt設定

     SAKURACLOUD_PROXYLB_SERVER0
     SAKURACLOUD_PROXYLB_SERVER1
     SAKURACLOUD_PROXYLB_COMMON_NAME => 証明書発行対象となるFQDN
     SAKURACLOUD_PROXYLB_ZONE_NAME   => さくらのクラウドDNSに登録されているDNSゾーン名

- IPv4アドレスの逆引き設定

     SAKURACLOUD_IPADDRESS
     SAKURACLOUD_HOSTNAME

- IPv6アドレスの逆引き設定

     SAKURACLOUD_IPV6ADDRESS
     SAKURACLOUD_IPV6HOSTNAME

- モバイルゲートウェイ/SIM

セキュアモバイルの利用権限のあるアカウントを利用すること

     SAKURACLOUD_SIM_ICCID
     SAKURACLOUD_SIM_PASSCODE

- SIMのログ

     SAKURACLOUD_SIM_ID => ログの参照対象のSIMのリソースID


ウェブアクセラレータ

証明書設定:
     SAKURACLOUD_WEBACCEL_SITE_ID
     SAKURACLOUD_WEBACCEL_CERT
     SAKURACLOUD_WEBACCEL_KEY

キャッシュ全削除:
     SAKURACLOUD_WEBACCEL_DOMAIN

キャッシュ削除(URL指定):
     SAKURACLOUD_WEBACCEL_URLS => キャッシュ削除対象のURLをカンマ区切りで指定
*/
package test
