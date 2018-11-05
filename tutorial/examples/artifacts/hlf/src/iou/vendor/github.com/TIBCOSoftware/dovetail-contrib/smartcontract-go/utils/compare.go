package utils

import (
	"strings"
)

func StringCompare(expected, actual string) int {
	return strings.Compare(expected, actual)
}

func IntCompare(expected, actual int) int {
	return expected - actual
}
