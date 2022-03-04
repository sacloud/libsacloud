# libsacloud

[![Go Reference](https://pkg.go.dev/badge/github.com/sacloud/libsacloud/v2.svg)](https://pkg.go.dev/github.com/sacloud/libsacloud/v2)
![Tests](https://github.com/sacloud/libsacloud/workflows/Tests/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/sacloud/libsacloud)](https://goreportcard.com/report/github.com/sacloud/libsacloud)

Library for SakuraCloud API

:bulb: 現在次期バージョンの開発が進められています。  
次期バージョンは以下2つのリポジトリに分割されます。

- [sacloud/sacloud-go](https://github.com/sacloud/sacloud-go): 高レベルAPI(libsacloudの`helper`パッケージなど)
- [sacloud/iaas-api-go](https://github.com/sacloud/iaas-api-go): IaaS関連API(libsacloudの`sacloud`パッケージなど)

これらのリリース後、sacloud配下の主要プロダクトがlibsacloudから移行完了するまでlibsacloudの開発は継続されます。

## Installation

Use go get.

    go get github.com/sacloud/libsacloud/v2

Then import the `sacloud` package into your own code.

    import "github.com/sacloud/libsacloud/v2/sacloud"

## License

  `libsacloud` Copyright (C) 2016-2022 The Libsacloud Authors.

  This project is published under [Apache 2.0 License](LICENSE).
