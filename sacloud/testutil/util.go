package testutil

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/trace"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// TestResourcePrefix テスト時に作成するリソースの名称に付与するプレフィックス
//
// このプレフィックスを持つリソースは受入テスト実行後に削除される
const TestResourcePrefix = "libsacloud-test-"

// ResourceName テスト時に作成するリソースの名称
func ResourceName(name string) string {
	return fmt.Sprintf("%s%s", TestResourcePrefix, name)
}

var testZone string
var apiCaller *sacloud.Client
var httpTrace bool

var accTestOnce sync.Once
var accTestMu sync.Mutex

// SingletonAPICaller 環境変数からシングルトンAPICallerを作成する
func SingletonAPICaller() *sacloud.Client {

	accTestMu.Lock()
	defer accTestMu.Unlock()
	accTestOnce.Do(func() {
		if !IsAccTest() {
			sacloud.DefaultStatePollInterval = 100 * time.Millisecond
			fake.SwitchFactoryFuncToFake()
		}

		if IsEnableTrace() || IsEnableAPITrace() {
			trace.AddClientFactoryHooks()
		}

		if IsEnableTrace() || IsEnableHTTPTrace() {
			httpTrace = true
		}

		//環境変数にトークン/シークレットがある場合のみテスト実施
		accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
		accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

		if accessToken == "" || accessTokenSecret == "" {
			log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
			os.Exit(0) // exit normal
		}
		client := sacloud.NewClient(accessToken, accessTokenSecret)
		client.DefaultTimeoutDuration = 30 * time.Minute
		client.UserAgent = fmt.Sprintf("test-libsacloud/%s", libsacloud.Version)
		client.AcceptLanguage = "en-US,en;q=0.9"

		client.RetryMax = 20
		client.RetryInterval = 3 * time.Second
		client.HTTPClient = &http.Client{
			Transport: &sacloud.RateLimitRoundTripper{RateLimitPerSec: 1},
		}
		if httpTrace {
			client.HTTPClient.Transport = &sacloud.TracingRoundTripper{
				Transport: client.HTTPClient.Transport,
			}
		}

		apiCaller = client
	})
	return apiCaller
}

// TestZone SAKURACLOUD_ZONE環境変数からテスト対象のゾーンを取得 デフォルトはtk1v
func TestZone() string {
	testZone := os.Getenv("SAKURACLOUD_ZONE")
	if testZone == "" {
		testZone = "tk1v"
	}
	return testZone
}

// IsAccTest TESTACC環境変数が指定されているか
func IsAccTest() bool {
	return os.Getenv("TESTACC") != ""
}

// IsEnableTrace SAKURACLOUD_TRACE環境変数が指定されているか
func IsEnableTrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE") != ""
}

// IsEnableAPITrace SAKURACLOUD_TRACE_API環境変数が指定されているか
func IsEnableAPITrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE_API") != ""
}

// IsEnableHTTPTrace SAKURACLOUD_TRACE_HTTP環境変数が指定されているか
func IsEnableHTTPTrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE_HTTP") != ""
}

// CleanupTestResources 指定プレフィックスを持つリソースの削除を行う
//
// TESTACC環境変数が設定されている場合のみ実施される
func CleanupTestResources(ctx context.Context, caller sacloud.APICaller, namePrefix string) error {
	if !IsAccTest() {
		return nil
	}

	if namePrefix == "" {
		cleanupFindCondition = &sacloud.FindCondition{
			Filter: search.Filter{
				search.Key("Name"): search.PartialMatch(namePrefix),
			},
		}
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
					multierror.Append(errs, err)
					return
				}
			}
			if target.deleteFunc != nil {
				if err := target.deleteFunc(ctx); err != nil {
					multierror.Append(errs, err)
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
				multierror.Append(errs, err)
				return
			}
			for _, v := range res {
				targets = append(targets, v)
			}
		}(finders[i])
	}
	wg.Wait()
	return targets
}

type cleanupTargetFindFunc func(context.Context, sacloud.APICaller) ([]*cleanupTarget, error)

var cleanupPrimaryGroup = []cleanupTargetFindFunc{
	findArchive,
	findAutoBackup,
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

func findCDROM(ctx context.Context, caller sacloud.APICaller) ([]*cleanupTarget, error) {
	op := sacloud.NewCDROMOp(caller)
	var res []*cleanupTarget

	for i := range types.ZoneNames {
		searched, err := op.Find(ctx, types.ZoneNames[i], cleanupFindCondition)
		if err != nil {
			return nil, err
		}
		for j := range searched.CDROMs {
			v := searched.CDROMs[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.Internets {
			v := searched.Internets[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.Switches {
			v := searched.Switches[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.PrivateHosts {
			v := searched.PrivateHosts[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.Archives {
			v := searched.Archives[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.AutoBackups {
			v := searched.AutoBackups[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.Databases {
			v := searched.Databases[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, types.ZoneNames[i], v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, types.ZoneNames[i], v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.Disks {
			v := searched.Disks[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.LoadBalancers {
			v := searched.LoadBalancers[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, types.ZoneNames[i], v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, types.ZoneNames[i], v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.MobileGateways {
			v := searched.MobileGateways[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					// delete sim routes
					if err := op.SetSIMRoutes(ctx, types.ZoneNames[i], v.ID, []*sacloud.MobileGatewaySIMRouteParam{}); err != nil {
						return err
					}

					// delete SIMs
					sims, err := op.ListSIM(ctx, types.ZoneNames[i], v.ID)
					if err != nil {
						return err
					}
					for _, sim := range sims {
						if err := op.DeleteSIM(ctx, types.ZoneNames[i], v.ID, types.StringID(sim.ResourceID)); err != nil {
							return err
						}
					}

					if err := op.Shutdown(ctx, types.ZoneNames[i], v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err = sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, types.ZoneNames[i], v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.NFS {
			v := searched.NFS[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, types.ZoneNames[i], v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, types.ZoneNames[i], v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.PacketFilters {
			v := searched.PacketFilters[j]
			res = append(res, &cleanupTarget{
				resource: v,
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.Servers {
			v := searched.Servers[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, types.ZoneNames[i], v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, types.ZoneNames[i], v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
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
		for j := range searched.VPCRouters {
			v := searched.VPCRouters[j]
			res = append(res, &cleanupTarget{
				resource: v,
				prepareFunc: func(ctx context.Context) error {
					if err := op.Shutdown(ctx, types.ZoneNames[i], v.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return op.Read(ctx, types.ZoneNames[i], v.ID)
					}).WaitForState(ctx)
					return err
				},
				deleteFunc: func(ctx context.Context) error {
					return op.Delete(ctx, types.ZoneNames[i], v.ID)
				},
			})
		}
	}
	return res, nil
}
