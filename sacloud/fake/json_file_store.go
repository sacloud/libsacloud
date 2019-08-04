package fake

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

const defaultJSONFilePath = "libsacloud-fake-store.json"

// JSONFileStore .
type JSONFileStore struct {
	Path       string
	Ctx        context.Context
	NoInitData bool

	mu    sync.Mutex
	cache map[string]map[string]interface{}
}

// NewJSONFileStore .
func NewJSONFileStore(path string) *JSONFileStore {
	return &JSONFileStore{
		Path:  path,
		cache: make(map[string]map[string]interface{}),
	}
}

// Init .
func (s *JSONFileStore) Init() error {
	if s.Ctx == nil {
		s.Ctx = context.Background()
	}
	if s.Path == "" {
		s.Path = defaultJSONFilePath
	}
	if stat, err := os.Stat(s.Path); err == nil {
		if stat.IsDir() {
			return fmt.Errorf("path %q is directory", s.Path)
		}
	} else {
		if _, err := os.Create(s.Path); err != nil {
			return err
		}
	}

	if err := s.load(); err != nil {
		return err
	}
	s.startWatcher()
	return nil
}

func (s *JSONFileStore) startWatcher() {
	ctx := s.Ctx
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	log.Printf("file watch start: %q", s.Path)

	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				switch {
				case event.Op&fsnotify.Write == fsnotify.Write,
					event.Op&fsnotify.Create == fsnotify.Create,
					event.Op&fsnotify.Rename == fsnotify.Rename:

					if err := s.load(); err != nil {
						log.Printf("reloading %q is failed: %s\n", s.Path, err)
					}

					if event.Op&fsnotify.Rename == fsnotify.Rename {
						watcher.Add(s.Path)
					}
					log.Printf("reloaded: %q\n", s.Path)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				panic(err)
			case <-ctx.Done():
				return
			}
		}
	}()
	watcher.Add(s.Path)
}

// NeedInitData .
func (s *JSONFileStore) NeedInitData() bool {
	return !s.NoInitData && len(s.cache[sacloud.APIDefaultZone]) < 2
}

// Put .
func (s *JSONFileStore) Put(resourceKey, zone string, id types.ID, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values == nil {
		values = map[string]interface{}{}
	}
	values[id.String()] = value
	s.cache[s.key(resourceKey, zone)] = values

	s.store()
}

// Get .
func (s *JSONFileStore) Get(resourceKey, zone string, id types.ID) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values == nil {
		return nil
	}
	return values[id.String()]
}

// List .
func (s *JSONFileStore) List(resourceKey, zone string) []interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	var ret []interface{}
	for _, v := range values {
		ret = append(ret, v)
	}
	return ret
}

// Delete .
func (s *JSONFileStore) Delete(resourceKey, zone string, id types.ID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values != nil {
		delete(values, id.String())
	}
	s.store()
}

