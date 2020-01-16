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

// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-interfaces'; DO NOT EDIT

package sacloud

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

/*************************************************
* ArchiveAPI
*************************************************/

// ArchiveAPI is interface for operate Archive resource
type ArchiveAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*ArchiveFindResult, error)
	Create(ctx context.Context, zone string, param *ArchiveCreateRequest) (*Archive, error)
	CreateBlank(ctx context.Context, zone string, param *ArchiveCreateBlankRequest) (*Archive, *FTPServer, error)
	Read(ctx context.Context, zone string, id types.ID) (*Archive, error)
	Update(ctx context.Context, zone string, id types.ID, param *ArchiveUpdateRequest) (*Archive, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	OpenFTP(ctx context.Context, zone string, id types.ID, openOption *OpenFTPRequest) (*FTPServer, error)
	CloseFTP(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* AuthStatusAPI
*************************************************/

// AuthStatusAPI is interface for operate AuthStatus resource
type AuthStatusAPI interface {
	Read(ctx context.Context) (*AuthStatus, error)
}

/*************************************************
* AutoBackupAPI
*************************************************/

// AutoBackupAPI is interface for operate AutoBackup resource
type AutoBackupAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*AutoBackupFindResult, error)
	Create(ctx context.Context, zone string, param *AutoBackupCreateRequest) (*AutoBackup, error)
	Read(ctx context.Context, zone string, id types.ID) (*AutoBackup, error)
	Update(ctx context.Context, zone string, id types.ID, param *AutoBackupUpdateRequest) (*AutoBackup, error)
	UpdateSettings(ctx context.Context, zone string, id types.ID, param *AutoBackupUpdateSettingsRequest) (*AutoBackup, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* BillAPI
*************************************************/

// BillAPI is interface for operate Bill resource
type BillAPI interface {
	ByContract(ctx context.Context, accountID types.ID) (*BillByContractResult, error)
	ByContractYear(ctx context.Context, accountID types.ID, year int) (*BillByContractYearResult, error)
	ByContractYearMonth(ctx context.Context, accountID types.ID, year int, month int) (*BillByContractYearMonthResult, error)
	Read(ctx context.Context, id types.ID) (*BillReadResult, error)
	Details(ctx context.Context, MemberCode string, id types.ID) (*BillDetailsResult, error)
	DetailsCSV(ctx context.Context, MemberCode string, id types.ID) (*BillDetailCSV, error)
}

/*************************************************
* BridgeAPI
*************************************************/

// BridgeAPI is interface for operate Bridge resource
type BridgeAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*BridgeFindResult, error)
	Create(ctx context.Context, zone string, param *BridgeCreateRequest) (*Bridge, error)
	Read(ctx context.Context, zone string, id types.ID) (*Bridge, error)
	Update(ctx context.Context, zone string, id types.ID, param *BridgeUpdateRequest) (*Bridge, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* CDROMAPI
*************************************************/

// CDROMAPI is interface for operate CDROM resource
type CDROMAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*CDROMFindResult, error)
	Create(ctx context.Context, zone string, param *CDROMCreateRequest) (*CDROM, *FTPServer, error)
	Read(ctx context.Context, zone string, id types.ID) (*CDROM, error)
	Update(ctx context.Context, zone string, id types.ID, param *CDROMUpdateRequest) (*CDROM, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	OpenFTP(ctx context.Context, zone string, id types.ID, openOption *OpenFTPRequest) (*FTPServer, error)
	CloseFTP(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* ContainerRegistryAPI
*************************************************/

// ContainerRegistryAPI is interface for operate ContainerRegistry resource
type ContainerRegistryAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*ContainerRegistryFindResult, error)
	Create(ctx context.Context, param *ContainerRegistryCreateRequest) (*ContainerRegistry, error)
	Read(ctx context.Context, id types.ID) (*ContainerRegistry, error)
	Update(ctx context.Context, id types.ID, param *ContainerRegistryUpdateRequest) (*ContainerRegistry, error)
	UpdateSettings(ctx context.Context, id types.ID, param *ContainerRegistryUpdateSettingsRequest) (*ContainerRegistry, error)
	Delete(ctx context.Context, id types.ID) error
	ListUsers(ctx context.Context, id types.ID) (*ContainerRegistryUsers, error)
	AddUser(ctx context.Context, id types.ID, param *ContainerRegistryUserCreateRequest) error
	UpdateUser(ctx context.Context, id types.ID, username string, param *ContainerRegistryUserUpdateRequest) error
	DeleteUser(ctx context.Context, id types.ID, username string) error
}

/*************************************************
* CouponAPI
*************************************************/

// CouponAPI is interface for operate Coupon resource
type CouponAPI interface {
	Find(ctx context.Context, accountID types.ID) (*CouponFindResult, error)
}

/*************************************************
* DatabaseAPI
*************************************************/

// DatabaseAPI is interface for operate Database resource
type DatabaseAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*DatabaseFindResult, error)
	Create(ctx context.Context, zone string, param *DatabaseCreateRequest) (*Database, error)
	Read(ctx context.Context, zone string, id types.ID) (*Database, error)
	Update(ctx context.Context, zone string, id types.ID, param *DatabaseUpdateRequest) (*Database, error)
	UpdateSettings(ctx context.Context, zone string, id types.ID, param *DatabaseUpdateSettingsRequest) (*Database, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Config(ctx context.Context, zone string, id types.ID) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	MonitorCPU(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*CPUTimeActivity, error)
	MonitorDisk(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*DiskActivity, error)
	MonitorInterface(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*InterfaceActivity, error)
	MonitorDatabase(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*DatabaseActivity, error)
	Status(ctx context.Context, zone string, id types.ID) (*DatabaseStatus, error)
}

