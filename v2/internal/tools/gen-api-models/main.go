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

package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud/v2/internal/define"
	"github.com/sacloud/libsacloud/v2/internal/tools"
)

const destination = "sacloud/zz_models.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-models: ")
}

func main() {
	outputPath := destination
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
		Template:   tmpl,
		Parameter:  define.APIs,
	})
	log.Printf("generated: %s\n", outputPath)
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-models'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatementsForModelDef "github.com/sacloud/libsacloud/v2/helper/validate" "github.com/sacloud/libsacloud/v2/sacloud/accessor" }}
	{{ . }}
{{- end }}
)

{{ range .Models }}

/************************************************* 
* {{.Name}}
*************************************************/

// {{ .Name }} represents API parameter/response structure
type {{ .Name }} struct {
	{{- range .Fields }}
	{{.Name}} {{.TypeName}} {{if .HasTag }}` + "`" + `{{.TagString}}` + "`" + `{{end}}
	{{- end }}
}

// Validate validates by field tags
func (o *{{ .Name}}) Validate() error {
	return validate.Struct(o)
}

// setDefaults implements sacloud.argumentDefaulter 
func (o *{{.Name}}) setDefaults() interface{} {
	return &struct {
	{{- range .Fields }}
	{{.Name}} {{.TypeName}} {{if .HasTag }}` + "`" + `{{.TagString}}` + "`" + `{{end}}
	{{- end }}
	{{- range .ConstFields }}
	{{.Name}} {{.TypeName}} {{if .HasTag }}` + "`" + `{{.TagString}}` + "`" + `{{end}}
	{{- end }}
	} {
	{{- range .Fields }}
	{{.Name}}: o.Get{{.Name}}(),
	{{- end }}
	{{- range .ConstFields }}
	{{.Name}}: {{.Value}},
	{{- end }}
	}
}

{{- $struct := .Name -}}
{{- range .Methods }}
// {{.Name}} {{if .Description}}{{.Description}}{{else}}.{{end}}
func (o *{{ $struct }}) {{ .Name }}({{ range .Arguments }}{{ .ArgName }} {{ .TypeName }},{{ end }}) ({{ range .ResultTypes }}{{.GoTypeSourceCode}},{{end}}) {
	{{ if .ResultTypes }}return {{ end }}accessor.{{if eq .AccessorFuncName ""}}{{.Name}}{{else}}{{.AccessorFuncName}}{{end}}(o,{{ range .Arguments }}{{ .ArgName }},{{ end }})
}
{{- end }}

{{- range .Fields }} {{ $name := .Name }}{{ $typeName := .TypeName }}
// Get{{$name}} returns value of {{$name}} 
func (o *{{ $struct }}) Get{{$name}}() {{$typeName}} {
	{{ if .DefaultValue -}}
	{{ if eq .Type.GoType "time.Time" -}}
	if o.{{$name}}.IsZero() {
		return {{.DefaultValue}}
	}
	{{ else -}}
	if o.{{$name}} == {{.Type.ZeroInitializeSourceCode}}{
		return {{.DefaultValue}}
	}
	{{ end -}}
	{{ end -}}
	return o.{{$name}}
}

// Set{{$name}} sets value to {{$name}} 
func (o *{{ $struct }}) Set{{$name}}(v {{$typeName}}) {
	o.{{$name}} = v
}

{{- range .Methods }}
// {{.Name}} {{if .Description}}{{.Description}}{{else}}.{{end}}
func (o *{{ $struct }}) {{ .Name }}({{ range .Arguments }}{{ .ArgName }} {{ .TypeName }},{{ end }}) ({{ range .ResultTypes }}{{.GoTypeSourceCode}},{{end}}) {
	{{ if .ResultTypes }}return {{ end }}accessor.{{if eq .AccessorFuncName ""}}{{.Name}}{{else}}{{.AccessorFuncName}}{{end}}(o,{{ range .Arguments }}{{ .ArgName }},{{ end }})
}
{{- end }}
{{- end }} {{/* end of range .Fields */}}

{{- end }} {{/* end of range .Models */}}
`
