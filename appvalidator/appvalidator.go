package appvalidator

import (
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func AppValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ValidPassword", ValidPassword)
	}
}

func ValidPassword(fl validator.FieldLevel) bool {
	if password, ok := fl.Field().Interface().(string); ok {
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
	}
	return false
}