/*************************************************
* DiskAPI
*************************************************/

// DiskAPI is interface for operate Disk resource
type DiskAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*DiskFindResult, error)
	Create(ctx context.Context, zone string, createParam *DiskCreateRequest, distantFrom []types.ID) (*Disk, error)
	Config(ctx context.Context, zone string, id types.ID, edit *DiskEditRequest) error
	CreateWithConfig(ctx context.Context, zone string, createParam *DiskCreateRequest, editParam *DiskEditRequest, bootAtAvailable bool, distantFrom []types.ID) (*Disk, error)
	ToBlank(ctx context.Context, zone string, id types.ID) error
	ResizePartition(ctx context.Context, zone string, id types.ID, param *DiskResizePartitionRequest) error
	ConnectToServer(ctx context.Context, zone string, id types.ID, serverID types.ID) error
	DisconnectFromServer(ctx context.Context, zone string, id types.ID) error
	Install(ctx context.Context, zone string, id types.ID, installParam *DiskInstallRequest, distantFrom []types.ID) (*Disk, error)
	Read(ctx context.Context, zone string, id types.ID) (*Disk, error)
	Update(ctx context.Context, zone string, id types.ID, param *DiskUpdateRequest) (*Disk, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Monitor(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*DiskActivity, error)
}

/*************************************************
* DiskPlanAPI
*************************************************/

// DiskPlanAPI is interface for operate DiskPlan resource
type DiskPlanAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*DiskPlanFindResult, error)
	Read(ctx context.Context, zone string, id types.ID) (*DiskPlan, error)
}

/*************************************************
* DNSAPI
*************************************************/

// DNSAPI is interface for operate DNS resource
type DNSAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*DNSFindResult, error)
	Create(ctx context.Context, param *DNSCreateRequest) (*DNS, error)
	Read(ctx context.Context, id types.ID) (*DNS, error)
	Update(ctx context.Context, id types.ID, param *DNSUpdateRequest) (*DNS, error)
	UpdateSettings(ctx context.Context, id types.ID, param *DNSUpdateSettingsRequest) (*DNS, error)
	Delete(ctx context.Context, id types.ID) error
}

/*************************************************
* GSLBAPI
*************************************************/

