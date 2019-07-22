package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud/v2/internal/define"
	"github.com/sacloud/libsacloud/v2/internal/tools"
)

const destination = "sacloud/search/keys/zz_search_keys.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-search-keys: ")
}

func main() {
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), destination),
		Template:   tmpl,
		Parameter:  define.APIs,
	})
	log.Printf("generated: %s\n", filepath.Join(destination))
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-search-keys'; DO NOT EDIT

package keys

{{ range . }} {{ $typeName := .TypeName }} {{ $resource := . }}
{{ range .Operations }}{{ if .SearchKeys }}
// {{.MethodName}}{{$typeName}} represents strong-typed filter keys for {{$typeName}}.{{.MethodName}}
var {{.MethodName}}{{$typeName}} = struct {
	{{ range .SearchKeys -}}
	{{.KeyName}} string
	{{ end -}}
}{
	{{ range .SearchKeys -}}
	{{.KeyName}}: "{{.SourceFieldName}}",
	{{ end -}}
}
{{ end }}{{- end -}}
{{ end }}
`
