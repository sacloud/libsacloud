package fake

import (
	"net"
	"sync"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

type valuePool struct {
	currentID            int64
	currentSharedIP      net.IP
	sharedNetMaskLen     int
	sharedDefaultGateway net.IP
	currentMACAddress    net.HardwareAddr
	mu                   sync.Mutex
}

var pool = &valuePool{
	currentID:            int64(100000000000),
	currentSharedIP:      net.IP{192, 0, 2, 2},
	sharedNetMaskLen:     24,
	sharedDefaultGateway: net.IP{192, 0, 2, 1},
	currentMACAddress:    net.HardwareAddr{0x00, 0x00, 0x5E, 0x00, 0x53, 0x00},
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