// GSLBAPI is interface for operate GSLB resource
type GSLBAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*GSLBFindResult, error)
	Create(ctx context.Context, param *GSLBCreateRequest) (*GSLB, error)
	Read(ctx context.Context, id types.ID) (*GSLB, error)
	Update(ctx context.Context, id types.ID, param *GSLBUpdateRequest) (*GSLB, error)
	UpdateSettings(ctx context.Context, id types.ID, param *GSLBUpdateSettingsRequest) (*GSLB, error)
	Delete(ctx context.Context, id types.ID) error
}

/*************************************************
* IconAPI
*************************************************/

// IconAPI is interface for operate Icon resource
type IconAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*IconFindResult, error)
	Create(ctx context.Context, param *IconCreateRequest) (*Icon, error)
	Read(ctx context.Context, id types.ID) (*Icon, error)
	Update(ctx context.Context, id types.ID, param *IconUpdateRequest) (*Icon, error)
	Delete(ctx context.Context, id types.ID) error
}

/*************************************************
* InterfaceAPI
*************************************************/

// InterfaceAPI is interface for operate Interface resource
type InterfaceAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*InterfaceFindResult, error)
	Create(ctx context.Context, zone string, param *InterfaceCreateRequest) (*Interface, error)
	Read(ctx context.Context, zone string, id types.ID) (*Interface, error)
	Update(ctx context.Context, zone string, id types.ID, param *InterfaceUpdateRequest) (*Interface, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Monitor(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*InterfaceActivity, error)
	ConnectToSharedSegment(ctx context.Context, zone string, id types.ID) error
	ConnectToSwitch(ctx context.Context, zone string, id types.ID, switchID types.ID) error
	DisconnectFromSwitch(ctx context.Context, zone string, id types.ID) error
	ConnectToPacketFilter(ctx context.Context, zone string, id types.ID, packetFilterID types.ID) error
	DisconnectFromPacketFilter(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* InternetAPI
*************************************************/

// InternetAPI is interface for operate Internet resource
type InternetAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*InternetFindResult, error)
	Create(ctx context.Context, zone string, param *InternetCreateRequest) (*Internet, error)
	Read(ctx context.Context, zone string, id types.ID) (*Internet, error)
	Update(ctx context.Context, zone string, id types.ID, param *InternetUpdateRequest) (*Internet, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	UpdateBandWidth(ctx context.Context, zone string, id types.ID, param *InternetUpdateBandWidthRequest) (*Internet, error)
	AddSubnet(ctx context.Context, zone string, id types.ID, param *InternetAddSubnetRequest) (*InternetSubnetOperationResult, error)
	UpdateSubnet(ctx context.Context, zone string, id types.ID, subnetID types.ID, param *InternetUpdateSubnetRequest) (*InternetSubnetOperationResult, error)
	DeleteSubnet(ctx context.Context, zone string, id types.ID, subnetID types.ID) error
	Monitor(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*RouterActivity, error)
	EnableIPv6(ctx context.Context, zone string, id types.ID) (*IPv6NetInfo, error)
	DisableIPv6(ctx context.Context, zone string, id types.ID, ipv6netID types.ID) error
}

/*************************************************
* InternetPlanAPI
*************************************************/

// InternetPlanAPI is interface for operate InternetPlan resource
type InternetPlanAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*InternetPlanFindResult, error)
	Read(ctx context.Context, zone string, id types.ID) (*InternetPlan, error)
}

/*************************************************
* IPAddressAPI
*************************************************/

// IPAddressAPI is interface for operate IPAddress resource
type IPAddressAPI interface {
	List(ctx context.Context, zone string) (*IPAddressListResult, error)
	Read(ctx context.Context, zone string, ipAddress string) (*IPAddress, error)
	UpdateHostName(ctx context.Context, zone string, ipAddress string, hostName string) (*IPAddress, error)
}

/*************************************************
* IPv6NetAPI
*************************************************/

// IPv6NetAPI is interface for operate IPv6Net resource
type IPv6NetAPI interface {
	List(ctx context.Context, zone string) (*IPv6NetListResult, error)
	Read(ctx context.Context, zone string, id types.ID) (*IPv6Net, error)
}

/*************************************************
* IPv6AddrAPI
*************************************************/

