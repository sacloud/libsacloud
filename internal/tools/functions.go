// Copyright 2016-2020 The Libsacloud Authors
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

package tools

import (
	"bytes"
	"fmt"
	"go/build"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/huandu/xstrings"
)

type TemplateConfig struct {
	OutputPath         string
	Template           string
	Parameter          interface{}
	PreventOverwriting bool
}

func WriteFileWithTemplate(config *TemplateConfig) bool {
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(config.Template))
	if err := t.Execute(buf, config.Parameter); err != nil {
		log.Fatalf("writing output: %s", err)
	}

	if config.PreventOverwriting {
		if _, err := os.Stat(config.OutputPath); err == nil {
			return false
		}
	}

	// write to file
	if err := ioutil.WriteFile(config.OutputPath, Sformat(buf.Bytes()), 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
	return true
}

// Gopath returns GOPATH
func Gopath() string {
	gopath := build.Default.GOPATH
	gopath = filepath.SplitList(gopath)[0]
	return gopath
}

func ProjectRootPath() string {
	return filepath.Join(Gopath(), "src/github.com/sacloud/libsacloud")
}

// Sformat formats go source codes
func Sformat(buf []byte) []byte {
	src, err := format.Source(buf)
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		log.Printf("generated: \n%s", string(buf))
		return buf
	}
	return src
}

var normalizationWords = map[string]string{
	"IP": "ip",
}

// NormalizeResourceName
func NormalizeResourceName(name string) string {
	n := name
	for k, v := range normalizationWords {
		if strings.HasPrefix(name, k) {
			n = strings.Replace(name, k, v, -1)
			break
		}
	}
	return n
}

func ToSnakeCaseName(name string) string {
	return strings.Replace(xstrings.ToSnakeCase(NormalizeResourceName(name)), "-", "_", -1)
}

func ToDashedName(name string) string {
	// From "CamelCase" to "dash-case"
	return strings.Replace(xstrings.ToSnakeCase(NormalizeResourceName(name)), "_", "-", -1)
}

func ToCamelCaseName(name string) string {
	return xstrings.ToCamelCase(strings.Replace(NormalizeResourceName(name), "-", "_", -1))
}

func ToCamelWithFirstLower(name string) string {
	return xstrings.FirstRuneToLower(xstrings.ToCamelCase(strings.Replace(NormalizeResourceName(name), "-", "_", -1)))
}

func ToCLIFlagName(name string) string {
	format := "--%s"
	if len(name) == 1 {
		format = "-%s"
	}
	return fmt.Sprintf(format, ToDashedName(name))
}

func FlattenStringList(list []string) string {
	if len(list) > 0 {
		return fmt.Sprintf("\"%s\"", strings.Join(list, "\",\""))
	}
	return ""
}

func FlattenIntList(list []int) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%d", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenUintList(list []uint) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%d", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}
func FlattenInt64List(list []int64) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%d", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenUint64List(list []uint64) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%d", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenFloatList(list []float32) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%f", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenFloat64List(list []float64) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%f", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}
