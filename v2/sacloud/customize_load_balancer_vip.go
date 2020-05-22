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

package sacloud

type LoadBalancerVirtualIPAddresses []*LoadBalancerVirtualIPAddress

// AddGSLBServer サーバの追加
func (o *LoadBalancerVirtualIPAddresses) Add(vip *LoadBalancerVirtualIPAddress) {
	if o.Exist(vip) {
		return // noop if already exists
	}
	*o = append(*o, vip)
}

// Exist サーバの存在確認
func (o *LoadBalancerVirtualIPAddresses) Exist(vip *LoadBalancerVirtualIPAddress) bool {
	for _, v := range *o {
		if v.VirtualIPAddress == vip.VirtualIPAddress {
			return true
		}
	}
	return false
}

// ExistAt サーバの存在確認
func (o *LoadBalancerVirtualIPAddresses) ExistAt(vip string) bool {
	return o.Exist(&LoadBalancerVirtualIPAddress{VirtualIPAddress: vip})
}

// Find サーバの検索
func (o *LoadBalancerVirtualIPAddresses) Find(vip *LoadBalancerVirtualIPAddress) *LoadBalancerVirtualIPAddress {
	for _, v := range *o {
		if v.VirtualIPAddress == vip.VirtualIPAddress {
			return v
		}
	}
	return nil
}

// FindAt サーバの検索
func (o *LoadBalancerVirtualIPAddresses) FindAt(vip string) *LoadBalancerVirtualIPAddress {
	return o.Find(&LoadBalancerVirtualIPAddress{VirtualIPAddress: vip})
}

// Update サーバの更新
func (o *LoadBalancerVirtualIPAddresses) Update(old *LoadBalancerVirtualIPAddress, new *LoadBalancerVirtualIPAddress) {
	for _, v := range *o {
		if v.VirtualIPAddress == old.VirtualIPAddress {
			*v = *new
			return
		}
	}
}

// UpdateAt サーバの更新
func (o *LoadBalancerVirtualIPAddresses) UpdateAt(vip string, new *LoadBalancerVirtualIPAddress) {
	o.Update(&LoadBalancerVirtualIPAddress{VirtualIPAddress: vip}, new)
}

// Delete サーバの削除
func (o *LoadBalancerVirtualIPAddresses) Delete(vip *LoadBalancerVirtualIPAddress) {
	var res []*LoadBalancerVirtualIPAddress
	for _, v := range *o {
		if v.VirtualIPAddress != vip.VirtualIPAddress {
			res = append(res, v)
		}
	}
	*o = res
}

// DeleteAt サーバの削除
func (o *LoadBalancerVirtualIPAddresses) DeleteAt(vip string) {
	o.Delete(&LoadBalancerVirtualIPAddress{VirtualIPAddress: vip})
}