var jsonResourceTypeMap = map[string]func() interface{}{
	ResourceArchive:         func() interface{} { return &sacloud.Archive{} },
	ResourceAuthStatus:      func() interface{} { return &sacloud.AuthStatus{} },
	ResourceAutoBackup:      func() interface{} { return &sacloud.AutoBackup{} },
	ResourceBill:            func() interface{} { return &sacloud.Bill{} },
	ResourceBridge:          func() interface{} { return &sacloud.Bridge{} },
	ResourceCDROM:           func() interface{} { return &sacloud.CDROM{} },
	ResourceCoupon:          func() interface{} { return &sacloud.Coupon{} },
	ResourceDatabase:        func() interface{} { return &sacloud.Database{} },
	ResourceDisk:            func() interface{} { return &sacloud.Disk{} },
	ResourceDiskPlan:        func() interface{} { return &sacloud.DiskPlan{} },
	ResourceDNS:             func() interface{} { return &sacloud.DNS{} },
	ResourceGSLB:            func() interface{} { return &sacloud.GSLB{} },
	ResourceIcon:            func() interface{} { return &sacloud.Icon{} },
	ResourceInterface:       func() interface{} { return &sacloud.Interface{} },
	ResourceInternet:        func() interface{} { return &sacloud.Internet{} },
	ResourceInternetPlan:    func() interface{} { return &sacloud.InternetPlan{} },
	ResourceIPAddress:       func() interface{} { return &sacloud.IPAddress{} },
	ResourceIPv6Net:         func() interface{} { return &sacloud.IPv6Net{} },
	ResourceIPv6Addr:        func() interface{} { return &sacloud.IPv6Addr{} },
	ResourceLicense:         func() interface{} { return &sacloud.License{} },
	ResourceLicenseInfo:     func() interface{} { return &sacloud.LicenseInfo{} },
	ResourceLoadBalancer:    func() interface{} { return &sacloud.LoadBalancer{} },
	ResourceMobileGateway:   func() interface{} { return &sacloud.MobileGateway{} },
	ResourceNFS:             func() interface{} { return &sacloud.NFS{} },
	ResourceNote:            func() interface{} { return &sacloud.Note{} },
	ResourcePacketFilter:    func() interface{} { return &sacloud.PacketFilter{} },
	ResourcePrivateHost:     func() interface{} { return &sacloud.PrivateHost{} },
	ResourcePrivateHostPlan: func() interface{} { return &sacloud.PrivateHostPlan{} },
	ResourceProxyLB:         func() interface{} { return &sacloud.ProxyLB{} },
	ResourceRegion:          func() interface{} { return &sacloud.Region{} },
	ResourceServer:          func() interface{} { return &sacloud.Server{} },
	ResourceServerPlan:      func() interface{} { return &sacloud.ServerPlan{} },
	ResourceServiceClass:    func() interface{} { return &sacloud.ServiceClass{} },
	ResourceSIM:             func() interface{} { return &sacloud.SIM{} },
	ResourceSimpleMonitor:   func() interface{} { return &sacloud.SimpleMonitor{} },
	ResourceSSHKey:          func() interface{} { return &sacloud.SSHKey{} },
	ResourceSwitch:          func() interface{} { return &sacloud.Switch{} },
	ResourceVPCRouter:       func() interface{} { return &sacloud.VPCRouter{} },
	ResourceWebAccel:        func() interface{} { return &sacloud.WebAccel{} },
	ResourceZone:            func() interface{} { return &sacloud.Zone{} },

	valuePoolResourceKey:         func() interface{} { return &valuePool{} },
	"BillDetails":                func() interface{} { return &[]*sacloud.BillDetail{} },
	"MobileGatewayDNS":           func() interface{} { return &sacloud.MobileGatewayDNSSetting{} },
	"MobileGatewaySIMRoutes":     func() interface{} { return &[]*sacloud.MobileGatewaySIMRoute{} },
	"MobileGatewaySIMs":          func() interface{} { return &[]*sacloud.MobileGatewaySIMInfo{} },
	"MobileGatewayTrafficConfig": func() interface{} { return &sacloud.MobileGatewayTrafficControl{} },
	"ProxyLBStatus":              func() interface{} { return &sacloud.ProxyLBHealth{} },
	"SIMNetworkOperator":         func() interface{} { return &[]*sacloud.SIMNetworkOperatorConfig{} },
}

func (s *JSONFileStore) unmarshalResource(resourceKey string, data []byte) (interface{}, error) {
	f, ok := jsonResourceTypeMap[resourceKey]
	if !ok {
		panic(fmt.Errorf("type %q is not registered", resourceKey))
	}
	v := f()
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *JSONFileStore) store() error {
	data, err := json.MarshalIndent(s.cache, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(s.Path, data, 0600)
}

func (s *JSONFileStore) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := ioutil.ReadFile(s.Path)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}

	var cache = make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &cache); err != nil {
		return err
	}

	var loaded = make(map[string]map[string]interface{})
	for cacheKey, values := range cache {
		resourceKey, _ := s.parseKey(cacheKey)

		var dest = make(map[string]interface{})
		for id, v := range values {
			data, err := json.Marshal(v)
			if err != nil {
				return err
			}
			cv, err := s.unmarshalResource(resourceKey, data)
			if err != nil {
				return err
			}
			dest[id] = cv
		}
		loaded[cacheKey] = dest
	}
	s.cache = loaded
	return nil
}

func (s *JSONFileStore) key(resourceKey, zone string) string {
	return fmt.Sprintf("%s/%s", resourceKey, zone)
}

func (s *JSONFileStore) parseKey(k string) (string, string) {
	ss := strings.Split(k, "/")
	if len(ss) == 2 {
		return ss[0], ss[1]
	}
	return "", ""
}

func (s *JSONFileStore) values(resourceKey, zone string) map[string]interface{} {
	return s.cache[s.key(resourceKey, zone)]
}
