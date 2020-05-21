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

package cdrom

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/sacloud/ftps"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/pkg/size"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Zone string `validate:"required" mapconv:"-"`

	Name        string `validate:"required"`
	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID
	SizeGB      int

	// for blank builder
	SourcePath   string `validate:"omitempty,file"`
	SourceReader io.Reader
}

func (r *CreateRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CreateRequest) toRequestParameter() (*sacloud.CDROMCreateRequest, error) {
	res := &sacloud.CDROMCreateRequest{}
	if err := mapconv.ConvertTo(r, res); err != nil {
		return nil, err
	}
	res.SizeMB = r.SizeGB * size.GiB
	return res, nil
}

func (s *Service) Create(req *CreateRequest) (*sacloud.CDROM, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*sacloud.CDROM, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var reader io.Reader
	switch req.SourcePath {
	case "":
		reader = req.SourceReader
	default:
		file, err := os.Open(req.SourcePath)
		if err != nil {
			return nil, fmt.Errorf("reading source file[%s] failed: %s", req.SourcePath, err)
		}
		defer file.Close() // nolint
		reader = file
	}

	params, err := req.toRequestParameter()
	if err != nil {
		return nil, err
	}

	client := sacloud.NewCDROMOp(s.caller)
	cdrom, ftpServer, err := client.Create(ctx, req.Zone, params)
	if err != nil {
		return nil, err
	}

	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	if err := ftpsClient.UploadReader("data.iso", reader); err != nil {
		return nil, err
	}

	if err := client.CloseFTP(ctx, req.Zone, cdrom.ID); err != nil {
		return nil, err
	}

	// refresh
	return client.Read(ctx, req.Zone, cdrom.ID)
}
