package ostype

import "github.com/sacloud/libsacloud/v2/sacloud/ostype"

//go:generate stringer -type=UnixPublicArchiveType
//go:generate stringer -type=WindowsPublicArchiveType

// UnixPublicArchiveType Unix系パブリックアーカイブ種別
type UnixPublicArchiveType int

const (
	// CentOS OS種別:CentOS
	CentOS UnixPublicArchiveType = iota
	// CentOS6 OS種別:CentOS6
	CentOS6
	// Ubuntu OS種別:Ubuntu
	Ubuntu
	// Debian OS種別:Debian
	Debian
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
)

// UnixPublicArchives UnixPublicArchiveTypeとsacloud/ostype/ArchiveOSTypeの対応マップ
var UnixPublicArchives = map[UnixPublicArchiveType]ostype.ArchiveOSType{
	CentOS:    ostype.CentOS,
	CentOS6:   ostype.CentOS6,
	Ubuntu:    ostype.Ubuntu,
	Debian:    ostype.Debian,
	CoreOS:    ostype.CoreOS,
	RancherOS: ostype.RancherOS,
	K3OS:      ostype.K3OS,
	Kusanagi:  ostype.Kusanagi,
	FreeBSD:   ostype.FreeBSD,
}
