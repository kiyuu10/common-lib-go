package types

import "net/url"

func urlQueryToMap(query url.Values) map[string]string {
	valueMap := make(map[string]string, len(query))
	for key, values := range query {
		valueMap[key] = values[0]
	}
	return valueMap
}
