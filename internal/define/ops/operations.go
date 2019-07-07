package ops

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

func find(resourceName string, nakedType meta.Type, findParam, result *schema.Model, payloadName string) *schema.Operation {
	if payloadName == "" {
		payloadName = names.ResourceFieldName(resourceName, schema.PayloadForms.Plural)
	}

	return &schema.Operation{
		ResourceName:    resourceName,
		Name:            "Find",
		PathFormat:      schema.DefaultPathFormat,
		Method:          http.MethodGet,
		RequestEnvelope: schema.RequestEnvelopeFromModel(findParam),
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.PassthroughModelArgument("conditions", findParam),
		},
		ResponseEnvelope: schema.ResponseEnvelopePlural(&schema.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: schema.Results{
			{
				SourceField: payloadName,
				DestField:   names.ResourceFieldName(resourceName, schema.PayloadForms.Plural),
				IsPlural:    true,
				Model:       result,
			},
		},
	}
}

// Find Find操作を定義
func Find(resourceName string, nakedType meta.Type, findParam, result *schema.Model) *schema.Operation {
	return find(resourceName, nakedType, findParam, result, "")
}

// FindAppliance Find操作を定義
func FindAppliance(resourceName string, nakedType meta.Type, findParam, result *schema.Model) *schema.Operation {
	return find(resourceName, nakedType, findParam, result, "Appliances")
}

// FindCommonServiceItem Find操作を定義
func FindCommonServiceItem(resourceName string, nakedType meta.Type, findParam, result *schema.Model) *schema.Operation {
	return find(resourceName, nakedType, findParam, result, "CommonServiceItems")
}

func create(resourceName string, nakedType meta.Type, createParam, result *schema.Model, payloadName string) *schema.Operation {
	if payloadName == "" {
		payloadName = names.ResourceFieldName(resourceName, schema.PayloadForms.Singular)
	}

	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "Create",
		PathFormat:   schema.DefaultPathFormat,
		Method:       http.MethodPost,
		RequestEnvelope: schema.RequestEnvelope(&schema.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.MappableArgument("param", createParam, payloadName),
		},
		ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: schema.Results{
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
func Create(resourceName string, nakedType meta.Type, createParam, result *schema.Model) *schema.Operation {
	return create(resourceName, nakedType, createParam, result, "")
}

// CreateAppliance Create操作を定義
func CreateAppliance(resourceName string, nakedType meta.Type, createParam, result *schema.Model) *schema.Operation {
	return create(resourceName, nakedType, createParam, result, "Appliance")
}

// CreateCommonServiceItem Create操作を定義
func CreateCommonServiceItem(resourceName string, nakedType meta.Type, createParam, result *schema.Model) *schema.Operation {
	return create(resourceName, nakedType, createParam, result, "CommonServiceItem")
}

func read(resourceName string, nakedType meta.Type, result *schema.Model, payloadName string) *schema.Operation {
	if payloadName == "" {
		payloadName = names.ResourceFieldName(resourceName, schema.PayloadForms.Singular)
	}

	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "Read",
		PathFormat:   schema.DefaultPathFormatWithID,
		Method:       http.MethodGet,
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
		},
		ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: schema.Results{
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
func Read(resourceName string, nakedType meta.Type, result *schema.Model) *schema.Operation {
	return read(resourceName, nakedType, result, "")
}

// ReadAppliance Read操作を定義
func ReadAppliance(resourceName string, nakedType meta.Type, result *schema.Model) *schema.Operation {
	return read(resourceName, nakedType, result, "Appliance")
}

// ReadCommonServiceItem Read操作を定義
func ReadCommonServiceItem(resourceName string, nakedType meta.Type, result *schema.Model) *schema.Operation {
	return read(resourceName, nakedType, result, "CommonServiceItem")
}

func update(resourceName string, nakedType meta.Type, updateParam, result *schema.Model, payloadName string) *schema.Operation {
	if payloadName == "" {
		payloadName = names.ResourceFieldName(resourceName, schema.PayloadForms.Singular)
	}

	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "Update",
		PathFormat:   schema.DefaultPathFormatWithID,
		Method:       http.MethodPut,
		RequestEnvelope: schema.RequestEnvelope(&schema.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
			schema.MappableArgument("param", updateParam, payloadName),
		},
		ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: schema.Results{
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
func Update(resourceName string, nakedType meta.Type, updateParam, result *schema.Model) *schema.Operation {
	return update(resourceName, nakedType, updateParam, result, "")
}

// UpdateAppliance Update操作を定義
func UpdateAppliance(resourceName string, nakedType meta.Type, updateParam, result *schema.Model) *schema.Operation {
	return update(resourceName, nakedType, updateParam, result, "Appliance")
}

// UpdateCommonServiceItem Update操作を定義
func UpdateCommonServiceItem(resourceName string, nakedType meta.Type, updateParam, result *schema.Model) *schema.Operation {
	return update(resourceName, nakedType, updateParam, result, "CommonServiceItem")
}

// Delete Delete操作を定義
func Delete(resourceName string) *schema.Operation {
	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "Delete",
		PathFormat:   schema.DefaultPathFormatWithID,
		Method:       http.MethodDelete,
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
		},
	}
}

// Config Config操作を定義
func Config(resourceName string) *schema.Operation {
	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "Config",
		PathFormat:   schema.IDAndSuffixPathFormat("config"),
		Method:       http.MethodPut,
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
		},
	}
}

// Boot リソースに対するBoot操作を定義
func Boot(resourceName string) *schema.Operation {
	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "Boot",
		PathFormat:   schema.IDAndSuffixPathFormat("power"),
		Method:       http.MethodPut,
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
		},
	}
}

