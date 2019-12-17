// Copyright 2016-2019 The Libsacloud Authors
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

package dsl

import (
	"fmt"
	"strings"

	"github.com/huandu/xstrings"
)

func uniqStrings(ss []string) []string {
	seen := make(map[string]struct{}, len(ss))
	i := 0
	for _, v := range ss {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		ss[i] = v
		i++
	}
	return ss[:i]
}

func wrapByDoubleQuote(targets ...string) []string {
	var ss []string
	for _, s := range targets {
		ss = append(ss, fmt.Sprintf(`"%s"`, s))
	}
	return ss
}

func toSnakeCaseName(name string) string {
	return strings.Replace(xstrings.ToSnakeCase(normalizeResourceName(name)), "-", "_", -1)
}

var normalizationWords = map[string]string{
	"IP": "ip",
}

func normalizeResourceName(name string) string {
	n := name
	for k, v := range normalizationWords {
		if strings.HasPrefix(name, k) {
			n = strings.Replace(name, k, v, -1)
			break
		}
	}
	return n
}

func firstRuneToLower(name string) string {
	return xstrings.FirstRuneToLower(name)
}
