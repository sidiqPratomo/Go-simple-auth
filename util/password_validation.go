package util

import (
	"unicode"
)

func ValidatePassword(password string) bool {
	min := 8
	if len(password) < min {
		return false
	}

	var number bool
	var upper bool
	var lower bool

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			lower = true
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsNumber(char):
			number = true
		}
		if number && upper && lower {
			return true
		}
	}
	return false
}