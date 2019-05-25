package types

import (
	"fmt"
	"sort"
	"strings"
)

// VPCFirewallPort パケットフィルターのルールでのポート型
//
// 1～65535 の整数、その範囲指定（ハイフン区切り）、または複数指定（コンマ区切り 6個まで）の形式の形式のみ受け付ける
type VPCFirewallPort string

// SetPort 単一のポート番号の設定
func (p *VPCFirewallPort) SetPort(port int) {
	*p = VPCFirewallPort(fmt.Sprintf("%d", port))
}

// SetPortRange ポート範囲指定
func (p *VPCFirewallPort) SetPortRange(from, to int) {
	*p = VPCFirewallPort(fmt.Sprintf("%d-%d", from, to))
}

// SetPortMultiple 複数のポートを指定(6個まで)
func (p *VPCFirewallPort) SetPortMultiple(ports ...int) {
	sort.Ints(ports)
	strPort := make([]string, len(ports))
	for i := range ports {
		strPort[i] = fmt.Sprintf("%d", ports[i])
	}

	*p = VPCFirewallPort(strings.Join(strPort, ","))
}

// IsEmpty 値が指定されているか
func (p *VPCFirewallPort) IsEmpty() bool {
	return p != nil && p.String() != ""
}

// String 文字列表現
func (p *VPCFirewallPort) String() string {
	return string(*p)
}

// Equal 指定のパケットフィルタポートと同じ値を持つか
func (p *VPCFirewallPort) Equal(p2 *PacketFilterPort) bool {
	return p.String() == p2.String()
}
