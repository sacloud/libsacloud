package types

// VPCRouterPlans VPCルータのプラン
var VPCRouterPlans = struct {
	// Standard スタンダードプラン シングル構成/最大スループット 80Mbps/一部機能は利用不可
	Standard ID
	// Premium プレミアムプラン 冗長構成/最大スループット400Mbps
	Premium ID
	// HighSpec ハイスペックプラン 冗長構成/最大スループット1,200Mbps
	HighSpec ID
}{
	Standard: ID(1),
	Premium:  ID(2),
	HighSpec: ID(3),
}