// IPv6AddrAPI is interface for operate IPv6Addr resource
type IPv6AddrAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*IPv6AddrFindResult, error)
	Create(ctx context.Context, zone string, param *IPv6AddrCreateRequest) (*IPv6Addr, error)
	Read(ctx context.Context, zone string, ipv6addr string) (*IPv6Addr, error)
	Update(ctx context.Context, zone string, ipv6addr string, param *IPv6AddrUpdateRequest) (*IPv6Addr, error)
	Delete(ctx context.Context, zone string, ipv6addr string) error
}

/*************************************************
* LicenseAPI
*************************************************/

// LicenseAPI is interface for operate License resource
type LicenseAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*LicenseFindResult, error)
	Create(ctx context.Context, param *LicenseCreateRequest) (*License, error)
	Read(ctx context.Context, id types.ID) (*License, error)
	Update(ctx context.Context, id types.ID, param *LicenseUpdateRequest) (*License, error)
	Delete(ctx context.Context, id types.ID) error
}

/*************************************************
* LicenseInfoAPI
*************************************************/

// LicenseInfoAPI is interface for operate LicenseInfo resource
type LicenseInfoAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*LicenseInfoFindResult, error)
	Read(ctx context.Context, id types.ID) (*LicenseInfo, error)
}

/*************************************************
* LoadBalancerAPI
*************************************************/

// LoadBalancerAPI is interface for operate LoadBalancer resource
type LoadBalancerAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*LoadBalancerFindResult, error)
	Create(ctx context.Context, zone string, param *LoadBalancerCreateRequest) (*LoadBalancer, error)
	Read(ctx context.Context, zone string, id types.ID) (*LoadBalancer, error)
	Update(ctx context.Context, zone string, id types.ID, param *LoadBalancerUpdateRequest) (*LoadBalancer, error)
	UpdateSettings(ctx context.Context, zone string, id types.ID, param *LoadBalancerUpdateSettingsRequest) (*LoadBalancer, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Config(ctx context.Context, zone string, id types.ID) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	MonitorInterface(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*InterfaceActivity, error)
	Status(ctx context.Context, zone string, id types.ID) (*LoadBalancerStatusResult, error)
}

/*************************************************
* MobileGatewayAPI
*************************************************/

// MobileGatewayAPI is interface for operate MobileGateway resource
type MobileGatewayAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*MobileGatewayFindResult, error)
	Create(ctx context.Context, zone string, param *MobileGatewayCreateRequest) (*MobileGateway, error)
	Read(ctx context.Context, zone string, id types.ID) (*MobileGateway, error)
	Update(ctx context.Context, zone string, id types.ID, param *MobileGatewayUpdateRequest) (*MobileGateway, error)
	UpdateSettings(ctx context.Context, zone string, id types.ID, param *MobileGatewayUpdateSettingsRequest) (*MobileGateway, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Config(ctx context.Context, zone string, id types.ID) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	ConnectToSwitch(ctx context.Context, zone string, id types.ID, switchID types.ID) error
	DisconnectFromSwitch(ctx context.Context, zone string, id types.ID) error
	GetDNS(ctx context.Context, zone string, id types.ID) (*MobileGatewayDNSSetting, error)
	SetDNS(ctx context.Context, zone string, id types.ID, param *MobileGatewayDNSSetting) error
	GetSIMRoutes(ctx context.Context, zone string, id types.ID) ([]*MobileGatewaySIMRoute, error)
	SetSIMRoutes(ctx context.Context, zone string, id types.ID, param []*MobileGatewaySIMRouteParam) error
	ListSIM(ctx context.Context, zone string, id types.ID) ([]*MobileGatewaySIMInfo, error)
	AddSIM(ctx context.Context, zone string, id types.ID, param *MobileGatewayAddSIMRequest) error
	DeleteSIM(ctx context.Context, zone string, id types.ID, simID types.ID) error
	Logs(ctx context.Context, zone string, id types.ID) ([]*MobileGatewaySIMLogs, error)
	GetTrafficConfig(ctx context.Context, zone string, id types.ID) (*MobileGatewayTrafficControl, error)
	SetTrafficConfig(ctx context.Context, zone string, id types.ID, param *MobileGatewayTrafficControl) error
	DeleteTrafficConfig(ctx context.Context, zone string, id types.ID) error
	TrafficStatus(ctx context.Context, zone string, id types.ID) (*MobileGatewayTrafficStatus, error)
	MonitorInterface(ctx context.Context, zone string, id types.ID, index int, condition *MonitorCondition) (*InterfaceActivity, error)
}

