package chapter_4

import (
	"fmt"
	"sort"
)

func RunTask5(data map[string]map[string]string) {
	var groups []string
	for group := range data {
		groups = append(groups, group)
	}

	sort.Strings(groups)
	for _, group := range groups {
		fmt.Printf("[%s]\n", group)

		var keys []string
		for key := range data[group] {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			fmt.Printf("%s=%s\n", key, data[group][key])
		}
	}
}
