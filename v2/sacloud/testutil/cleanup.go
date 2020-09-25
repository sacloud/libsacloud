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

package testutil

import (
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func ComposeCleanupFuncs(funcs ...func(context.Context, sacloud.APICaller) error) func(context.Context, sacloud.APICaller) error {
	return func(ctx context.Context, caller sacloud.APICaller) error {
		for _, f := range funcs {
			if err := f(ctx, caller); err != nil {
				return err
			}
		}
		return nil
	}
}

func ComposeCleanupResourceFunc(prefix string, targets ...CleanupTarget) func(context.Context, sacloud.APICaller) error {
	var funcs []func(context.Context, sacloud.APICaller) error
	for _, t := range targets {
		funcs = append(funcs, func(ctx context.Context, caller sacloud.APICaller) error {
			return CleanupResource(ctx, caller, prefix, t)
		})
	}
	return ComposeCleanupFuncs(funcs...)
}

func CleanupResource(ctx context.Context, caller sacloud.APICaller, prefix string, target CleanupTarget) error {
	if !IsAccTest() {
		return nil
	}
	if prefix == "" {
		prefix = TestResourcePrefix
	}
	cleanupFindCondition = &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key("Name"): search.PartialMatch(prefix),
		},
	}

	searched, err := target.finder(ctx, caller)
	if err != nil {
		return err
	}

	var errs *multierror.Error
	var wg sync.WaitGroup
	for i := range searched {
		wg.Add(1)
		go func(target *cleanupTarget) {
			defer wg.Done()
			if target.prepareFunc != nil {
				if err := target.prepareFunc(ctx); err != nil {
					multierror.Append(errs, err) // nolint
					return
				}
			}
			if target.deleteFunc != nil {
				if err := target.deleteFunc(ctx); err != nil {
					multierror.Append(errs, err) // nolint
					return
				}
			}
		}(searched[i])
	}
	wg.Wait()

	return errs.ErrorOrNil()
}

type CleanupTarget struct {
	finder cleanupTargetFindFunc
}

// CleanupTargets クリーンアップ対象のリソース。CleanupResourceに渡す
var CleanupTargets = struct {
	Archive           CleanupTarget
	AutoBackup        CleanupTarget
	Bridge            CleanupTarget
	ContainerRegistry CleanupTarget
	CDROM             CleanupTarget
	Database          CleanupTarget
	Disk              CleanupTarget
	DNS               CleanupTarget
	GSLB              CleanupTarget
	Icon              CleanupTarget
	Internet          CleanupTarget
	License           CleanupTarget
	LoadBalancer      CleanupTarget
	MobileGateway     CleanupTarget
	NFS               CleanupTarget
	Note              CleanupTarget
	PacketFilter      CleanupTarget
	PrivateHost       CleanupTarget
	ProxyLB           CleanupTarget
	Server            CleanupTarget
	SIM               CleanupTarget
	SimpleMonitor     CleanupTarget
	SSHKey            CleanupTarget
	Switch            CleanupTarget
	VPCRouter         CleanupTarget
}{
	Archive:           CleanupTarget{finder: findArchive},
	AutoBackup:        CleanupTarget{finder: findAutoBackup},
	Bridge:            CleanupTarget{finder: findBridge},
	ContainerRegistry: CleanupTarget{finder: findContainerRegistry},
	CDROM:             CleanupTarget{finder: findCDROM},
	Database:          CleanupTarget{finder: findDatabase},
	Disk:              CleanupTarget{finder: findDisk},
	DNS:               CleanupTarget{finder: findDNS},
	GSLB:              CleanupTarget{finder: findGSLB},
	Icon:              CleanupTarget{finder: findIcon},
	Internet:          CleanupTarget{finder: findInternet},
	License:           CleanupTarget{finder: findLicense},
	LoadBalancer:      CleanupTarget{finder: findLoadBalancer},
	MobileGateway:     CleanupTarget{finder: findMobileGateway},
	NFS:               CleanupTarget{finder: findNFS},
	Note:              CleanupTarget{finder: findNote},
	PacketFilter:      CleanupTarget{finder: findPacketFilter},
	PrivateHost:       CleanupTarget{finder: findPrivateHost},
	ProxyLB:           CleanupTarget{finder: findProxyLB},
	Server:            CleanupTarget{finder: findServer},
	SIM:               CleanupTarget{finder: findSIM},
	SimpleMonitor:     CleanupTarget{finder: findSimpleMonitor},
	SSHKey:            CleanupTarget{finder: findSSHKey},
	Switch:            CleanupTarget{finder: findSwitch},
	VPCRouter:         CleanupTarget{finder: findVPCRouter},
}

// CleanupTestResources 指定プレフィックスを持つリソースの削除を行う
//
// TESTACC環境変数が設定されている場合のみ実施される
func CleanupTestResources(ctx context.Context, caller sacloud.APICaller) error {
	if !IsAccTest() {
		return nil
	}
	cleanupFindCondition = &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key("Name"): search.PartialMatch(TestResourcePrefix),
		},
	}
	var errs *multierror.Error

	// cleanup: primary group
	doCleanup(ctx, correctCleanupTargets(ctx, caller, cleanupPrimaryGroup, errs), errs)
	// cleanup: secondary group
	doCleanup(ctx, correctCleanupTargets(ctx, caller, cleanupSecondaryGroup, errs), errs)

	return errs.ErrorOrNil()
}

func doCleanup(ctx context.Context, targets []*cleanupTarget, errs *multierror.Error) {
	var wg sync.WaitGroup
	for i := range targets {
		wg.Add(1)
		go func(target *cleanupTarget) {
			defer wg.Done()
			if target.prepareFunc != nil {
				if err := target.prepareFunc(ctx); err != nil {
					multierror.Append(errs, err) // nolint
					return
				}
			}
			if target.deleteFunc != nil {
				if err := target.deleteFunc(ctx); err != nil {
					multierror.Append(errs, err) // nolint
					return
				}
			}
		}(targets[i])
	}
	wg.Wait()
}