/*************************************************
* NFSAPI
*************************************************/

// NFSAPI is interface for operate NFS resource
type NFSAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*NFSFindResult, error)
	Create(ctx context.Context, zone string, param *NFSCreateRequest) (*NFS, error)
	Read(ctx context.Context, zone string, id types.ID) (*NFS, error)
	Update(ctx context.Context, zone string, id types.ID, param *NFSUpdateRequest) (*NFS, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	MonitorFreeDiskSize(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*FreeDiskSizeActivity, error)
	MonitorInterface(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*InterfaceActivity, error)
}

/*************************************************
* NoteAPI
*************************************************/

// NoteAPI is interface for operate Note resource
type NoteAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*NoteFindResult, error)
	Create(ctx context.Context, param *NoteCreateRequest) (*Note, error)
	Read(ctx context.Context, id types.ID) (*Note, error)
	Update(ctx context.Context, id types.ID, param *NoteUpdateRequest) (*Note, error)
	Delete(ctx context.Context, id types.ID) error
}

/*************************************************
* PacketFilterAPI
*************************************************/

// PacketFilterAPI is interface for operate PacketFilter resource
type PacketFilterAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*PacketFilterFindResult, error)
	Create(ctx context.Context, zone string, param *PacketFilterCreateRequest) (*PacketFilter, error)
	Read(ctx context.Context, zone string, id types.ID) (*PacketFilter, error)
	Update(ctx context.Context, zone string, id types.ID, param *PacketFilterUpdateRequest) (*PacketFilter, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* PrivateHostAPI
*************************************************/

// PrivateHostAPI is interface for operate PrivateHost resource
type PrivateHostAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*PrivateHostFindResult, error)
	Create(ctx context.Context, zone string, param *PrivateHostCreateRequest) (*PrivateHost, error)
	Read(ctx context.Context, zone string, id types.ID) (*PrivateHost, error)
	Update(ctx context.Context, zone string, id types.ID, param *PrivateHostUpdateRequest) (*PrivateHost, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* PrivateHostPlanAPI
*************************************************/

// PrivateHostPlanAPI is interface for operate PrivateHostPlan resource
type PrivateHostPlanAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*PrivateHostPlanFindResult, error)
	Read(ctx context.Context, zone string, id types.ID) (*PrivateHostPlan, error)
}

/*************************************************
* ProxyLBAPI
*************************************************/

// ProxyLBAPI is interface for operate ProxyLB resource
type ProxyLBAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*ProxyLBFindResult, error)
	Create(ctx context.Context, param *ProxyLBCreateRequest) (*ProxyLB, error)
	Read(ctx context.Context, id types.ID) (*ProxyLB, error)
	Update(ctx context.Context, id types.ID, param *ProxyLBUpdateRequest) (*ProxyLB, error)
	UpdateSettings(ctx context.Context, id types.ID, param *ProxyLBUpdateSettingsRequest) (*ProxyLB, error)
	Delete(ctx context.Context, id types.ID) error
	ChangePlan(ctx context.Context, id types.ID, param *ProxyLBChangePlanRequest) (*ProxyLB, error)
	GetCertificates(ctx context.Context, id types.ID) (*ProxyLBCertificates, error)
	SetCertificates(ctx context.Context, id types.ID, param *ProxyLBSetCertificatesRequest) (*ProxyLBCertificates, error)
	DeleteCertificates(ctx context.Context, id types.ID) error
	RenewLetsEncryptCert(ctx context.Context, id types.ID) error
	HealthStatus(ctx context.Context, id types.ID) (*ProxyLBHealth, error)
	MonitorConnection(ctx context.Context, id types.ID, condition *MonitorCondition) (*ConnectionActivity, error)
}

/*************************************************
* RegionAPI
*************************************************/

