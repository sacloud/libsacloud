package fake

import (
	"net"
	"sync"

	"github.com/sacloud/libsacloud/v2/pkg/cidr"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type valuePool struct {
	currentID            int64
	currentSharedIP      net.IP
	sharedNetMaskLen     int
	sharedDefaultGateway net.IP
	currentMACAddress    net.HardwareAddr
	mu                   sync.Mutex
	currentSubnets       map[int]*net.IPNet
}

var pool = &valuePool{
	currentID:            int64(100000000000),
	currentSharedIP:      net.IP{192, 0, 2, 2},
	sharedNetMaskLen:     24,
	sharedDefaultGateway: net.IP{192, 0, 2, 1},
	currentMACAddress:    net.HardwareAddr{0x00, 0x00, 0x5E, 0x00, 0x53, 0x00},
	currentSubnets: map[int]*net.IPNet{
		24: {
			IP:   net.IP{24, 0, 0, 0},
			Mask: net.IPMask{255, 255, 255, 0},
		},
		25: {
			IP:   net.IP{25, 0, 0, 0},
			Mask: net.IPMask{255, 255, 255, 128},
		},
		26: {
			IP:   net.IP{26, 0, 0, 0},
			Mask: net.IPMask{255, 255, 255, 192},
		},
		27: {
			IP:   net.IP{27, 0, 0, 0},
			Mask: net.IPMask{255, 255, 255, 224},
		},
		28: {
			IP:   net.IP{28, 0, 0, 0},
			Mask: net.IPMask{255, 255, 255, 240},
		},
	},
}

func (p *valuePool) generateID() types.ID {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.currentID++
	return types.ID(p.currentID)
}

func (p *valuePool) nextSharedIP() net.IP {
	p.mu.Lock()
	defer p.mu.Unlock()

	ip := p.currentSharedIP.To4()
	ip[3]++
	p.currentSharedIP = ip
	return ip
}

func (p *valuePool) nextMACAddress() net.HardwareAddr {
	p.mu.Lock()
	defer p.mu.Unlock()

	mac := []byte(p.currentMACAddress)
	mac[5]++
	p.currentMACAddress = net.HardwareAddr(mac)
	return p.currentMACAddress
}

func (p *valuePool) nextSubnet(maskLen int) *assignedSubnet {
	p.mu.Lock()
	defer p.mu.Unlock()

	next, _ := cidr.NextSubnet(p.currentSubnets[maskLen], maskLen) // ignore result
	p.currentSubnets[maskLen] = next

	count := cidr.AddressCount(next)
	current := next.IP
	var defaultGateway, networkAddr string

	var addresses []string
	for i := uint64(0); i < count; i++ {
		// [0]: ネットワークアドレス
		// [1:3]: ルータ自身が利用
		// [len]: ブロードキャスト
		if i < 4 || i == count-1 {
			if i == 0 {
				networkAddr = current.String()
			}
			if i == 1 {
				defaultGateway = current.String()
			}
			current = cidr.Inc(current)
			continue
		}
		addresses = append(addresses, current.String())
		current = cidr.Inc(current)
	}

	return &assignedSubnet{
		defaultRoute:   defaultGateway,
		networkAddress: networkAddr,
		networkMaskLen: maskLen,
		addresses:      addresses,
	}
}

func (p *valuePool) nextSubnetFull(maskLen int, defaultRoute string) *assignedSubnet {
	p.mu.Lock()
	defer p.mu.Unlock()

	next, _ := cidr.NextSubnet(p.currentSubnets[maskLen], maskLen) // ignore result
	p.currentSubnets[maskLen] = next

	count := cidr.AddressCount(next)
	current := next.IP
	var networkAddr string

	var addresses []string
	for i := uint64(0); i < count; i++ {
		addresses = append(addresses, current.String())
		current = cidr.Inc(current)
	}

	return &assignedSubnet{
		defaultRoute:   defaultRoute,
		networkAddress: networkAddr,
		networkMaskLen: maskLen,
		addresses:      addresses,
	}
}

type assignedSubnet struct {
	defaultRoute   string
	networkMaskLen int
	networkAddress string
	addresses      []string
}
