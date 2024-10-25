package util

import (
	"regexp"
)

func RegexValidate(str string, rgx string) bool {
	regex, err := regexp.Compile(rgx)
	if err != nil {
		return false
	}

	if !regex.Match([]byte(str)) {
		return false
	}

	return true
}
