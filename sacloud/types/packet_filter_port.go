package types

import "fmt"

// PacketFilterPort パケットフィルターのルールでのポート型
//
// 0〜65535 の整数、その範囲指定（ハイフン区切り）の形式のみ受け付ける
type PacketFilterPort string

// SetPort 単一のポート番号の設定
func (p *PacketFilterPort) SetPort(port int) {
	*p = PacketFilterPort(fmt.Sprintf("%d", port))
}

// SetPortRange ポート範囲指定
func (p *PacketFilterPort) SetPortRange(from, to int) {
	*p = PacketFilterPort(fmt.Sprintf("%d-%d", from, to))
}

// IsEmpty 値が指定されているか
func (p *PacketFilterPort) IsEmpty() bool {
	return p != nil && p.String() != ""
}

// String 文字列表現
func (p *PacketFilterPort) String() string {
	return string(*p)
}

// Equal 指定のパケットフィルタポートと同じ値を持つか
func (p *PacketFilterPort) Equal(p2 *PacketFilterPort) bool {
	return p.String() == p2.String()
}
