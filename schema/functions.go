package schema

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
