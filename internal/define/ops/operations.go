package ops

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

func find(resourceName string, nakedType meta.Type, findParam, result *dsl.Model, payloadName string) *dsl.Operation {
	if payloadName == "" {
		payloadName = names.ResourceFieldName(resourceName, dsl.PayloadForms.Plural)
	}

	return &dsl.Operation{
		ResourceName:     resourceName,
		Name:             "Find",
		PathFormat:       dsl.DefaultPathFormat,
		Method:           http.MethodGet,
		UseWrappedResult: true,
		RequestEnvelope:  dsl.RequestEnvelopeFromModel(findParam),
		Arguments: dsl.Arguments{
			dsl.PassthroughModelArgument("conditions", findParam),
		},
		ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: dsl.Results{
			{
				SourceField: payloadName,
				DestField:   names.ResourceFieldName(resourceName, dsl.PayloadForms.Plural),
				IsPlural:    true,
				Model:       result,
			},
		},
	}
}

// Find Find操作を定義
func Find(resourceName string, nakedType meta.Type, findParam, result *dsl.Model) *dsl.Operation {
	return find(resourceName, nakedType, findParam, result, "")
}

// FindAppliance Find操作を定義
func FindAppliance(resourceName string, nakedType meta.Type, findParam, result *dsl.Model) *dsl.Operation {
	return find(resourceName, nakedType, findParam, result, "Appliances")
}

// FindCommonServiceItem Find操作を定義
func FindCommonServiceItem(resourceName string, nakedType meta.Type, findParam, result *dsl.Model) *dsl.Operation {
	return find(resourceName, nakedType, findParam, result, "CommonServiceItems")
}

// List List操作(パラメータのないFind)を定義
func List(resourceName string, nakedType meta.Type, result *dsl.Model) *dsl.Operation {
	return &dsl.Operation{
		ResourceName:     resourceName,
		Name:             "List",
		PathFormat:       dsl.DefaultPathFormat,
		Method:           http.MethodGet,
		UseWrappedResult: true,
		ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: names.ResourceFieldName(resourceName, dsl.PayloadForms.Plural),
		}),
		Results: dsl.Results{
			{
				SourceField: names.ResourceFieldName(resourceName, dsl.PayloadForms.Plural),
				DestField:   names.ResourceFieldName(resourceName, dsl.PayloadForms.Plural),
				IsPlural:    true,
				Model:       result,
			},
		},
	}
}

