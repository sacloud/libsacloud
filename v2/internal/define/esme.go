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

package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

const (
	esmeAPIName     = "ESME"
	esmeAPIPathName = "commonserviceitem"
)

var esmeAPI = &dsl.Resource{
	Name:       esmeAPIName,
	PathName:   esmeAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(esmeAPIName, esmeNakedType, findParameter, esmeView),

		// create
		ops.CreateCommonServiceItem(esmeAPIName, esmeNakedType, esmeCreateParam, esmeView),

		// read
		ops.ReadCommonServiceItem(esmeAPIName, esmeNakedType, esmeView),

		// update
		ops.UpdateCommonServiceItem(esmeAPIName, esmeNakedType, esmeUpdateParam, esmeView),

		// delete
		ops.Delete(esmeAPIName),

		// SendMessageWithGeneratedOTP
		{
			ResourceName: esmeAPIName,
			Name:         "SendMessageWithGeneratedOTP",
			PathFormat:   dsl.IDAndSuffixPathFormat("esme/2fa/otp"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.ESMESendSMSRequest{}),
				Name: esmeAPIName,
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", esmeSendMessageWithGeneratedOTPParam, "ESME"),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.ESMESendSMSResponse{}),
				Name: esmeAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: esmeAPIName,
					DestField:   esmeSendMessageResult.Name,
					Model:       esmeSendMessageResult,
				},
			},
		},

		// SendMessageWithInputtedOTP
		{
			ResourceName: esmeAPIName,
			Name:         "SendMessageWithInputtedOTP",
			PathFormat:   dsl.IDAndSuffixPathFormat("esme/2fa"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.ESMESendSMSRequest{}),
				Name: esmeAPIName,
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", esmeSendMessageWithInputtedOTPParam, "ESME"),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.ESMESendSMSResponse{}),
				Name: esmeAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: esmeAPIName,
					DestField:   esmeSendMessageResult.Name,
					Model:       esmeSendMessageResult,
				},
			},
		},

		// Logs
		{
			ResourceName: esmeAPIName,
			Name:         "Logs",
			PathFormat:   dsl.IDAndSuffixPathFormat("esme/logs"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "ESME",
				Type: meta.Static(&naked.ESMELogs{}),
			}),
			Results: dsl.Results{
				{
					SourceField: "ESME.Logs",
					DestField:   "Logs",
					IsPlural:    true,
					Model:       esmeLogsModel,
				},
			},
		},
	},
}

var (
	esmeNakedType = meta.Static(naked.ESME{})

	esmeView = &dsl.Model{
		Name:      esmeAPIName,
		NakedType: esmeNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	esmeCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(esmeAPIName),
		NakedType: esmeNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"esme"`,
			},
		},

		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	esmeUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(esmeAPIName),
		NakedType: esmeNakedType,
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	esmeSendMessageWithGeneratedOTPParam = &dsl.Model{
		Name:      "ESMESendMessageWithGeneratedOTPRequest",
		NakedType: meta.Static(naked.ESMESendSMSRequest{}),
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "OTPOperation",
				Type:  meta.Static(types.EOTPOperation("")),
				Value: `"` + types.OTPOperations.Generate.String() + `"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Def("Destination", meta.TypeString), // 宛先 81開始
			fields.Def("Sender", meta.TypeString),
		},
	}

	esmeSendMessageWithInputtedOTPParam = &dsl.Model{
		Name:      "ESMESendMessageWithInputtedOTPRequest",
		NakedType: meta.Static(naked.ESMESendSMSRequest{}),
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "OTPOperation",
				Type:  meta.Static(types.EOTPOperation("")),
				Value: `"` + types.OTPOperations.Input.String() + `"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Def("Destination", meta.TypeString), // 宛先 81開始
			fields.Def("Sender", meta.TypeString),
			fields.Def("OTP", meta.TypeString),
		},
	}

	esmeSendMessageResult = &dsl.Model{
		Name:      "ESMESendMessageResult",
		NakedType: meta.Static(naked.ESMESendSMSResponse{}),
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Def("MessageID", meta.TypeString),
			fields.Def("Status", meta.TypeString), // TODO typesに型定義したいが不明な値があるためstringとしている
			fields.Def("OTP", meta.TypeString),
		},
	}

	esmeLogsModel = &dsl.Model{
		Name:      esmeAPIName + "Logs",
		NakedType: meta.Static(naked.ESMELog{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("MessageID", meta.TypeString),
			fields.Def("Status", meta.TypeString), // TODO typesに型定義したいが不明な値があるためstringとしている
			fields.Def("OTP", meta.TypeString),
			fields.Def("Destination", meta.TypeString),
			fields.Def("SentAt", meta.TypeTime),
			fields.Def("DoneAt", meta.TypeTime),
			fields.Def("RetryCount", meta.TypeInt),
		},
	}
)
