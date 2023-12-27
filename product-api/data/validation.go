package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidationError wraps the validators FieldError so we do not
// expose this to out code.
type ValidationError struct {
	validator.FieldError
}

// Error wraps to error that human can read.
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error()) // look at line 17
	}

	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate: validate}
}

func (v *Validation) Validate(i any) ValidationErrors {
	errs := v.validate.Struct(i)
	if errs == nil {
		return nil
	}

	if len(errs.(validator.ValidationErrors)) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs.(validator.ValidationErrors) {
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// validateSKU implement validator.Func
func validateSKU(fl validator.FieldLevel) bool {
	// *: This line may be suddenly interrupt app because it can panic
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)

	// find all match string
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
}
