package validator

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterTagName used to replace the field name with json tag for the error message
func RegisterTagName() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get(JSON)
		if name == Underscore || name == EmptyString {
			return fld.Name
		}

		return name
	})
}

// RegisterCustomValidation use add custom validator
func RegisterCustomValidation(validate *validator.Validate) {
	validate.RegisterValidation(alpha, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), alphaRegex)
	})

	validate.RegisterValidation(alphaNumeric, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), alphaNumericRegex)
	})

	validate.RegisterValidation(email, func(fl validator.FieldLevel) bool {
		return validateEmail(fl.Field().String())
	})

	validate.RegisterValidation(timestamp, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), timestampRegex)
	})

	validate.RegisterValidation(alphaNumericWithHyphenSpace, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), alphaNumericWithHyphenSpaceRegex)
	})

	validate.RegisterValidation(alphaNumericWithHyphen, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), alphaNumericWithHyphenRegex)
	})

}

// RegisterCustomTranslation use add custom validator translation
func RegisterCustomTranslation(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(alpha, trans, func(ut ut.Translator) error {
		return ut.Add(alpha, "{0} must contain alpha characters", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(alpha, fe.Field())
		return t
	})

	validate.RegisterTranslation(email, trans, func(ut ut.Translator) error {
		return ut.Add(email, "{0} must be a valid email address", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(email, fe.Field())
		return t
	})

	validate.RegisterTranslation(timestamp, trans, func(ut ut.Translator) error {
		return ut.Add(timestamp, "{0} must be a valid UTC timestamp", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(timestamp, fe.Field())
		return t
	})

	validate.RegisterTranslation(alphaNumericWithHyphenSpace, trans, func(ut ut.Translator) error {
		return ut.Add(alphaNumericWithHyphenSpace, "{0} must contain alpha-numaric characters with hyphen and space", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(alphaNumericWithHyphenSpace, fe.Field())
		return t
	})

	validate.RegisterTranslation(alphaNumericWithHyphen, trans, func(ut ut.Translator) error {
		return ut.Add(alphaNumericWithHyphen, "{0} must contain alpha-numaric characters with hyphen", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(alphaNumericWithHyphen, fe.Field())
		return t
	})

	validate.RegisterTranslation(timeonly, trans, func(ut ut.Translator) error {
		return ut.Add(timeonly, "{0} must be a valid time in the format 'HH:MM:SS'", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(timeonly, fe.Field())
		return t
	})
}
