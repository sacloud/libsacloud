// Copyright 2016-2021 The Libsacloud Authors
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

// generated by 'github.com/sacloud/libsacloud/v2/internal/tools/gen-api-fake-op'; DO NOT EDIT

package fake

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func TestResourceOps(t *testing.T) {

	if op, ok := NewArchiveOp().(sacloud.ArchiveAPI); !ok {
		t.Fatalf("%s is not sacloud.Archive", op)
	}

	if op, ok := NewAuthStatusOp().(sacloud.AuthStatusAPI); !ok {
		t.Fatalf("%s is not sacloud.AuthStatus", op)
	}

	if op, ok := NewAutoBackupOp().(sacloud.AutoBackupAPI); !ok {
		t.Fatalf("%s is not sacloud.AutoBackup", op)
	}

	if op, ok := NewBillOp().(sacloud.BillAPI); !ok {
		t.Fatalf("%s is not sacloud.Bill", op)
	}

	if op, ok := NewBridgeOp().(sacloud.BridgeAPI); !ok {
		t.Fatalf("%s is not sacloud.Bridge", op)
	}

	if op, ok := NewCDROMOp().(sacloud.CDROMAPI); !ok {
		t.Fatalf("%s is not sacloud.CDROM", op)
	}

	if op, ok := NewContainerRegistryOp().(sacloud.ContainerRegistryAPI); !ok {
		t.Fatalf("%s is not sacloud.ContainerRegistry", op)
	}

	if op, ok := NewCouponOp().(sacloud.CouponAPI); !ok {
		t.Fatalf("%s is not sacloud.Coupon", op)
	}

	if op, ok := NewDatabaseOp().(sacloud.DatabaseAPI); !ok {
		t.Fatalf("%s is not sacloud.Database", op)
	}

	if op, ok := NewDiskOp().(sacloud.DiskAPI); !ok {
		t.Fatalf("%s is not sacloud.Disk", op)
	}

	if op, ok := NewDiskPlanOp().(sacloud.DiskPlanAPI); !ok {
		t.Fatalf("%s is not sacloud.DiskPlan", op)
	}

	if op, ok := NewDNSOp().(sacloud.DNSAPI); !ok {
		t.Fatalf("%s is not sacloud.DNS", op)
	}

	if op, ok := NewESMEOp().(sacloud.ESMEAPI); !ok {
		t.Fatalf("%s is not sacloud.ESME", op)
	}

	if op, ok := NewGSLBOp().(sacloud.GSLBAPI); !ok {
		t.Fatalf("%s is not sacloud.GSLB", op)
	}

	if op, ok := NewIconOp().(sacloud.IconAPI); !ok {
		t.Fatalf("%s is not sacloud.Icon", op)
	}

	if op, ok := NewInterfaceOp().(sacloud.InterfaceAPI); !ok {
		t.Fatalf("%s is not sacloud.Interface", op)
	}

	if op, ok := NewInternetOp().(sacloud.InternetAPI); !ok {
		t.Fatalf("%s is not sacloud.Internet", op)
	}

	if op, ok := NewInternetPlanOp().(sacloud.InternetPlanAPI); !ok {
		t.Fatalf("%s is not sacloud.InternetPlan", op)
	}

	if op, ok := NewIPAddressOp().(sacloud.IPAddressAPI); !ok {
		t.Fatalf("%s is not sacloud.IPAddress", op)
	}

	if op, ok := NewIPv6NetOp().(sacloud.IPv6NetAPI); !ok {
		t.Fatalf("%s is not sacloud.IPv6Net", op)
	}

	if op, ok := NewIPv6AddrOp().(sacloud.IPv6AddrAPI); !ok {
		t.Fatalf("%s is not sacloud.IPv6Addr", op)
	}

	if op, ok := NewLicenseOp().(sacloud.LicenseAPI); !ok {
		t.Fatalf("%s is not sacloud.License", op)
	}

	if op, ok := NewLicenseInfoOp().(sacloud.LicenseInfoAPI); !ok {
		t.Fatalf("%s is not sacloud.LicenseInfo", op)
	}

	if op, ok := NewLoadBalancerOp().(sacloud.LoadBalancerAPI); !ok {
		t.Fatalf("%s is not sacloud.LoadBalancer", op)
	}

	if op, ok := NewLocalRouterOp().(sacloud.LocalRouterAPI); !ok {
		t.Fatalf("%s is not sacloud.LocalRouter", op)
	}

	if op, ok := NewMobileGatewayOp().(sacloud.MobileGatewayAPI); !ok {
		t.Fatalf("%s is not sacloud.MobileGateway", op)
	}

	if op, ok := NewNFSOp().(sacloud.NFSAPI); !ok {
		t.Fatalf("%s is not sacloud.NFS", op)
	}

	if op, ok := NewNoteOp().(sacloud.NoteAPI); !ok {
		t.Fatalf("%s is not sacloud.Note", op)
	}

	if op, ok := NewPacketFilterOp().(sacloud.PacketFilterAPI); !ok {
		t.Fatalf("%s is not sacloud.PacketFilter", op)
	}

	if op, ok := NewPrivateHostOp().(sacloud.PrivateHostAPI); !ok {
		t.Fatalf("%s is not sacloud.PrivateHost", op)
	}

	if op, ok := NewPrivateHostPlanOp().(sacloud.PrivateHostPlanAPI); !ok {
		t.Fatalf("%s is not sacloud.PrivateHostPlan", op)
	}

	if op, ok := NewProxyLBOp().(sacloud.ProxyLBAPI); !ok {
		t.Fatalf("%s is not sacloud.ProxyLB", op)
	}

	if op, ok := NewRegionOp().(sacloud.RegionAPI); !ok {
		t.Fatalf("%s is not sacloud.Region", op)
	}

	if op, ok := NewServerOp().(sacloud.ServerAPI); !ok {
		t.Fatalf("%s is not sacloud.Server", op)
	}

	if op, ok := NewServerPlanOp().(sacloud.ServerPlanAPI); !ok {
		t.Fatalf("%s is not sacloud.ServerPlan", op)
	}

	if op, ok := NewServiceClassOp().(sacloud.ServiceClassAPI); !ok {
		t.Fatalf("%s is not sacloud.ServiceClass", op)
	}

	if op, ok := NewSIMOp().(sacloud.SIMAPI); !ok {
		t.Fatalf("%s is not sacloud.SIM", op)
	}

	if op, ok := NewSimpleMonitorOp().(sacloud.SimpleMonitorAPI); !ok {
		t.Fatalf("%s is not sacloud.SimpleMonitor", op)
	}

	if op, ok := NewSSHKeyOp().(sacloud.SSHKeyAPI); !ok {
		t.Fatalf("%s is not sacloud.SSHKey", op)
	}

	if op, ok := NewSubnetOp().(sacloud.SubnetAPI); !ok {
		t.Fatalf("%s is not sacloud.Subnet", op)
	}

	if op, ok := NewSwitchOp().(sacloud.SwitchAPI); !ok {
		t.Fatalf("%s is not sacloud.Switch", op)
	}

	if op, ok := NewVPCRouterOp().(sacloud.VPCRouterAPI); !ok {
		t.Fatalf("%s is not sacloud.VPCRouter", op)
	}

	if op, ok := NewWebAccelOp().(sacloud.WebAccelAPI); !ok {
		t.Fatalf("%s is not sacloud.WebAccel", op)
	}

	if op, ok := NewZoneOp().(sacloud.ZoneAPI); !ok {
		t.Fatalf("%s is not sacloud.Zone", op)
	}

}
