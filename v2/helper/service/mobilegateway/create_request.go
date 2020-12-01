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

package mobilegateway

import (
	mobileGatewayBuilder "github.com/sacloud/libsacloud/v2/helper/builder/mobilegateway"
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Zone string `request:"-" validate:"required"`

	Name                            string `validate:"required"`
	Description                     string `validate:"min=0,max=512"`
	Tags                            types.Tags
	IconID                          types.ID
	PrivateInterface                *PrivateInterfaceSetting `validate:"omitempty"`
	StaticRoutes                    []*sacloud.MobileGatewayStaticRoute
	SimRoutes                       []*SIMRouteSetting
	InternetConnectionEnabled       bool
	InterDeviceCommunicationEnabled bool
	DNS                             *sacloud.MobileGatewayDNSSetting
	SIMs                            []*SIMSetting
	TrafficConfig                   *sacloud.MobileGatewayTrafficControl

	NoWait bool
}

func (req *CreateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateRequest) Builder(caller sacloud.APICaller) (*mobileGatewayBuilder.Builder, error) {
	builder := &mobileGatewayBuilder.Builder{Client: mobileGatewayBuilder.NewAPIClient(caller)}
	if err := service.RequestConvertTo(req, builder); err != nil {
		return nil, err
	}
	return builder, nil
}
