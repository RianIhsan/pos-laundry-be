package utils

import "strconv"

func ConvertStringToUint(s string) (uint64, error) {
	result, err := strconv.ParseUint(s, 10, 64)
	return result, err
}
