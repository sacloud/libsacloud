package types

// LoadBalancerPlans ロードバランサーのプラン
var LoadBalancerPlans = struct {
	// Standard スタンダード
	Standard ID
	// Premium プレミアム
	Premium ID
}{
	Standard: ID(1),
	Premium:  ID(2),
}
