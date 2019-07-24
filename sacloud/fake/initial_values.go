package fake

import (
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var (
	accountID    = types.ID(123456789012)
	accountName  = "fakeアカウント"
	accountCode  = "fake"
	accountClass = "member"
	memberCode   = "fake-member"
	memberClass  = "member"
)

var authStatus = &sacloud.AuthStatus{
	AccountID:          accountID,
	AccountName:        accountName,
	AccountCode:        accountCode,
	AccountClass:       accountClass,
	MemberCode:         memberCode,
	MemberClass:        memberClass,
	AuthClass:          types.AuthClasses.Account,
	AuthMethod:         types.AuthMethods.APIKey,
	IsAPIKey:           true,
	ExternalPermission: types.ExternalPermission("bill+eventlog+cdn"),
	OperationPenalty:   types.OperationPenalties.None,
	Permission:         types.Permissions.Create,
}

var zones = []string{"tk1a", "is1a", "is1b", "tk1v"}

var zoneIDs = map[string]types.ID{
	"tk1a": types.ID(21001),
	"is1a": types.ID(31001),
	"is1b": types.ID(31002),
	"tk1v": types.ID(29001),
}

var sharedSegmentSwitch = &sacloud.Switch{
	ID:             pool.generateID(),
	Name:           "スイッチ",
	Scope:          types.Scopes.Shared,
	Description:    "共有セグメント用スイッチ",
	NetworkMaskLen: pool.sharedNetMaskLen,
	DefaultRoute:   pool.sharedDefaultGateway.String(),
}

func init() {
	initArchives()
	initBills()
	initCoupons()
	initCDROMs()
	initNotes()
	initSwitch()
	initZones()
	initRegions()
	initPrivateHostPlan()
	initDiskPlan()
	initLicensePlan()
	initInternetPlan()
	initServerPlan()
	initServiceClass()
}

func initArchives() {
	archives := []*sacloud.Archive{
		{
			ID:                   pool.generateID(),
			Name:                 "CentOS 7.6 (1810) 64bit",
			Tags:                 []string{"@size-extendable", "arch-64bit", "current-stable", "distro-centos", "distro-ver-7.6", "os-linux"},
			DisplayOrder:         1,
			Scope:                types.Scopes.Shared,
			Availability:         types.Availabilities.Available,
			SizeMB:               20480,
			DiskPlanID:           types.ID(2),
			DiskPlanName:         "標準プラン",
			DiskPlanStorageClass: "iscsi9999",
		},
		{
			ID:                   pool.generateID(),
			Name:                 "Ubuntu Server 18.04.2 LTS 64bit",
			Tags:                 []string{"@size-extendable", "arch-64bit", "current-stable", "distro-ubuntu", "distro-ver-18.04.2", "os-linux"},
			DisplayOrder:         2,
			Scope:                types.Scopes.Shared,
			Availability:         types.Availabilities.Available,
			SizeMB:               20480,
			DiskPlanID:           types.ID(2),
			DiskPlanName:         "標準プラン",
			DiskPlanStorageClass: "iscsi9999",
		},
	}
	for _, zone := range zones {
		for _, archive := range archives {
			s.setArchive(zone, archive)
		}
	}
}

func initBills() {
	bills := []*sacloud.Bill{
		{
			ID:             pool.generateID(),
			Amount:         1080,
			Date:           time.Now(),
			MemberID:       "dummy00000",
			Paid:           false,
			PayLimit:       time.Now().AddDate(0, 1, 0),
			PaymentClassID: 999,
		},
	}
	for _, bill := range bills {
		s.setBill(sacloud.APIDefaultZone, bill)
		initBillDetails(bill.ID)
	}
}

func initBillDetails(billID types.ID) {
	details := []*sacloud.BillDetail{
		{
			ID:             pool.generateID(),
			Amount:         108,
			Description:    "description",
			ServiceClassID: 999,
			Usage:          100,
			Zone:           "tk1a",
			ContractEndAt:  time.Now(),
		},
	}
	s.setWithID(ResourceBill+"Details", sacloud.APIDefaultZone, details, billID)
}

func initCDROMs() {
	cdroms := []*sacloud.CDROM{
		{
			ID:           pool.generateID(),
			Name:         "dummy",
			Description:  "dummy",
			DisplayOrder: 999,
			Tags:         types.Tags{"current-stable", "os-linux"},
			Availability: types.Availabilities.Available,
			Scope:        types.Scopes.Shared,
			CreatedAt:    time.Now(),
			ModifiedAt:   time.Now(),
		},
	}
	for _, zone := range zones {
		for _, c := range cdroms {
			s.setCDROM(zone, c)
		}
	}
}

func initCoupons() {
	coupons := []*sacloud.Coupon{
		{
			ID:             pool.generateID(),
			MemberID:       "dummy00000",
			ContractID:     pool.generateID(),
			ServiceClassID: 999,
			Discount:       20000,
			AppliedAt:      time.Now().AddDate(0, -1, 0),
			UntilAt:        time.Now().AddDate(0, 1, 0),
		},
	}
	for _, c := range coupons {
		s.setCoupon(sacloud.APIDefaultZone, c)
	}
}

func initNotes() {
	notes := []*sacloud.Note{
		{
			ID:           1,
			Name:         "fake",
			Availability: types.Availabilities.Available,
			Scope:        types.Scopes.Shared,
			Class:        "shell",
		},
	}
	for _, note := range notes {
		s.setNote(sacloud.APIDefaultZone, note)
	}
}

func initSwitch() {
	for _, zone := range zones {
		s.setSwitch(zone, sharedSegmentSwitch)
	}
}

func initZones() {
	// zones
	s.setZone(sacloud.APIDefaultZone, &sacloud.Zone{
		ID:           21001,
		Name:         "tk1a",
		Description:  "東京第1ゾーン",
		DisplayOrder: 1,
	})
	s.setZone(sacloud.APIDefaultZone, &sacloud.Zone{
		ID:           31001,
		Name:         "is1a",
		Description:  "石狩第1ゾーン",
		DisplayOrder: 2,
	})
	s.setZone(sacloud.APIDefaultZone, &sacloud.Zone{
		ID:           31002,
		Name:         "is1b",
		Description:  "石狩第2ゾーン",
		DisplayOrder: 3,
	})
	s.setZone(sacloud.APIDefaultZone, &sacloud.Zone{
		ID:           29001,
		Name:         "tk1v",
		Description:  "Sandbox",
		DisplayOrder: 4,
		IsDummy:      true,
	})
}

func initRegions() {
	s.setRegion(sacloud.APIDefaultZone, &sacloud.Region{
		ID:          210,
		Name:        "東京",
		Description: "東京",
		NameServers: []string{
			"210.188.224.10",
			"210.188.224.11",
		},
	})
	s.setRegion(sacloud.APIDefaultZone, &sacloud.Region{
		ID:          290,
		Name:        "Sandbox",
		Description: "Sandbox",
		NameServers: []string{
			"133.242.0.3",
			"133.242.0.4",
		},
	})
	s.setRegion(sacloud.APIDefaultZone, &sacloud.Region{
		ID:          310,
		Name:        "石狩",
		Description: "石狩",
		NameServers: []string{
			"133.242.0.3",
			"133.242.0.4",
		},
	})
}

func initPrivateHostPlan() {
	s.setPrivateHostPlan("is1b", &sacloud.PrivateHostPlan{
		ID:           112900526366,
		Name:         "200Core 224GB 標準",
		Class:        "dynamic",
		CPU:          200,
		MemoryMB:     229376,
		Availability: types.Availabilities.Available,
	})
	s.setPrivateHostPlan("tk1a", &sacloud.PrivateHostPlan{
		ID:           112900526366,
		Name:         "200Core 224GB 標準",
		Class:        "dynamic",
		CPU:          200,
		MemoryMB:     229376,
		Availability: types.Availabilities.Available,
	})
}

func initServerPlan() {
	plans := []*sacloud.ServerPlan{
		{
			ID:           pool.generateID(),
			Name:         "プラン/1Core-1GB",
			CPU:          1,
			MemoryMB:     1024,
			Commitment:   types.Commitments.Standard,
			Generation:   100,
			Availability: types.Availabilities.Available,
		},
		{
			ID:           pool.generateID(),
			Name:         "プラン/2Core-4GB",
			CPU:          2,
			MemoryMB:     4 * 1024,
			Commitment:   types.Commitments.Standard,
			Generation:   100,
			Availability: types.Availabilities.Available,
		},
		// TODO add more plans
	}

	for _, zone := range zones {
		for _, plan := range plans {
			s.setServerPlan(zone, plan)
		}
	}
}

func initInternetPlan() {
	bandWidthList := []int{100, 250500, 1000, 1500, 2000, 2500, 3000, 5000}

	var plans []*sacloud.InternetPlan

	for _, bw := range bandWidthList {
		plans = append(plans, &sacloud.InternetPlan{
			ID:            types.ID(bw),
			BandWidthMbps: bw,
			Name:          fmt.Sprintf("%dMbps共有", bw),
			Availability:  types.Availabilities.Available,
		})
	}

	for _, zone := range zones {
		for _, plan := range plans {
			s.setInternetPlan(zone, plan)
		}
	}
}

func initDiskPlan() {
	plans := []*sacloud.DiskPlan{
		{
			ID:           2,
			Name:         "HDDプラン",
			Availability: types.Availabilities.Available,
			StorageClass: "iscsi1204",
			Size: []*sacloud.DiskPlanSizeInfo{
				{
					Availability:  types.Availabilities.Available,
					DisplaySize:   20,
					DisplaySuffix: "GB",
					SizeMB:        20 * 1024,
				},
				{
					Availability:  types.Availabilities.Available,
					DisplaySize:   40,
					DisplaySuffix: "GB",
					SizeMB:        40 * 1024,
				},
			},
		},
		{
			ID:           4,
			Name:         "SSDプラン",
			Availability: types.Availabilities.Available,
			StorageClass: "iscsi1204",
			Size: []*sacloud.DiskPlanSizeInfo{
				{
					Availability:  types.Availabilities.Available,
					DisplaySize:   20,
					DisplaySuffix: "GB",
					SizeMB:        20 * 1024,
				},
				{
					Availability:  types.Availabilities.Available,
					DisplaySize:   40,
					DisplaySuffix: "GB",
					SizeMB:        40 * 1024,
				},
			},
		},
		// TODO add more size-info
	}

	for _, zone := range zones {
		for _, plan := range plans {
			s.setDiskPlan(zone, plan)
		}
	}
}

func initLicensePlan() {
	plans := []*sacloud.LicenseInfo{
		{
			ID:         types.ID(10001),
			Name:       "Windows RDS SAL",
			TermsOfUse: "1ライセンスにつき、1人のユーザが利用できます。",
		},
	}

	for _, zone := range zones {
		for _, plan := range plans {
			s.setLicenseInfo(zone, plan)
		}
	}
}

func initServiceClass() {
	classes := []*sacloud.ServiceClass{
		{
			ID:               types.ID(50050),
			ServiceClassName: "plan/1",
			ServiceClassPath: "cloud/plan/1",
			DisplayName:      "プラン1(ディスクなし)",
			IsPublic:         true,
			Price: &sacloud.Price{
				Base:    0,
				Daily:   108,
				Hourly:  10,
				Monthly: 2139,
			},
		},
		{
			ID:               types.ID(50051),
			ServiceClassName: "plan/2",
			ServiceClassPath: "cloud/plan/2",
			DisplayName:      "プラン2(ディスクなし)",
			IsPublic:         true,
			Price: &sacloud.Price{
				Base:    0,
				Daily:   172,
				Hourly:  17,
				Monthly: 3425,
			},
		},
	}

	for _, zone := range zones {
		for _, class := range classes {
			class.Price.Zone = zone
			s.setServiceClass(zone, class)
		}
	}
}
