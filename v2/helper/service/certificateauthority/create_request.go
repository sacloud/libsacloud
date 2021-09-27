// Copyright 2016-2021 The Libsacloud Authors
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

package certificateauthority

import (
	"time"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Name        string `validate:"required"`
	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID

	Country          string
	Organization     string
	OrganizationUnit []string
	CommonName       string
	NotAfter         time.Time

	Clients []*ClientCert // Note: API的に証明書の削除はできないため、指定した以上の証明書が存在する可能性がある
	Servers []*ServerCert // Note: API的に証明書の削除はできないため、指定した以上の証明書が存在する可能性がある

	WaitDuration time.Duration // 証明書発行待ち時間、省略した場合10秒
}

func (req *CreateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateRequest) ApplyRequest() *ApplyRequest {
	return &ApplyRequest{
		Name:             req.Name,
		Description:      req.Description,
		Tags:             req.Tags,
		IconID:           req.IconID,
		Country:          req.Country,
		Organization:     req.Organization,
		OrganizationUnit: req.OrganizationUnit,
		CommonName:       req.CommonName,
		NotAfter:         req.NotAfter,
		Clients:          req.Clients,
		Servers:          req.Servers,
		WaitDuration:     req.WaitDuration,
	}
}