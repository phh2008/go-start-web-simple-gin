package util

import (
	"regexp"
	"strings"
)

var snakeReg = regexp.MustCompile("[A-Z][a-z]")

const underline = "_"

// SnakeCase 驼峰转下划线
func SnakeCase(src string) string {
	str := snakeReg.ReplaceAllStringFunc(src, func(s string) string {
		return underline + s
	})
	return strings.ToLower(strings.TrimLeft(str, underline))
}
