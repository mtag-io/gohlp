package _map

import (
	"github/mtag-io/gohlp/slice"
	"strings"
)

func MGet[T any](m map[string]interface{}, key string) T {
	var none T
	if strings.Contains(key, ".") {
		allKeys := strings.Split(key, ".")
		keys, lastKey := slice.Pop(allKeys)

		for _, k := range keys {
			if val, ok := m[k]; ok {
				m, _ = val.(map[string]interface{})
			} else {
				return none
			}
		}
		return m[lastKey].(T)
	}
	return m[key].(T)
}
