package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud-v2/internal/define"
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/tools"
)

const (
	apisDestination = "sacloud/fake/zz_api_ops.go"
	opsDestination  = "sacloud/fake/ops_%s.go"
	testDestination = "sacloud/fake/zz_api_ops_test.go"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-fake-op: ")
}

func main() {
	// generate xxxOp
	outputPath := apisDestination
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
		Template:   apisTmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", outputPath)

	// generate funcs
	schema.IsOutOfSacloudPackage = true
	for _, resource := range define.Resources {
		dest := fmt.Sprintf(opsDestination, resource.FileSafeName())
		wrote := tools.WriteFileWithTemplate(&tools.TemplateConfig{
			OutputPath:         filepath.Join(tools.ProjectRootPath(), dest),
			Template:           opsTmpl,
			Parameter:          resource,
			PreventOverwriting: true,
		})
		if wrote {
			log.Printf("generated: %s\n", filepath.Join(dest))
		}
	}

	// generate xxxOp
	outputPath = testDestination
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
		Template:   testTmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", outputPath)

}

const apisTmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-fake-op'; DO NOT EDIT

package fake

import (
{{- range .ImportStatements "github.com/sacloud/libsacloud-v2/sacloud" "github.com/sacloud/libsacloud-v2/sacloud/types"}}
	{{ . }}
{{- end }}
)

// SwitchFactoryFuncToFake switches sacloud.xxxAPI's factory methods to use fake client
func SwitchFactoryFuncToFake() {
{{ range . -}}
	sacloud.SetClientFactoryFunc(Resource{{.TypeName}}, func(caller sacloud.APICaller) interface{} {
		return New{{ .TypeName }}Op()
	})
{{ end -}}
}


{{ range . }}{{ $typeName := .TypeName}}

/************************************************* 
* {{$typeName}}Op
*************************************************/

// {{ .TypeName }}Op is fake implementation of {{ .TypeName }}API interface
type {{ .TypeName }}Op struct{
	key string
}

// New{{ $typeName}}Op creates new {{ $typeName}}Op instance
func New{{ $typeName}}Op() sacloud.{{ $typeName}}API {
	return &{{$typeName}}Op {
		key: Resource{{$typeName}},
	}
}
{{ end -}}
`

const opsTmpl = `package fake

import (
{{- range .ImportStatements "context" "github.com/sacloud/libsacloud-v2/sacloud" "github.com/sacloud/libsacloud-v2/sacloud/types"}}
	{{ . }}
{{- end }}
)

{{ range .Operations }}
// {{ .MethodName }} is fake implementation
func (o *{{ .ResourceTypeName }}Op) {{ .MethodName }}(ctx context.Context{{ range .AllArguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
{{ if eq .MethodName "Find" -}}
	results, _ := find(o.key, {{if .ResourceIsGlobal}}sacloud.DefaultZone{{else}}zone{{end}}, conditions)
	var values []*sacloud.{{.ResourceTypeName}}
	for _, res := range results {
		dest := &sacloud.{{.ResourceTypeName}}{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return values, nil
{{ else if eq .MethodName "Create" -}}
	result := &sacloud.{{.ResourceTypeName}}{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	// TODO core logic is not implemented

	s.set{{.ResourceTypeName}}({{if .ResourceIsGlobal}}sacloud.DefaultZone{{else}}zone{{end}}, result)
	return result, nil
{{ else if eq .MethodName "Read" -}}
	value := s.get{{.ResourceTypeName}}ByID({{if .ResourceIsGlobal}}sacloud.DefaultZone{{else}}zone{{end}}, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.{{.ResourceTypeName}}{}
	copySameNameField(value, dest)
	return dest, nil
{{ else if eq .MethodName "Update" -}}
	value, err := o.Read(ctx, {{if .ResourceIsGlobal}}sacloud.DefaultZone{{else}}zone{{end}}, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	// TODO core logic is not implemented

	return value, nil
{{ else if eq .MethodName "Delete" -}}
	_, err := o.Read(ctx, {{if .ResourceIsGlobal}}sacloud.DefaultZone{{else}}zone{{end}}, id)
	if err != nil {
		return err
	}

	// TODO core logic is not implemented

	s.delete(o.key, {{if .ResourceIsGlobal}}sacloud.DefaultZone{{else}}zone{{end}}, id)
	return nil
{{ else -}}
	// TODO not implemented
	err := errors.New("not implements")
	return {{.ReturnErrorStatement}}
{{ end -}}
}
{{ end }}
`

const testTmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-fake-op'; DO NOT EDIT

package fake

import (
{{- range .ImportStatements "testing" "github.com/sacloud/libsacloud-v2/sacloud" "github.com/sacloud/libsacloud-v2/sacloud/types"}}
        {{ . }}
{{- end }}
)

func TestResourceOps(t *testing.T) {
{{ range . }}
        if op, ok := New{{.TypeName}}Op().(sacloud.{{.TypeName}}API); !ok {
                t.Fatalf("%s is not sacloud.{{.TypeName}}", op)
        }
{{ end }}
}`