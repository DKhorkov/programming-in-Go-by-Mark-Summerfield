package chapter_4

import (
	"strings"
)

func RunTask4(data []string) map[string]map[string]string {
	result := make(map[string]map[string]string)
	var groupName string
	for _, v := range data {
		v = strings.TrimSpace(v)
		if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
			groupName = v[1 : len(v)-1]
			result[groupName] = make(map[string]string)
			continue
		}

		if len(v) > 0 && !strings.HasPrefix(v, ";") {
			split := strings.SplitN(v, "=", 2)
			subKey, subValue := split[0], split[1]
			result[groupName][subKey] = subValue
		}
	}

	return result
}
