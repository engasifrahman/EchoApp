package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

var Validate = validator.New()

func RegisterAlphaSpacesNumbersValidation(v *validator.Validate) error {
	return v.RegisterValidation("alpha_spaces_numbers", func(fl validator.FieldLevel) bool {
		// Allows letters, numbers, and spaces (like Hotel Sheraton 99)
		re := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)
		return re.MatchString(fl.Field().String())
	})
}

func FormatValidationError(err error) map[string][]string {
	res := make(map[string][]string)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			field := strings.ToLower(e.Field())

			var message string
			switch e.Tag() {
			case "required":
				message = fmt.Sprintf("The %s field is required.", field)
			case "min":
				message = fmt.Sprintf("The %s must be at least %s characters.", field, e.Param())
			case "max":
				message = fmt.Sprintf("The %s may not be greater than %s characters.", field, e.Param())
			case "email":
				message = fmt.Sprintf("The %s must be a valid email address.", field)
			case "len":
				message = fmt.Sprintf("The %s must be %s characters long.", field, e.Param())
			case "numeric":
				message = fmt.Sprintf("The %s must be a number.", field)
			case "gte":
				message = fmt.Sprintf("The %s must be greater than or equal to %s.", field, e.Param())
			case "lte":
				message = fmt.Sprintf("The %s must be less than or equal to %s.", field, e.Param())
			case "alpha":
				message = fmt.Sprintf("The %s may only contain letters.", field)
			case "alphanum":
				message = fmt.Sprintf("The %s may only contain letters and numbers.", field)
			case "url":
				message = fmt.Sprintf("The %s must be a valid URL.", field)
			case "uuid":
				message = fmt.Sprintf("The %s must be a valid UUID.", field)
			default:
				message = fmt.Sprintf("The %s is invalid.", field)
			}

			res[field] = append(res[field], message)
		}
	}
	return res
}
