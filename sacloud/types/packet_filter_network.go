package types

import "fmt"

// PacketFilterNetwork パケットフィルタのルールでの送信元ネットワーク(アドレス/範囲)
//
// A.A.A.A、A.A.A.A/N (N=1〜31)、またはA.A.A.A/M.M.M.M形式を指定可能
type PacketFilterNetwork string

// SetAddress 単一のIPアドレスを指定
func (p *PacketFilterNetwork) SetAddress(ip string) {
	*p = PacketFilterNetwork(ip)
}

// SetAddressRange アドレス範囲を指定
func (p *PacketFilterNetwork) SetAddressRange(from, to string) {
	*p = PacketFilterNetwork(fmt.Sprintf("%s/%s", from, to))
}

// SetNetworkAddress ネットワークアドレスを指定
func (p *PacketFilterNetwork) SetNetworkAddress(networkAddr string, maskLen int) {
	*p = PacketFilterNetwork(fmt.Sprintf("%s/%d", networkAddr, maskLen))
}

// String 文字列表現
func (p *PacketFilterNetwork) String() string {
	return string(*p)
}

// Equal 指定の送信元ネットワークと同じ値を持つか
//
// "192.168.0.0/192.168.0.255"と"192.168.0.0/24"は異なる値とみなす
func (p *PacketFilterNetwork) Equal(p2 *PacketFilterNetwork) bool {
	return p.String() == p2.String()
}
