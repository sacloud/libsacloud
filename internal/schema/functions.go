package schema

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
