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

package vpcrouter

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `request:"-" validate:"required"`
	ID   types.ID `request:"-" validate:"required"`

	Name        *string     `request:",omitempty" validate:"omitempty,min=1"`
	Description *string     `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags        *types.Tags `request:",omitempty"`
	IconID      *types.ID   `request:",omitempty"`

	NICSetting            *PremiumNICSetting             `request:",omitempty"`
	AdditionalNICSettings *[]AdditionalPremiumNICSetting `request:",omitempty"`
	RouterSetting         *RouterSetting                 `request:",omitempty"`
	NoWait                bool

	SettingsHash string
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) ApplyRequest(ctx context.Context, caller sacloud.APICaller) (*ApplyRequest, error) {
	current, err := sacloud.NewVPCRouterOp(caller).Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}

	if current.PlanID == types.VPCRouterPlans.Standard {
		return nil, fmt.Errorf("target VPCRouter(Zone=%s ID=%q) is not a premium or higher plan", req.Zone, req.ID)
	}
	if current.Availability != types.Availabilities.Available {
		return nil, fmt.Errorf("target VPCRouter has invalid Availability: %v", current.Availability)
	}

	var additionalNICs []AdditionalNICSettingHolder
	for _, nic := range current.Interfaces {
		if nic.Index == 0 {
			continue
		}
		var setting *sacloud.VPCRouterInterfaceSetting
		for _, s := range current.Settings.Interfaces {
			if s.Index == nic.Index {
				setting = s
				break
			}
		}
		if setting == nil {
			continue
		}

		additionalNICs = append(additionalNICs, &AdditionalPremiumNICSetting{
			SwitchID:         nic.SwitchID,
			IPAddresses:      setting.IPAddress,
			VirtualIPAddress: setting.VirtualIPAddress,
			NetworkMaskLen:   setting.NetworkMaskLen,
			Index:            setting.Index,
		})
	}

	var nicSetting *PremiumNICSetting
	for _, s := range current.Settings.Interfaces {
		if s.Index == 0 {
			nicSetting = &PremiumNICSetting{
				SwitchID:         current.Interfaces[0].SwitchID,
				IPAddresses:      s.IPAddress,
				VirtualIPAddress: s.VirtualIPAddress,
				IPAliases:        s.IPAliases,
			}
			break
		}
	}

	applyRequest := &ApplyRequest{
		Zone:                  req.Zone,
		ID:                    req.ID,
		Name:                  current.Name,
		Description:           current.Description,
		Tags:                  current.Tags,
		IconID:                current.IconID,
		PlanID:                current.PlanID,
		NICSetting:            nicSetting,
		AdditionalNICSettings: additionalNICs,
		RouterSetting: &RouterSetting{
			VRID:                      current.Settings.VRID,
			InternetConnectionEnabled: current.Settings.InternetConnectionEnabled,
			StaticNAT:                 current.Settings.StaticNAT,
			PortForwarding:            current.Settings.PortForwarding,
			Firewall:                  current.Settings.Firewall,
			DHCPServer:                current.Settings.DHCPServer,
			DHCPStaticMapping:         current.Settings.DHCPStaticMapping,
			PPTPServer:                current.Settings.PPTPServer,
			L2TPIPsecServer:           current.Settings.L2TPIPsecServer,
			RemoteAccessUsers:         current.Settings.RemoteAccessUsers,
			SiteToSiteIPsecVPN:        current.Settings.SiteToSiteIPsecVPN,
			StaticRoute:               current.Settings.StaticRoute,
			SyslogHost:                current.Settings.SyslogHost,
		},
		NoWait: false,
	}

	if err := service.RequestConvertTo(req, applyRequest); err != nil {
		return nil, err
	}
	return applyRequest, nil
}