// Shutdown リソースに対するシャットダウン操作を定義
func Shutdown(resourceName string) *schema.Operation {
	param := &schema.Model{
		Name: "ShutdownOption",
		Fields: []*schema.FieldDesc{
			{
				Name: "Force",
				Type: meta.TypeFlag,
			},
		},
	}
	return &schema.Operation{
		ResourceName:    resourceName,
		Name:            "Shutdown",
		PathFormat:      schema.IDAndSuffixPathFormat("power"),
		Method:          http.MethodDelete,
		RequestEnvelope: schema.RequestEnvelopeFromModel(param),
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
			schema.PassthroughModelArgument("shutdownOption", param),
		},
	}
}

// Reset リソースに対するリセット操作を定義
func Reset(resourceName string) *schema.Operation {
	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "Reset",
		PathFormat:   schema.IDAndSuffixPathFormat("reset"),
		Method:       http.MethodPut,
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
		},
	}
}

// Status ステータス取得操作を定義
func Status(resourceName string, nakedType meta.Type, result *schema.Model) *schema.Operation {
	payloadName := names.ResourceFieldName(resourceName, schema.PayloadForms.Singular)
	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "Status",
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
		},
		PathFormat: schema.IDAndSuffixPathFormat("status"),
		Method:     http.MethodGet,
		ResponseEnvelope: schema.ResponseEnvelopePlural(&schema.EnvelopePayloadDesc{
			Type: nakedType,
			Name: payloadName,
		}),
		Results: schema.Results{
			{
				SourceField: payloadName,
				DestField:   "Status",
				IsPlural:    true,
				Model:       result,
			},
		},
	}
}

// OpenFTP FTPオープン操作を定義
func OpenFTP(resourceName string, openParam, result *schema.Model) *schema.Operation {
	return &schema.Operation{
		ResourceName:    resourceName,
		Name:            "OpenFTP",
		PathFormat:      schema.IDAndSuffixPathFormat("ftp"),
		Method:          http.MethodPut,
		RequestEnvelope: schema.RequestEnvelopeFromModel(openParam),
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
			schema.PassthroughModelArgument("openOption", openParam),
		},
		ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
			Name: result.Name,
			Type: meta.Static(naked.OpeningFTPServer{}),
		}),
		Results: schema.Results{
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
func CloseFTP(resourceName string) *schema.Operation {
	return &schema.Operation{
		ResourceName: resourceName,
		Name:         "CloseFTP",
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
		},
		PathFormat: schema.IDAndSuffixPathFormat("ftp"),
		Method:     http.MethodDelete,
	}
}

// WithIDAction ID+αのみを引数にとるシンプルなオペレーションを定義
func WithIDAction(resourceName, opName, method, pathSuffix string, arguments ...*schema.Argument) *schema.Operation {
	args := schema.Arguments{schema.ArgumentZone, schema.ArgumentID}
	args = append(args, arguments...)
	return &schema.Operation{
		ResourceName: resourceName,
		Name:         opName,
		PathFormat:   schema.IDAndSuffixPathFormat(pathSuffix),
		Method:       method,
		Arguments:    args,
	}
}

// Monitor アクティビティモニタ取得操作を定義
func Monitor(resourceName string, monitorParam, result *schema.Model) *schema.Operation {
	return &schema.Operation{
		ResourceName:    resourceName,
		Name:            "Monitor",
		PathFormat:      schema.IDAndSuffixPathFormat("monitor"),
		Method:          http.MethodGet,
		RequestEnvelope: schema.RequestEnvelopeFromModel(monitorParam),
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
			schema.PassthroughModelArgument("condition", monitorParam),
		},
		ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
			Type: meta.Static(naked.MonitorValues{}),
			Name: "Data",
		}),
		Results: schema.Results{
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
func MonitorChild(resourceName, funcNameSuffix, childResourceName string, monitorParam, result *schema.Model) *schema.Operation {
	return &schema.Operation{
		ResourceName:    resourceName,
		Name:            "Monitor" + funcNameSuffix,
		PathFormat:      schema.IDAndSuffixPathFormat(childResourceName + "/monitor"),
		Method:          http.MethodGet,
		RequestEnvelope: schema.RequestEnvelopeFromModel(monitorParam),
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
			schema.PassthroughModelArgument("condition", monitorParam),
		},
		ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
			Type: meta.Static(naked.MonitorValues{}),
			Name: "Data",
		}),
		Results: schema.Results{
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
func MonitorChildBy(resourceName, funcNameSuffix, childResourceName string, monitorParam, result *schema.Model) *schema.Operation {
	pathSuffix := childResourceName + "/{{if eq .index 0}}{{.index}}{{end}}/monitor"
	return &schema.Operation{
		ResourceName:    resourceName,
		Name:            "Monitor" + funcNameSuffix,
		PathFormat:      schema.IDAndSuffixPathFormat(pathSuffix),
		Method:          http.MethodGet,
		RequestEnvelope: schema.RequestEnvelopeFromModel(monitorParam),
		Arguments: schema.Arguments{
			schema.ArgumentZone,
			schema.ArgumentID,
			{
				Name: "index",
				Type: meta.TypeInt,
			},
			schema.PassthroughModelArgument("condition", monitorParam),
		},
		ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
			Type: meta.Static(naked.MonitorValues{}),
			Name: "Data",
		}),
		Results: schema.Results{
			{
				SourceField: "Data",
				DestField:   result.Name,
				IsPlural:    false,
				Model:       result,
			},
		},
	}
}
