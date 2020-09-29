// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ostype is define OS type of SakuraCloud public archive
package ostype

//go:generate stringer -type=ArchiveOSType

// ArchiveOSType パブリックアーカイブOS種別
type ArchiveOSType int

const (
	// Custom OS種別:カスタム
	Custom ArchiveOSType = iota

	// CentOS OS種別:CentOS
	CentOS
	// CentOS8 OS種別:CentOS8
	CentOS8
	// CentOS7 OS種別:CentOS7
	CentOS7
	// CentOS6 OS種別:CentOS6
	CentOS6

	// Ubuntu OS種別:Ubuntu
	Ubuntu
	// Ubuntu1804 OS種別:Ubuntu(Focal Fossa)
	Ubuntu2004
	// Ubuntu1804 OS種別:Ubuntu(Bionic)
	Ubuntu1804
	// Ubuntu1604 OS種別:Ubuntu(Xenial)
	Ubuntu1604

	// Debian OS種別:Debian
	Debian
	// Debian10 OS種別:Debian10
	Debian10
	// Debian9 OS種別:Debian9
	Debian9

	// CoreOS OS種別:CoreOS
	CoreOS

	// RancherOS OS種別:RancherOS
	RancherOS
	// K3OS OS種別: k3OS
	K3OS

	// Kusanagi OS種別:Kusanagi(CentOS)
	Kusanagi

	// FreeBSD OS種別:FreeBSD
	FreeBSD

	// Windows2016 OS種別:Windows Server 2016 Datacenter Edition
	Windows2016
	// Windows2016RDS OS種別:Windows Server 2016 RDS
	Windows2016RDS
	// Windows2016RDSOffice OS種別:Windows Server 2016 RDS(Office)
	Windows2016RDSOffice
	// Windows2016SQLServerWeb OS種別:Windows Server 2016 SQLServer(Web)
	Windows2016SQLServerWeb
	// Windows2016SQLServerStandard OS種別:Windows Server 2016 SQLServer 2016(Standard)
	Windows2016SQLServerStandard
	// Windows2016SQLServer2017Standard OS種別:Windows Server 2016 SQLServer 2017(Standard)
	Windows2016SQLServer2017Standard
	// Windows2016SQLServer2017Enterprise OS種別:Windows Server 2016 SQLServer 2017(Enterprise)
	Windows2016SQLServer2017Enterprise
	// Windows2016SQLServerStandardAll OS種別:Windows Server 2016 SQLServer(Standard) + RDS + Office
	Windows2016SQLServerStandardAll
	// Windows2016SQLServer2017StandardAll OS種別:Windows Server 2016 SQLServer 2017(Standard) + RDS + Office
	Windows2016SQLServer2017StandardAll

	// Windows2019 OS種別:Windows Server 2019 Datacenter Edition
	Windows2019
	// Windows2019RDS OS種別:Windows Server 2019 RDS
	Windows2019RDS

	// Windows2019RDSOffice2019 OS種別:Windows Server 2019 RDS + Office 2019
	Windows2019RDSOffice2019

	// Windows2019SQLServer2017Web OS種別:Windows Server 2019 + SQLServer 2017(Web)
	Windows2019SQLServer2017Web
	// Windows2019SQLServer2019Web OS種別:Windows Server 2019 + SQLServer 2019(Web)
	Windows2019SQLServer2019Web

	// Windows2019SQLServer2017Standard OS種別:Windows Server 2019 + SQLServer 2017(Standard)
	Windows2019SQLServer2017Standard
	// Windows2019SQLServer2019Standard OS種別:Windows Server 2019 + SQLServer 2019(Standard)
	Windows2019SQLServer2019Standard

	// Windows2019SQLServer2017Enterprise OS種別:Windows Server 2019 + SQLServer 2017(Enterprise)
	Windows2019SQLServer2017Enterprise
	// Windows2019SQLServer2019Enterprise OS種別:Windows Server 2019 + SQLServer 2019(Enterprise)
	Windows2019SQLServer2019Enterprise

	// Windows2019SQLServer2017StandardAll OS種別:Windows Server 2019 + SQLServer 2017(Standard) + RDS + Office
	Windows2019SQLServer2017StandardAll
	// Windows2019SQLServer2019StandardAll OS種別:Windows Server 2019 + SQLServer 2019(Standard) + RDS + Office
	Windows2019SQLServer2019StandardAll
)

// ArchiveOSTypes アーカイブ種別のリスト
var ArchiveOSTypes = []ArchiveOSType{
	CentOS,
	CentOS8,
	CentOS7,
	CentOS6,
	Ubuntu,
	Ubuntu2004,
	Ubuntu1804,
	Ubuntu1604,
	Debian,
	Debian10,
	Debian9,
	CoreOS,
	RancherOS,
	K3OS,
	Kusanagi,
	FreeBSD,
	Windows2016,
	Windows2016RDS,
	Windows2016RDSOffice,
	Windows2016SQLServerWeb,
	Windows2016SQLServerStandard,
	Windows2016SQLServer2017Standard,
	Windows2016SQLServer2017Enterprise,
	Windows2016SQLServerStandardAll,
	Windows2016SQLServer2017StandardAll,
	Windows2019,
	Windows2019RDS,
	Windows2019RDSOffice2019,
	Windows2019SQLServer2017Web,
	Windows2019SQLServer2019Web,
	Windows2019SQLServer2017Standard,
	Windows2019SQLServer2019Standard,
	Windows2019SQLServer2017Enterprise,
	Windows2019SQLServer2019Enterprise,
	Windows2019SQLServer2017StandardAll,
	Windows2019SQLServer2019StandardAll,
}

