// Copyright 2016-2022 The Libsacloud Authors
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

package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	archiveAPIName     = "Archive"
	archiveAPIPathName = "archive"
)

var archiveAPI = &dsl.Resource{
	Name:       archiveAPIName,
	PathName:   archiveAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(archiveAPIName, archiveNakedType, findParameter, archiveView),

		// create
		ops.Create(archiveAPIName, archiveNakedType, archiveCreateParam, archiveView),

		// CreateBlank
		{
			ResourceName: archiveAPIName,
			Name:         "CreateBlank",
			PathFormat:   dsl.DefaultPathFormat,
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Name: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
				Type: archiveNakedType,
			}),
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
					Type: archiveNakedType,
				},
				&dsl.EnvelopePayloadDesc{
					Name: models.ftpServer().Name,
					Type: meta.Static(naked.OpeningFTPServer{}),
				},
			),
			Arguments: dsl.Arguments{
				dsl.MappableArgument("param", archiveCreateBlankParam, names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular)),
			},
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
					DestField:   archiveView.Name,
					IsPlural:    false,
					Model:       archiveView,
				},
				{
					SourceField: models.ftpServer().Name,
					DestField:   models.ftpServer().Name,
					IsPlural:    false,
					Model:       models.ftpServer(),
				},
			},
		},

		// TODO 他ゾーンからの転送コピー作成

		// read
		ops.Read(archiveAPIName, archiveNakedType, archiveView),

		// update
		ops.Update(archiveAPIName, archiveNakedType, archiveUpdateParam, archiveView),

		// delete
		ops.Delete(archiveAPIName),

		// openFTP
		ops.OpenFTP(archiveAPIName, models.ftpServerOpenParameter(), models.ftpServer()),

		// closeFTP
		ops.CloseFTP(archiveAPIName),

		// Share
		&dsl.Operation{
			ResourceName: archiveAPIName,
			Name:         "Share",
			PathFormat:   dsl.IDAndSuffixPathFormat("ftp"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: "Shared", // sacloudパッケージ内でMarshalJSON時に設定される
					Type: meta.TypeFlag,
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: "ArchiveShareInfo",
					Type: meta.Static(naked.ArchiveShareInfo{}),
				},
			),
			Results: dsl.Results{
				{
					SourceField: "ArchiveShareInfo",
					DestField:   "ArchiveShareInfo",
					IsPlural:    false,
					Model:       archiveShareInfo,
				}, // sacloudパッケージ内のcustomized_envelopeで設定される
			},
		},
		// CreateFromShared
		{
			ResourceName: archiveAPIName,
			Name:         "CreateFromShared",
			PathFormat:   dsl.DefaultPathFormat + "/{{.sourceArchiveID}}/to/zone/{{.destZoneID}}",
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Name: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
				Type: archiveNakedType,
			}),
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
					Type: archiveNakedType,
				},
			),
			Arguments: dsl.Arguments{
				&dsl.Argument{Name: "sourceArchiveID", Type: meta.TypeID},
				&dsl.Argument{Name: "destZoneID", Type: meta.TypeID},
				dsl.MappableArgument("param", archiveCreateFromSharedParam, names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular)),
			},
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
					DestField:   archiveView.Name,
					IsPlural:    false,
					Model:       archiveView,
				},
			},
		},
		// Transfer
		{
			ResourceName: archiveAPIName,
			Name:         "Transfer",
			PathFormat:   dsl.DefaultPathFormat + "/{{.sourceArchiveID}}/to/zone/{{.destZoneID}}",
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Name: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
				Type: archiveNakedType,
			}),
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
					Type: archiveNakedType,
				},
			),
			Arguments: dsl.Arguments{
				&dsl.Argument{Name: "sourceArchiveID", Type: meta.TypeID},
				&dsl.Argument{Name: "destZoneID", Type: meta.TypeID},
				dsl.MappableArgument("param", archiveTransferParam, names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular)),
			},
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
					DestField:   archiveView.Name,
					IsPlural:    false,
					Model:       archiveView,
				},
			},
		},
	},
}

var (
	archiveNakedType = meta.Static(naked.Archive{})

	archiveView = &dsl.Model{
		Name:      archiveAPIName,
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.DisplayOrder(),
			fields.Availability(),
			fields.Scope(),
			fields.SizeMB(),
			fields.MigratedMB(),
			fields.DiskPlanID(),
			fields.DiskPlanName(),
			fields.DiskPlanStorageClass(),
			fields.SourceDiskID(),
			fields.SourceDiskAvailability(),
			fields.SourceArchiveID(),
			fields.SourceArchiveAvailability(),
			fields.BundleInfo(),
			fields.Storage(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.OriginalArchiveID(),
			fields.SourceInfo(),
		},
	}

	archiveCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(archiveAPIName),
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.SourceDiskID(),
			fields.SourceArchiveID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	archiveCreateBlankParam = &dsl.Model{
		Name:      "ArchiveCreateBlankRequest",
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	archiveUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(archiveAPIName),
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	archiveShareInfo = &dsl.Model{
		Name:      "ArchiveShareInfo",
		NakedType: meta.Static(naked.ArchiveShareInfo{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("SharedKey", meta.Static(types.ArchiveShareKey(""))),
		},
	}

	archiveCreateFromSharedParam = &dsl.Model{
		Name:      names.CreateParameterName(archiveAPIName) + "FromShared",
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.Def("SourceSharedKey", meta.Static(types.ArchiveShareKey(""))),
		},
	}

	archiveTransferParam = &dsl.Model{
		Name:      "ArchiveTransferRequest",
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
