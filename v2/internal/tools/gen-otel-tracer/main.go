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
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/tools"
)

const destination = "sacloud/trace/otel/zz_api_tracer.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-otel-tracer: ")
}

func main() {
	dsl.IsOutOfSacloudPackage = true

	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), destination),
		Template:   tmpl,
		Parameter:  define.APIs,
	})
	log.Printf("generated: %s\n", filepath.Join(destination))
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-otel-tracer'; DO NOT EDIT

package otel

import (
	"github.com/sacloud/libsacloud/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

func addClientFactoryHooks(cnf *config) {
{{ range . -}} 
	sacloud.AddClientFacotyHookFunc("{{.TypeName}}", func(in interface{}) interface{} {
		return new{{.TypeName}}Tracer(in.(sacloud.{{.TypeName}}API), cnf)
	})
{{ end -}}
}

{{ range . }} {{$typeName := .TypeName}} {{ $resource := . }}
/************************************************* 
* {{ $typeName }}Tracer
*************************************************/

// {{ $typeName }}Tracer is for trace {{ $typeName }}Op operations
type {{ $typeName }}Tracer struct {
	Internal sacloud.{{$typeName}}API
	config *config
}

// New{{ $typeName}}Tracer creates new {{ $typeName}}Tracer instance
func new{{ $typeName}}Tracer(in sacloud.{{$typeName}}API, cnf *config) sacloud.{{$typeName}}API {
	return &{{ $typeName}}Tracer {
		Internal: in,
		config: cnf,
	}
}

{{ range .Operations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
// {{ .MethodName }} is API call with trace log
func (t *{{ $typeName }}Tracer) {{ .MethodName }}(ctx context.Context{{if not $resource.IsGlobal}}, zone string{{end}}{{ range .Arguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
	var span trace.Span
	options := append(t.config.SpanStartOptions, trace.WithAttributes(
{{if not $resource.IsGlobal -}}
		label.String("libsacloud.api.arguments.zone", zone),
{{ end -}}
{{ range .Arguments -}}
		label.Any("libsacloud.api.arguments.{{.ArgName}}", {{.ArgName}}),
{{ end -}}
	))
	ctx, span = t.config.Tracer.Start(ctx, "{{ $typeName }}API.{{ .MethodName }}", options...)
	defer func() {
		span.End()
	}()

	// for http trace
	ctx = httptrace.WithClientTrace(ctx, otelhttptrace.NewClientTrace(ctx))
	{{range .ResultsTypeInfo}}{{.VarName}}, {{end}}err := t.Internal.{{ .MethodName }}(ctx{{if not $resource.IsGlobal}}, zone{{end}}{{ range .Arguments }}, {{ .ArgName }}{{ end }})

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}else {
		span.SetStatus(codes.Ok, "")
		{{range .ResultsTypeInfo}}span.SetAttributes(label.Any("libsacloud.api.results.{{.VarName}}", {{.VarName}}))
{{ end }}
	}
	return {{range .ResultsTypeInfo}}{{.VarName}}, {{end}}err
}
{{- end -}}

{{ end }}
`
