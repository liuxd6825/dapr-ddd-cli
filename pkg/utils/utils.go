package utils

import "errors"

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