// OSTypeShortNames OSTypeとして利用できる文字列のリスト
var OSTypeShortNames = []string{
	"centos", "centos8", "centos7", "centos6",
	"ubuntu", "ubuntu2004", "ubuntu1804", "ubuntu1604",
	"debian", "debian10", "debian9",
	"coreos", "rancheros", "k3os", "kusanagi", "freebsd",
	"windows2016", "windows2016-rds", "windows2016-rds-office",
	"windows2016-sql-web", "windows2016-sql-standard", "windows2016-sql-standard-all",
	"windows2016-sql2017-standard", "windows2016-sql2017-enterprise", "windows2016-sql2017-standard-all",
	"windows2019", "windows2019-rds", "windows2019-rds-office2019",
	"windows2019-sql2017-web", "windows2019-sql2019-web",
	"windows2019-sql2017-standard", "windows2019-sql2019-standard",
	"windows2019-sql2017-enterprise", "windows2019-sql2019-enterprise",
	"windows2019-sql2017-standard-all", "windows2019-sql2019-standard-all",
}

// IsWindows Windowsか
func (o ArchiveOSType) IsWindows() bool {
	switch o {
	case Windows2016, Windows2016RDS, Windows2016RDSOffice,
		Windows2016SQLServerWeb, Windows2016SQLServerStandard, Windows2016SQLServerStandardAll,
		Windows2016SQLServer2017Standard, Windows2016SQLServer2017Enterprise, Windows2016SQLServer2017StandardAll,
		Windows2019, Windows2019RDS,
		Windows2019RDSOffice2019,
		Windows2019SQLServer2017Web, Windows2019SQLServer2019Web,
		Windows2019SQLServer2017Standard, Windows2019SQLServer2019Standard,
		Windows2019SQLServer2017Enterprise, Windows2019SQLServer2019Enterprise,
		Windows2019SQLServer2017StandardAll, Windows2019SQLServer2019StandardAll:
		return true
	default:
		return false
	}
}

// IsSupportDiskEdit ディスクの修正機能をフルサポートしているか(Windowsは一部サポートのためfalseを返す)
func (o ArchiveOSType) IsSupportDiskEdit() bool {
	switch o {
	case CentOS, CentOS8, CentOS7, CentOS6,
		Ubuntu, Ubuntu2004, Ubuntu1804, Ubuntu1604,
		Debian, Debian10, Debian9,
		CoreOS, RancherOS, K3OS, Kusanagi, FreeBSD:
		return true
	default:
		return false
	}
}

// StrToOSType 文字列からArchiveOSTypesへの変換
func StrToOSType(osType string) ArchiveOSType {
	switch osType {
	case "centos":
		return CentOS
	case "centos8":
		return CentOS8
	case "centos7":
		return CentOS7
	case "centos6":
		return CentOS6
	case "ubuntu":
		return Ubuntu
	case "ubuntu2004":
		return Ubuntu2004
	case "ubuntu1804":
		return Ubuntu1804
	case "ubuntu1604":
		return Ubuntu1604
	case "debian":
		return Debian
	case "debian10":
		return Debian10
	case "debian9":
		return Debian9
	case "coreos":
		return CoreOS
	case "rancheros":
		return RancherOS
	case "k3os":
		return K3OS
	case "kusanagi":
		return Kusanagi
	case "freebsd":
		return FreeBSD
	case "windows2016":
		return Windows2016
	case "windows2016-rds":
		return Windows2016RDS
	case "windows2016-rds-office":
		return Windows2016RDSOffice
	case "windows2016-sql-web":
		return Windows2016SQLServerWeb
	case "windows2016-sql-standard":
		return Windows2016SQLServerStandard
	case "windows2016-sql2017-standard":
		return Windows2016SQLServer2017Standard
	case "windows2016-sql2017-enterprise":
		return Windows2016SQLServer2017Enterprise
	case "windows2016-sql-standard-all":
		return Windows2016SQLServerStandardAll
	case "windows2016-sql2017-standard-all":
		return Windows2016SQLServer2017StandardAll
	case "windows2019":
		return Windows2019
	case "windows2019-rds":
		return Windows2019RDS
	case "windows2019-rds-office2019":
		return Windows2019RDSOffice2019
	case "windows2019-sql2017-web":
		return Windows2019SQLServer2017Web
	case "windows2019-sql2019-web":
		return Windows2019SQLServer2019Web
	case "windows2019-sql2017-standard":
		return Windows2019SQLServer2017Standard
	case "windows2019-sql2019-standard":
		return Windows2019SQLServer2019Standard
	case "windows2019-sql2017-enterprise":
		return Windows2019SQLServer2017Enterprise
	case "windows2019-sql2019-enterprise":
		return Windows2019SQLServer2019Enterprise
	case "windows2019-sql2017-standard-all":
		return Windows2019SQLServer2017StandardAll
	case "windows2019-sql2019-standard-all":
		return Windows2019SQLServer2019StandardAll
	default:
		return Custom
	}
}
