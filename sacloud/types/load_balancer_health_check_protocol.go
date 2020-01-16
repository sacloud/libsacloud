// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// ELoadBalancerHealthCheckProtocol ロードバランサ 監視プロトコル
type ELoadBalancerHealthCheckProtocol string

// String ロードバランサ 監視プロトコルの文字列表現
func (p ELoadBalancerHealthCheckProtocol) String() string {
	return string(p)
}

// LoadBalancerHealthCheckProtocols ロードバランサ 監視プロトコル
var LoadBalancerHealthCheckProtocols = struct {
	// Unknown 不明
	Unknown ELoadBalancerHealthCheckProtocol
	// HTTP http
	HTTP ELoadBalancerHealthCheckProtocol
	// HTTPS https
	HTTPS ELoadBalancerHealthCheckProtocol
	// TCP tcp
	TCP ELoadBalancerHealthCheckProtocol
	// Ping ping
	Ping ELoadBalancerHealthCheckProtocol
}{
	Unknown: ELoadBalancerHealthCheckProtocol(""),
	HTTP:    ELoadBalancerHealthCheckProtocol("http"),
	HTTPS:   ELoadBalancerHealthCheckProtocol("https"),
	TCP:     ELoadBalancerHealthCheckProtocol("tcp"),
	Ping:    ELoadBalancerHealthCheckProtocol("ping"),
}

// LoadBalancerHealthCheckProtocolsStrings 有効なロードバランサ監視プロトコルを示す文字列のリスト
//
// Unknown(空文字)は含まない
func LoadBalancerHealthCheckProtocolsStrings() []string {
	return []string{
		LoadBalancerHealthCheckProtocols.HTTP.String(),
		LoadBalancerHealthCheckProtocols.HTTPS.String(),
		LoadBalancerHealthCheckProtocols.TCP.String(),
		LoadBalancerHealthCheckProtocols.Ping.String(),
	}
}
