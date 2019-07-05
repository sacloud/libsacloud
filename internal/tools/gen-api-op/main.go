package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud/v2/internal/define"
	"github.com/sacloud/libsacloud/v2/internal/tools"
)

const destination = "sacloud/zz_api_ops.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-op: ")
}

func main() {
	outputPath := destination
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
		Template:   tmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", outputPath)
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-op'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "context" "encoding/json" "github.com/sacloud/libsacloud/v2/sacloud/naked" "github.com/sacloud/libsacloud/v2/pkg/mapconv" }}
	{{ . }}
{{- end }}
)

func init() {
{{ range . }}
	SetClientFactoryFunc("{{.TypeName}}", func(caller APICaller) interface{} {
		return &{{ .TypeName }}Op {
			Client: caller,
			PathSuffix: "{{.GetPathSuffix}}",
			PathName: "{{.GetPathName}}",
		}
	})
{{ end -}}
}

{{ range . }}{{ $typeName := .TypeName}}

/************************************************* 
* {{$typeName}}Op
*************************************************/

// {{ .TypeName }}Op implements {{ .TypeName }}API interface
type {{ .TypeName }}Op struct{
	// Client APICaller
    Client APICaller
	// PathSuffix is used when building URL
	PathSuffix string
	// PathName is used when building URL
	PathName string
}

// New{{ $typeName}}Op creates new {{ $typeName}}Op instance
func New{{ $typeName}}Op(caller APICaller) {{ $typeName}}API {
	return GetClientFactoryFunc("{{$typeName}}")(caller).({{$typeName}}API)
}

{{ range .Operations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
// {{ .MethodName }} is API call
func (o *{{ $typeName }}Op) {{ .MethodName }}(ctx context.Context{{ range .AllArguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
	url, err := buildURL("{{.GetPathFormat}}", map[string]interface{}{
		"rootURL": SakuraCloudAPIRoot,
		"pathSuffix": o.PathSuffix,
		"pathName": o.PathName,
		{{- range .AllArguments }}
		"{{.Name}}": {{.Name}},
		{{- end }}
	})
	if err != nil {
		return {{ $returnErrStatement }}
	}

	var body interface{}
{{ if .HasRequestEnvelope }}
	{{- range .AllArguments }}
	if {{.ArgName}} == {{.ZeroValueOnSource}} {
		{{.ArgName}} = {{.ZeroInitializer}}	
	}
	{{- end }}
	args := &struct {
		{{- range .AllArguments }}
		Arg{{ .ArgName }} {{ .TypeName }} {{.MapConvTagSrc}}
		{{- end }}
	}{
		{{- range .AllArguments }}
		Arg{{ .ArgName }}:{{ .ArgName}},
		{{- end }}
	}

	v := &{{.RequestEnvelopeStructName}}{}
	if err := mapconv.ConvertTo(args, v); err != nil {
		return {{ $returnErrStatement }}
	}
	body = v
{{ end }}

	{{ if .HasResponseEnvelope -}}
	data, err := o.Client.Do(ctx, "{{.GetMethod}}", url, body)
	{{ else -}}
	_, err = o.Client.Do(ctx, "{{.GetMethod}}", url, body)
	{{ end -}}
	if err != nil {
		return {{ $returnErrStatement }}
	}

	{{ if .HasResponseEnvelope -}}
	nakedResponse := &{{.ResponseEnvelopeStructName}}{}
	if err := json.Unmarshal(data, nakedResponse); err != nil {
		return {{ $returnErrStatement }}
	}

	{{ if .IsResponseSingular -}}
	{{ range $i,$v := .AllResults -}}
	payload{{$i}} := {{$v.ZeroInitializeSourceCode}}
	if err := payload{{$i}}.convertFrom(nakedResponse.{{.DestField}}); err != nil {
		return {{ $returnErrStatement }}
	}
	{{ end -}}
	{{- else if .IsResponsePlural -}}
	{{ range $i,$v := .AllResults -}}
	var payload{{$i}} []{{$v.GoTypeSourceCode}}
	for _ , v := range nakedResponse.{{.DestField}} {
		payload := {{$v.ZeroInitializeSourceCode}}
		if err := payload.convertFrom(v); err != nil {
			return {{ $returnErrStatement }}
		}
		payload{{$i}} = append(payload{{$i}}, payload)
	}
	{{ end -}}
	{{ end -}}
	{{ end -}}

	return {{range $i,$v := .AllResults}}payload{{$i}},{{ end }} nil
}
{{ end -}}
{{ end -}}
`
