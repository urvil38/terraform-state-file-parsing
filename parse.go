package main

import "encoding/json"

func parseInstance[T any](s []any) T {
	var si []T
	v, _ := json.Marshal(s)
	json.Unmarshal(v, &si)
	if len(si) > 0 {
		return si[0]
	}
	var zero T
	return zero
}
