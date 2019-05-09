package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud-v2/internal/define"
	"github.com/sacloud/libsacloud-v2/internal/tools"
)

const destination = "sacloud/zz_%s_op.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-op: ")
}

func main() {
	for _, resource := range define.Resources {
		outputPath := fmt.Sprintf(destination, resource.FileSafeName())
		tools.WriteFileWithTemplate(&tools.TemplateConfig{
			OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
			Template:   tmpl,
			Parameter:  resource,
		})
		log.Printf("generated: %s\n", outputPath)

	}

}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-op'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "context" "encoding/json" "github.com/sacloud/libsacloud-v2/sacloud/naked" }}
	{{ . }}
{{- end }}
)

// {{ .TypeName }}Op implements {{ .TypeName }}API interface
type {{ .TypeName }}Op struct{
	// Client APICaller
    Client APICaller
	// PathSuffix is used when building URL
	PathSuffix string
	// PathName is used when building URL
	PathName string
}

// New{{ .TypeName }}Op creates new {{ .TypeName }}Op instance
func New{{ .TypeName }}Op(client APICaller) *{{ .TypeName }}Op {
	return &{{ .TypeName }}Op {
    	Client: client,
		PathSuffix: "{{$.GetPathSuffix}}",
		PathName: "{{$.GetPathName}}",
	}
}

{{ range .AllOperations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
// {{ .MethodName }} is API call
func (o *{{ $.TypeName }}Op) {{ .MethodName }}(ctx context.Context{{ range .AllArguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
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
	{{- $structName := .RequestEnvelopeStructName -}}
	{{ range .MapDestinationDeciders }} 
	{
		if body == nil {
			body = &{{$structName}}{}
		}
		v := body.(*{{$structName}})
		n, err := {{.ArgName}}.toNaked()
		if err != nil {
			return {{ $returnErrStatement }}
		}
		v.{{.DestinationFieldName}} = n 
		body = v
	}
	{{ end }}
	{{ range .PassthroughFieldDeciders}} 
	{{- $argName := .ArgName -}}
	{
		if body == nil {
			body = &{{$structName}}{}
		}
		v := body.(*{{$structName}})
		{{- range .PassthroughFieldNames }}
		v.{{.}} = {{$argName}}.{{.}}
		{{- end }}
		body = v
	}
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
	if err := payload{{$i}}.parseNaked(nakedResponse.{{.SourceField}}); err != nil {
		return {{ $returnErrStatement }}
	}
	{{ end -}}
	{{- else if .IsResponsePlural -}}
	{{ range $i,$v := .AllResults -}}
	var payload{{$i}} []{{$v.GoTypeSourceCode}}
	for _ , v := range nakedResponse.{{.SourceField}} {
		payload := {{$v.ZeroInitializeSourceCode}}
		if err := payload.parseNaked(v); err != nil {
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
`