// RegionAPI is interface for operate Region resource
type RegionAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*RegionFindResult, error)
	Read(ctx context.Context, id types.ID) (*Region, error)
}

/*************************************************
* ServerAPI
*************************************************/

// ServerAPI is interface for operate Server resource
type ServerAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*ServerFindResult, error)
	Create(ctx context.Context, zone string, param *ServerCreateRequest) (*Server, error)
	Read(ctx context.Context, zone string, id types.ID) (*Server, error)
	Update(ctx context.Context, zone string, id types.ID, param *ServerUpdateRequest) (*Server, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	DeleteWithDisks(ctx context.Context, zone string, id types.ID, disks *ServerDeleteWithDisksRequest) error
	ChangePlan(ctx context.Context, zone string, id types.ID, plan *ServerChangePlanRequest) (*Server, error)
	InsertCDROM(ctx context.Context, zone string, id types.ID, insertParam *InsertCDROMRequest) error
	EjectCDROM(ctx context.Context, zone string, id types.ID, ejectParam *EjectCDROMRequest) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	SendKey(ctx context.Context, zone string, id types.ID, keyboardParam *SendKeyRequest) error
	GetVNCProxy(ctx context.Context, zone string, id types.ID) (*VNCProxyInfo, error)
	Monitor(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*CPUTimeActivity, error)
}

/*************************************************
* ServerPlanAPI
*************************************************/

// ServerPlanAPI is interface for operate ServerPlan resource
type ServerPlanAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*ServerPlanFindResult, error)
	Read(ctx context.Context, zone string, id types.ID) (*ServerPlan, error)
}

/*************************************************
* ServiceClassAPI
*************************************************/

// ServiceClassAPI is interface for operate ServiceClass resource
type ServiceClassAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*ServiceClassFindResult, error)
}

/*************************************************
* SIMAPI
*************************************************/

// SIMAPI is interface for operate SIM resource
type SIMAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*SIMFindResult, error)
	Create(ctx context.Context, param *SIMCreateRequest) (*SIM, error)
	Read(ctx context.Context, id types.ID) (*SIM, error)
	Update(ctx context.Context, id types.ID, param *SIMUpdateRequest) (*SIM, error)
	Delete(ctx context.Context, id types.ID) error
	Activate(ctx context.Context, id types.ID) error
	Deactivate(ctx context.Context, id types.ID) error
	AssignIP(ctx context.Context, id types.ID, param *SIMAssignIPRequest) error
	ClearIP(ctx context.Context, id types.ID) error
	IMEILock(ctx context.Context, id types.ID, param *SIMIMEILockRequest) error
	IMEIUnlock(ctx context.Context, id types.ID) error
	Logs(ctx context.Context, id types.ID) (*SIMLogsResult, error)
	GetNetworkOperator(ctx context.Context, id types.ID) ([]*SIMNetworkOperatorConfig, error)
	SetNetworkOperator(ctx context.Context, id types.ID, configs []*SIMNetworkOperatorConfig) error
	MonitorSIM(ctx context.Context, id types.ID, condition *MonitorCondition) (*LinkActivity, error)
	Status(ctx context.Context, id types.ID) (*SIMInfo, error)
}

/*************************************************
* SimpleMonitorAPI
*************************************************/

// SimpleMonitorAPI is interface for operate SimpleMonitor resource
type SimpleMonitorAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*SimpleMonitorFindResult, error)
	Create(ctx context.Context, param *SimpleMonitorCreateRequest) (*SimpleMonitor, error)
	Read(ctx context.Context, id types.ID) (*SimpleMonitor, error)
	Update(ctx context.Context, id types.ID, param *SimpleMonitorUpdateRequest) (*SimpleMonitor, error)
	UpdateSettings(ctx context.Context, id types.ID, param *SimpleMonitorUpdateSettingsRequest) (*SimpleMonitor, error)
	Delete(ctx context.Context, id types.ID) error
	MonitorResponseTime(ctx context.Context, id types.ID, condition *MonitorCondition) (*ResponseTimeSecActivity, error)
	HealthStatus(ctx context.Context, id types.ID) (*SimpleMonitorHealthStatus, error)
}