func create(resourceName string, nakedType meta.Type, createParam, result *dsl.Model, payloadName string) *dsl.Operation {
	if payloadName == "" {
		payloadName = names.ResourceFieldName(resourceName, dsl.PayloadForms.Singular)
	}

	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "Create",
		PathFormat:   dsl.DefaultPathFormat,
		Method:       http.MethodPost,
		RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Arguments: dsl.Arguments{
			dsl.MappableArgument("param", createParam, payloadName),
		},
		ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: dsl.Results{
			{
				SourceField: payloadName,
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}

// Create Create操作を定義
func Create(resourceName string, nakedType meta.Type, createParam, result *dsl.Model) *dsl.Operation {
	return create(resourceName, nakedType, createParam, result, "")
}

// CreateAppliance Create操作を定義
func CreateAppliance(resourceName string, nakedType meta.Type, createParam, result *dsl.Model) *dsl.Operation {
	return create(resourceName, nakedType, createParam, result, "Appliance")
}

// CreateCommonServiceItem Create操作を定義
func CreateCommonServiceItem(resourceName string, nakedType meta.Type, createParam, result *dsl.Model) *dsl.Operation {
	return create(resourceName, nakedType, createParam, result, "CommonServiceItem")
}

func read(resourceName string, nakedType meta.Type, result *dsl.Model, payloadName string) *dsl.Operation {
	if payloadName == "" {
		payloadName = names.ResourceFieldName(resourceName, dsl.PayloadForms.Singular)
	}

	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "Read",
		PathFormat:   dsl.DefaultPathFormatWithID,
		Method:       http.MethodGet,
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
		},
		ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: dsl.Results{
			{
				SourceField: payloadName,
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}

// Read Read操作を定義
func Read(resourceName string, nakedType meta.Type, result *dsl.Model) *dsl.Operation {
	return read(resourceName, nakedType, result, "")
}

// ReadAppliance Read操作を定義
func ReadAppliance(resourceName string, nakedType meta.Type, result *dsl.Model) *dsl.Operation {
	return read(resourceName, nakedType, result, "Appliance")
}

// ReadCommonServiceItem Read操作を定義
func ReadCommonServiceItem(resourceName string, nakedType meta.Type, result *dsl.Model) *dsl.Operation {
	return read(resourceName, nakedType, result, "CommonServiceItem")
}

func update(resourceName string, nakedType meta.Type, updateParam, result *dsl.Model, payloadName string) *dsl.Operation {
	if payloadName == "" {
		payloadName = names.ResourceFieldName(resourceName, dsl.PayloadForms.Singular)
	}

	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "Update",
		PathFormat:   dsl.DefaultPathFormatWithID,
		Method:       http.MethodPut,
		RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
			dsl.MappableArgument("param", updateParam, payloadName),
		},
		ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: dsl.Results{
			{
				SourceField: payloadName,
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}

// Update Update操作を定義
func Update(resourceName string, nakedType meta.Type, updateParam, result *dsl.Model) *dsl.Operation {
	return update(resourceName, nakedType, updateParam, result, "")
}

// UpdateAppliance Update操作を定義
func UpdateAppliance(resourceName string, nakedType meta.Type, updateParam, result *dsl.Model) *dsl.Operation {
	return update(resourceName, nakedType, updateParam, result, "Appliance")
}

// UpdateCommonServiceItem Update操作を定義
func UpdateCommonServiceItem(resourceName string, nakedType meta.Type, updateParam, result *dsl.Model) *dsl.Operation {
	return update(resourceName, nakedType, updateParam, result, "CommonServiceItem")
}

func patch(resourceName string, nakedType meta.Type, updateParam, result *dsl.Model, payloadName string) *dsl.Operation {
	op := update(resourceName, nakedType, updateParam, result, payloadName)
	op.Name = "Patch"
	op.IsPatch = true
	return op
}

// Patch Patch操作を定義
func Patch(resourceName string, nakedType meta.Type, updateParam, result *dsl.Model) *dsl.Operation {
	return patch(resourceName, nakedType, updateParam, result, "")
}

// PatchAppliance Patch操作を定義
func PatchAppliance(resourceName string, nakedType meta.Type, updateParam, result *dsl.Model) *dsl.Operation {
	return patch(resourceName, nakedType, updateParam, result, "Appliance")
}

// PatchCommonServiceItem Patch操作を定義
func PatchCommonServiceItem(resourceName string, nakedType meta.Type, updateParam, result *dsl.Model) *dsl.Operation {
	return patch(resourceName, nakedType, updateParam, result, "CommonServiceItem")
}

// Delete Delete操作を定義
func Delete(resourceName string) *dsl.Operation {
	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "Delete",
		PathFormat:   dsl.DefaultPathFormatWithID,
		Method:       http.MethodDelete,
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
		},
	}
}

// Config Config操作を定義
func Config(resourceName string) *dsl.Operation {
	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "Config",
		PathFormat:   dsl.IDAndSuffixPathFormat("config"),
		Method:       http.MethodPut,
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
		},
	}
}

// Boot リソースに対するBoot操作を定義
func Boot(resourceName string) *dsl.Operation {
	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "Boot",
		PathFormat:   dsl.IDAndSuffixPathFormat("power"),
		Method:       http.MethodPut,
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
		},
	}
}

// Shutdown リソースに対するシャットダウン操作を定義
func Shutdown(resourceName string) *dsl.Operation {
	param := &dsl.Model{
		Name: "ShutdownOption",
		Fields: []*dsl.FieldDesc{
			{
				Name: "Force",
				Type: meta.TypeFlag,
			},
		},
	}
	return &dsl.Operation{
		ResourceName:    resourceName,
		Name:            "Shutdown",
		PathFormat:      dsl.IDAndSuffixPathFormat("power"),
		Method:          http.MethodDelete,
		RequestEnvelope: dsl.RequestEnvelopeFromModel(param),
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
			dsl.PassthroughModelArgument("shutdownOption", param),
		},
	}
}

// Reset リソースに対するリセット操作を定義
func Reset(resourceName string) *dsl.Operation {
	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "Reset",
		PathFormat:   dsl.IDAndSuffixPathFormat("reset"),
		Method:       http.MethodPut,
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
		},
	}
}

// Status ステータス取得操作を定義
func Status(resourceName string, nakedType meta.Type, result *dsl.Model) *dsl.Operation {
	payloadName := names.ResourceFieldName(resourceName, dsl.PayloadForms.Singular)
	return &dsl.Operation{
		ResourceName:     resourceName,
		Name:             "Status",
		UseWrappedResult: true,
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
		},
		PathFormat: dsl.IDAndSuffixPathFormat("status"),
		Method:     http.MethodGet,
		ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: dsl.Results{
			{
				SourceField: payloadName,
				DestField:   "Status",
				IsPlural:    true,
				Model:       result,
			},
		},
	}
}

