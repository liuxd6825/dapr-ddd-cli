package utils

import (
	"errors"
	"strings"
)

func IsEmptyStr(v string) bool {
	if len(v) == 0 {
		return true
	}
	return false
}

func ValidEmptyStr(v string, msg string) error {
	if IsEmptyStr(v) {
		return errors.New(msg)
	}
	return nil
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
