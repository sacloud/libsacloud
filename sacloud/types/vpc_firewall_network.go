package types

import "fmt"

// VPCFirewallNetwork VPCルータのファイアウォールルールでの送信元ネットワーク(アドレス/範囲)
//
// A.A.A.A、A.A.A.A/N (N=1〜31)形式を指定可能
type VPCFirewallNetwork string

// SetAddress 単一のIPアドレスを指定
func (p *VPCFirewallNetwork) SetAddress(ip string) {
	*p = VPCFirewallNetwork(ip)
}

// SetNetworkAddress ネットワークアドレスを指定
func (p *VPCFirewallNetwork) SetNetworkAddress(networkAddr string, maskLen int) {
	*p = VPCFirewallNetwork(fmt.Sprintf("%s/%d", networkAddr, maskLen))
}

// String 文字列表現
func (p *VPCFirewallNetwork) String() string {
	return string(*p)
}

// Equal 指定の送信元ネットワークと同じ値を持つか
func (p *VPCFirewallNetwork) Equal(p2 *PacketFilterNetwork) bool {
	return p.String() == p2.String()
}
