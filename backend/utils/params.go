package utils

import (
	"net/url"
	"strconv"
)

// a helper to parse int params
func GetParams(param url.Values, key string, minValue int) int {

	value := param.Get(key)

	val, err := strconv.Atoi(value)

	if err != nil {
		return 0
	}

	return max(minValue, val)
}