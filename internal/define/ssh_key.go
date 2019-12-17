// Copyright 2016-2019 The Libsacloud Authors
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

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	sshKeyAPIName     = "SSHKey"
	sshKeyAPIPathName = "sshkey"
)

var sshKeyAPI = &dsl.Resource{
	Name:       sshKeyAPIName,
	PathName:   sshKeyAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.Find(sshKeyAPIName, sshKeyNakedType, findParameter, sshKeyView),

		// create
		ops.Create(sshKeyAPIName, sshKeyNakedType, sshKeyCreateParam, sshKeyView),

		// generate
		{
			ResourceName: sshKeyAPIName,
			Name:         "Generate",
			PathFormat:   dsl.DefaultPathFormat + "/generate",
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: sshKeyNakedType,
				Name: sshKeyAPIName,
			}),
			Arguments: dsl.Arguments{
				dsl.MappableArgument("param", sshKeyGenerateParam, sshKeyAPIName),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: sshKeyNakedType,
				Name: sshKeyAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: sshKeyAPIName,
					DestField:   sshKeyGeneratedView.Name,
					IsPlural:    false,
					Model:       sshKeyGeneratedView,
				},
			},
		},

		// read
		ops.Read(sshKeyAPIName, sshKeyNakedType, sshKeyView),

		// update
		ops.Update(sshKeyAPIName, sshKeyNakedType, sshKeyUpdateParam, sshKeyView),

		// delete
		ops.Delete(sshKeyAPIName),
	},
}

var (
	sshKeyNakedType = meta.Static(naked.SSHKey{})

	sshKeyFields = []*dsl.FieldDesc{
		fields.ID(),
		fields.Name(),
		fields.Description(),
		fields.CreatedAt(),
		fields.PublicKey(),
		fields.Fingerprint(),
	}

	sshKeyView = &dsl.Model{
		Name:      sshKeyAPIName,
		NakedType: sshKeyNakedType,
		Fields:    sshKeyFields,
	}

	sshKeyGeneratedView = &dsl.Model{
		Name:      sshKeyAPIName + "Generated",
		NakedType: sshKeyNakedType,
		Fields:    append(sshKeyFields, fields.PrivateKey()),
	}

	sshKeyCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(sshKeyAPIName),
		NakedType: sshKeyNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PublicKey(),
		},
	}

	sshKeyGenerateParam = &dsl.Model{
		Name:      names.RequestParameterName(sshKeyAPIName, "Generate"),
		NakedType: sshKeyNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PassPhrase(),
		},
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "GenerateFormat",
				Type:  meta.TypeString,
				Value: `"openssh"`,
			},
		},
	}

	sshKeyUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(sshKeyAPIName),
		NakedType: sshKeyNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
		},
	}
)