// HealthStatus ステータス取得操作を定義(シンプル監視)
func HealthStatus(resourceName string, nakedType meta.Type, result *dsl.Model) *dsl.Operation {
	payloadName := names.ResourceFieldName(resourceName, dsl.PayloadForms.Singular)
	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "HealthStatus",
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
		},
		PathFormat: dsl.IDAndSuffixPathFormat("health"),
		Method:     http.MethodGet,
		ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: dsl.Results{
			{
				SourceField: payloadName,
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}

// OpenFTP FTPオープン操作を定義
func OpenFTP(resourceName string, openParam, result *dsl.Model) *dsl.Operation {
	return &dsl.Operation{
		ResourceName:    resourceName,
		Name:            "OpenFTP",
		PathFormat:      dsl.IDAndSuffixPathFormat("ftp"),
		Method:          http.MethodPut,
		RequestEnvelope: dsl.RequestEnvelopeFromModel(openParam),
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
			dsl.PassthroughModelArgument("openOption", openParam),
		},
		ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
			Name: result.Name,
			Type: meta.Static(naked.OpeningFTPServer{}),
		}),
		Results: dsl.Results{
			{
				SourceField: result.Name,
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}

// CloseFTP FTPクローズ操作を定義
func CloseFTP(resourceName string) *dsl.Operation {
	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         "CloseFTP",
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
		},
		PathFormat: dsl.IDAndSuffixPathFormat("ftp"),
		Method:     http.MethodDelete,
	}
}

// WithIDAction ID+αのみを引数にとるシンプルなオペレーションを定義
func WithIDAction(resourceName, opName, method, pathSuffix string, arguments ...*dsl.Argument) *dsl.Operation {
	args := dsl.Arguments{dsl.ArgumentID}
	args = append(args, arguments...)
	return &dsl.Operation{
		ResourceName: resourceName,
		Name:         opName,
		PathFormat:   dsl.IDAndSuffixPathFormat(pathSuffix),
		Method:       method,
		Arguments:    args,
	}
}

// Monitor アクティビティモニタ取得操作を定義
func Monitor(resourceName string, monitorParam, result *dsl.Model) *dsl.Operation {
	return &dsl.Operation{
		ResourceName:    resourceName,
		Name:            "Monitor",
		PathFormat:      dsl.IDAndSuffixPathFormat("monitor"),
		Method:          http.MethodGet,
		RequestEnvelope: dsl.RequestEnvelopeFromModel(monitorParam),
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
			dsl.PassthroughModelArgument("condition", monitorParam),
		},
		ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
			Type: meta.Static(naked.MonitorValues{}),
			Name: "Data",
		}),
		Results: dsl.Results{
			{
				SourceField: "Data",
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}

// MonitorChild アクティビティモニタ取得操作を定義
func MonitorChild(resourceName, funcNameSuffix, childResourceName string, monitorParam, result *dsl.Model) *dsl.Operation {
	return &dsl.Operation{
		ResourceName:    resourceName,
		Name:            "Monitor" + funcNameSuffix,
		PathFormat:      dsl.IDAndSuffixPathFormat(childResourceName + "/monitor"),
		Method:          http.MethodGet,
		RequestEnvelope: dsl.RequestEnvelopeFromModel(monitorParam),
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
			dsl.PassthroughModelArgument("condition", monitorParam),
		},
		ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
			Type: meta.Static(naked.MonitorValues{}),
			Name: "Data",
		}),
		Results: dsl.Results{
			{
				SourceField: "Data",
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}

// MonitorChildBy アプライアンスなどでの内部リソースインデックスを持つアクティビティモニタ取得操作を定義
func MonitorChildBy(resourceName, funcNameSuffix, childResourceName string, monitorParam, result *dsl.Model) *dsl.Operation {
	pathSuffix := childResourceName + "/{{if eq .index 0}}{{.index}}{{end}}/monitor"
	return &dsl.Operation{
		ResourceName:    resourceName,
		Name:            "Monitor" + funcNameSuffix,
		PathFormat:      dsl.IDAndSuffixPathFormat(pathSuffix),
		Method:          http.MethodGet,
		RequestEnvelope: dsl.RequestEnvelopeFromModel(monitorParam),
		Arguments: dsl.Arguments{
			dsl.ArgumentID,
			{
				Name: "index",
				Type: meta.TypeInt,
			},
			dsl.PassthroughModelArgument("condition", monitorParam),
		},
		ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
			Type: meta.Static(naked.MonitorValues{}),
			Name: "Data",
		}),
		Results: dsl.Results{
			{
				SourceField: "Data",
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}