func correctCleanupTargets(ctx context.Context, caller sacloud.APICaller, finders []cleanupTargetFindFunc, errs *multierror.Error) []*cleanupTarget {
	var targets []*cleanupTarget
	var wg sync.WaitGroup
	for i := range finders {
		wg.Add(1)
		go func(finder cleanupTargetFindFunc) {
			defer wg.Done()

			res, err := finder(ctx, caller)
			if err != nil {
				multierror.Append(errs, err) // nolint
				return
			}
			targets = append(targets, res...)
		}(finders[i])
	}
	wg.Wait()
	return targets
}

type cleanupTargetFindFunc func(context.Context, sacloud.APICaller) ([]*cleanupTarget, error)

var cleanupPrimaryGroup = []cleanupTargetFindFunc{
	findArchive,
	findAutoBackup,
	findContainerRegistry,
	findDatabase,
	findDNS,
	findGSLB,
	findIcon,
	findLicense,
	findLoadBalancer,
	findNFS,
	findNote,
	findPacketFilter,
	findProxyLB,
	findServer,
	findSimpleMonitor,
	findSSHKey,
	findVPCRouter,
	findMobileGateway,
}

var cleanupSecondaryGroup = []cleanupTargetFindFunc{
	findBridge,
	findCDROM,
	findDisk,
	findInternet,
	findSwitch,
	findPrivateHost,
	findSIM,
}

var cleanupFindCondition = &sacloud.FindCondition{
	Filter: search.Filter{
		search.Key("Name"): search.PartialMatch(TestResourcePrefix),
	},
}

type cleanupTarget struct {
	resource    interface{}
	prepareFunc func(context.Context) error
	deleteFunc  func(context.Context) error
}

func findBridge(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewBridgeOp(caller)
	searched, err := op.Find(ctx, sacloud.APIDefaultZone, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	var res []*cleanupTarget
	for i := range searched.Bridges {
		v := searched.Bridges[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, sacloud.APIDefaultZone, v.ID)
			},
		})
	}
	return res, nil
}

func findContainerRegistry(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewContainerRegistryOp(caller)
	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	var res []*cleanupTarget
	for i := range searched.ContainerRegistries {
		v := searched.ContainerRegistries[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findCDROM(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewCDROMOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.CDROMs {
			v := searched.CDROMs[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findInternet(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewInternetOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.Internet {
			v := searched.Internet[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findSwitch(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewSwitchOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.Switches {
			v := searched.Switches[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findPrivateHost(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewPrivateHostOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.PrivateHosts {
			v := searched.PrivateHosts[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findArchive(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewArchiveOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.Archives {
			v := searched.Archives[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findAutoBackup(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewAutoBackupOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.AutoBackups {
			v := searched.AutoBackups[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findDatabase(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewDatabaseOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.Databases {
			v := searched.Databases[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, zone, v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, zone, v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findDisk(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewDiskOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.Disks {
			v := searched.Disks[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findDNS(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewDNSOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.DNS {
		v := searched.DNS[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findGSLB(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewGSLBOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.GSLBs {
		v := searched.GSLBs[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findIcon(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewIconOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.Icons {
		v := searched.Icons[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findLicense(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewLicenseOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.Licenses {
		v := searched.Licenses[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findLoadBalancer(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewLoadBalancerOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.LoadBalancers {
			v := searched.LoadBalancers[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, zone, v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, zone, v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findMobileGateway(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewMobileGatewayOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.MobileGateways {
			v := searched.MobileGateways[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					// delete sim routes
					if err := op.SetSIMRoutes(ctx, zone, v.ID, []*sacloud.MobileGatewaySIMRouteParam{}); err != nil {
						return err
					}

					// delete SIMs
					sims, err := op.ListSIM(ctx, zone, v.ID)
					if err != nil {
						return err
					}
					for _, sim := range sims {
						if err := op.DeleteSIM(ctx, zone, v.ID, types.StringID(sim.ResourceID)); err != nil {
							return err
						}
					}

					if err := op.Shutdown(ctx, zone, v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err = sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, zone, v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findNFS(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewNFSOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.NFS {
			v := searched.NFS[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, zone, v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, zone, v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findNote(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewNoteOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.Notes {
		v := searched.Notes[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findPacketFilter(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewPacketFilterOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.PacketFilters {
			v := searched.PacketFilters[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findProxyLB(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewProxyLBOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.ProxyLBs {
		v := searched.ProxyLBs[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findServer(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewServerOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.Servers {
			v := searched.Servers[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, zone, v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, zone, v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}

func findSIM(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewSIMOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.SIMs {
		v := searched.SIMs[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findSimpleMonitor(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewSimpleMonitorOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.SimpleMonitors {
		v := searched.SimpleMonitors[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findSSHKey(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewSSHKeyOp(caller)
	var res []*cleanupTarget

	searched, err := op.Find(ctx, cleanupFindCondition)
	if err != nil {
		return nil, err
	}
	for i := range searched.SSHKeys {
		v := searched.SSHKeys[i]
		res = append(res, &cleanupTarget{
			resource: v,
			deleteFunc: func(ctx context.Context) error {
				return op.Delete(ctx, v.ID)
			},
		})
	}
	return res, nil
}

func findVPCRouter(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewVPCRouterOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		zone := types.ZoneNames[i]
		for j := range searched.VPCRouters {
			v := searched.VPCRouters[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, zone, v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, zone, v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, zone, v.ID)
				},
			})
		}
	}
	return res, nil
}