/*************************************************
* SSHKeyAPI
*************************************************/

// SSHKeyAPI is interface for operate SSHKey resource
type SSHKeyAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*SSHKeyFindResult, error)
	Create(ctx context.Context, param *SSHKeyCreateRequest) (*SSHKey, error)
	Generate(ctx context.Context, param *SSHKeyGenerateRequest) (*SSHKeyGenerated, error)
	Read(ctx context.Context, id types.ID) (*SSHKey, error)
	Update(ctx context.Context, id types.ID, param *SSHKeyUpdateRequest) (*SSHKey, error)
	Delete(ctx context.Context, id types.ID) error
}

/*************************************************
* SubnetAPI
*************************************************/

// SubnetAPI is interface for operate Subnet resource
type SubnetAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*SubnetFindResult, error)
	Read(ctx context.Context, zone string, id types.ID) (*Subnet, error)
}

/*************************************************
* SwitchAPI
*************************************************/

// SwitchAPI is interface for operate Switch resource
type SwitchAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*SwitchFindResult, error)
	Create(ctx context.Context, zone string, param *SwitchCreateRequest) (*Switch, error)
	Read(ctx context.Context, zone string, id types.ID) (*Switch, error)
	Update(ctx context.Context, zone string, id types.ID, param *SwitchUpdateRequest) (*Switch, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	ConnectToBridge(ctx context.Context, zone string, id types.ID, bridgeID types.ID) error
	DisconnectFromBridge(ctx context.Context, zone string, id types.ID) error
	GetServers(ctx context.Context, zone string, id types.ID) (*SwitchGetServersResult, error)
}

/*************************************************
* VPCRouterAPI
*************************************************/

// VPCRouterAPI is interface for operate VPCRouter resource
type VPCRouterAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) (*VPCRouterFindResult, error)
	Create(ctx context.Context, zone string, param *VPCRouterCreateRequest) (*VPCRouter, error)
	Read(ctx context.Context, zone string, id types.ID) (*VPCRouter, error)
	Update(ctx context.Context, zone string, id types.ID, param *VPCRouterUpdateRequest) (*VPCRouter, error)
	UpdateSettings(ctx context.Context, zone string, id types.ID, param *VPCRouterUpdateSettingsRequest) (*VPCRouter, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Config(ctx context.Context, zone string, id types.ID) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	ConnectToSwitch(ctx context.Context, zone string, id types.ID, nicIndex int, switchID types.ID) error
	DisconnectFromSwitch(ctx context.Context, zone string, id types.ID, nicIndex int) error
	MonitorInterface(ctx context.Context, zone string, id types.ID, index int, condition *MonitorCondition) (*InterfaceActivity, error)
	Status(ctx context.Context, zone string, id types.ID) (*VPCRouterStatus, error)
}

/*************************************************
* WebAccelAPI
*************************************************/

// WebAccelAPI is interface for operate WebAccel resource
type WebAccelAPI interface {
	List(ctx context.Context) (*WebAccelListResult, error)
	Read(ctx context.Context, id types.ID) (*WebAccel, error)
	ReadCertificate(ctx context.Context, id types.ID) (*WebAccelCerts, error)
	CreateCertificate(ctx context.Context, id types.ID, param *WebAccelCertRequest) (*WebAccelCerts, error)
	UpdateCertificate(ctx context.Context, id types.ID, param *WebAccelCertRequest) (*WebAccelCerts, error)
	DeleteCertificate(ctx context.Context, id types.ID) error
	DeleteAllCache(ctx context.Context, param *WebAccelDeleteAllCacheRequest) error
	DeleteCache(ctx context.Context, param *WebAccelDeleteCacheRequest) ([]*WebAccelDeleteCacheResult, error)
}

/*************************************************
* ZoneAPI
*************************************************/

// ZoneAPI is interface for operate Zone resource
type ZoneAPI interface {
	Find(ctx context.Context, conditions *FindCondition) (*ZoneFindResult, error)
	Read(ctx context.Context, id types.ID) (*Zone, error)
}
