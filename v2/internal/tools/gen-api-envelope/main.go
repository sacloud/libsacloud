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

const destination = "sacloud/zz_envelopes.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-envelope: ")
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

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-envelope'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "github.com/sacloud/libsacloud/v2/sacloud/types" "github.com/sacloud/libsacloud/v2/sacloud/naked" "github.com/sacloud/libsacloud/v2/sacloud/search" }}
	{{ . }}
{{- end }}
)

{{- range . }}
{{- range .Operations -}}

{{ if .HasRequestEnvelope }}
// {{ .RequestEnvelopeStructName }} is envelop of API request
type {{ .RequestEnvelopeStructName }} struct {
{{ if .IsRequestSingular }}
	{{- range .RequestPayloads}}
	{{.Name}} {{.TypeName}} {{.TagString}}
	{{- end }}
{{- else if .IsRequestPlural -}}
	{{- range .RequestPayloads}}
	{{.Name}} []{{.TypeName}} {{.TagString}}
	{{- end }}
{{ end }}
}
{{ end }}

{{ if .HasResponseEnvelope }}
// {{ .ResponseEnvelopeStructName }} is envelop of API response
type {{ .ResponseEnvelopeStructName }} struct {
{{- if .IsResponsePlural -}}
	Total       int        ` + "`" + `json:",omitempty"` + "`" + ` // トータル件数
	From        int        ` + "`" + `json:",omitempty"` + "`" + ` // ページング開始ページ
	Count       int        ` + "`" + `json:",omitempty"` + "`" + ` // 件数
{{ else }}
	IsOk    bool  ` + "`" + `json:"is_ok,omitempty"` + "`" + ` // is_ok項目
	Success types.APIResult  ` + "`" + `json:",omitempty"` + "`" + `      // success項目
{{ end }}
{{ if .IsResponseSingular }}
	{{- range .ResponsePayloads}}
	{{.Name}} {{.TypeName}} {{.TagString}}
	{{- end }}
{{- else if .IsResponsePlural -}}
	{{- range .ResponsePayloads}}
	{{.Name}} []{{.TypeName}} {{.TagString}}
	{{- end }}
{{ end }}
}
{{ end }}

{{- end -}}
{{- end -}}
`
