package fake

import (
	"net"
	"sync"
)

type addressPool struct {
	currentSharedIP      net.IP
	sharedNetMaskLen     int
	sharedDefaultGateway net.IP
	currentMACAddress    net.HardwareAddr
	mu                   sync.Mutex
}

var addrPool = &addressPool{
	currentSharedIP:      net.IP{192, 0, 2, 2},
	sharedNetMaskLen:     24,
	sharedDefaultGateway: net.IP{192, 0, 2, 1},
	currentMACAddress:    net.HardwareAddr{0x00, 0x00, 0x5E, 0x00, 0x53, 0x00},
}

func (p *addressPool) nextSharedIP() net.IP {
	p.mu.Lock()
	defer p.mu.Unlock()

	ip := p.currentSharedIP.To4()
	ip[3]++
	p.currentSharedIP = ip
	return ip
}

func (p *addressPool) nextMACAddress() net.HardwareAddr {
	p.mu.Lock()
	defer p.mu.Unlock()

	mac := []byte(p.currentMACAddress)
	mac[5]++
	p.currentMACAddress = net.HardwareAddr(mac)
	return p.currentMACAddress
}
