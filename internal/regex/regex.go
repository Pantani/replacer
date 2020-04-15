package regex

import (
	"regexp"
)

const (
	REGEX = "[A-Za-z0-9]+"
)

func IsValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	reg := regexp.MustCompile(REGEX)
	result := reg.ReplaceAllString(name, "")
	return len(result) == 0
}
