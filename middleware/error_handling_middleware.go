package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
	"github.com/sidiqPratomo/DJKI-Pengaduan/dto"
)

func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		firstError := c.Errors[0].Err
		if firstError != nil {
			statusCode, errorResponse := checkError(firstError)
			c.AbortWithStatusJSON(statusCode, errorResponse)
		}
	}
}

func checkError(err error) (int, dto.ErrorResponse) {
	switch e := err.(type) {
	case *apperror.AppError:
		return e.Code, dto.ErrorResponse{Message: e.Message}
	case validator.ValidationErrors:
		details := generateValidationErrors(e)
		return http.StatusBadRequest, dto.ErrorResponse{Message: appconstant.MsgBadRequest, Details: details}
	default:
		return http.StatusInternalServerError, dto.ErrorResponse{Message: appconstant.MsgInternalServerError}
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	case "lte":
		return "should be less than or equal to " + fe.Param()
	case "gte":
		return "should be greater than or equal to " + fe.Param()
	case "max":
		return "should be max " + fe.Param() + " characters"
	case "email":
		return "should be in valid email format"
	case "ltefield":
		return "should be less than field " + fe.Param()
	case "ValidPassword":
		return "should be minimum 8 characters, contains at least 1 uppercase letter, and contains 1 number"
	case "latitude":
		return "should be in latitude format"
	case "longitude":
		return "should be in longitude format"
	case "number":
		return "should be a number"
	case "datetime":
		return "should be in valid date format yyyy-mm-dd"
	}
	return "unknown error"
}

func generateValidationErrors(ve validator.ValidationErrors) []dto.ValidationErrorDetails {
	details := make([]dto.ValidationErrorDetails, len(ve))
	for i, fe := range ve {
		details[i] = dto.ValidationErrorDetails{Field: fe.Field(), Message: getErrorMsg(fe)}
	}
	return details
}
