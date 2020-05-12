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

package archive

import (
	"context"
	"fmt"
	"os"

	"github.com/sacloud/ftps"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UploadRequest struct {
	Zone string   `validate:"required" mapconv:"-"`
	ID   types.ID `validate:"required" mapconv:"-"`

	Path string `validate:"omitempty,file"`
}

func (r *UploadRequest) Validate() error {
	return validate.Struct(r)
}

func (s *Service) Upload(req *UploadRequest) error {
	return s.UploadWithContext(context.Background(), req)
}

func (s *Service) UploadWithContext(ctx context.Context, req *UploadRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewArchiveOp(s.caller)
	archive, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return fmt.Errorf("reading archive[%s] failed: %s", req.ID, err)
	}

	if archive.Scope != types.Scopes.User {
		return fmt.Errorf("archive[%s] is not allowed to download", req.ID)
	}

	ftpServer, err := client.OpenFTP(ctx, req.Zone, archive.ID, &sacloud.OpenFTPRequest{ChangePassword: true})
	if err != nil {
		return fmt.Errorf("requesting FTP server information failed: %s", err)
	}

	// upload
	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)

	var file *os.File
	switch req.Path {
	case "":
		file = os.Stdin
	default:
		f, err := os.Open(req.Path)
		if err != nil {
			return fmt.Errorf("opening upload file failed: %s", err)
		}
		defer f.Close()
		file = f
	}

	if err := ftpsClient.UploadFile("upload.raw", file); err != nil {
		return fmt.Errorf("uploading file to archive[%s] failed: %s", archive.ID, err)
	}

	// close FTP
	if err := client.CloseFTP(ctx, req.Zone, archive.ID); err != nil {
		return fmt.Errorf("closing FTP server failed: %s", err)
	}
	return nil
}
