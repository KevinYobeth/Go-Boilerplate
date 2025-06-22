package validator

import (
	"reflect"

	"github.com/go-playground/locales/en"
	"github.com/leebenson/conform"
	"github.com/samber/lo"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateStruct(s any) error {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.NewIncorrectInputError(nil, "input must be a non-nil pointer to struct")
	}

	conformErr := conform.Strings(s)
	if conformErr != nil {
		return errors.NewGenericError(conformErr, "failed to conform struct")
	}

	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		field, ok := fld.Tag.Lookup("field")
		if ok {
			return field
		}

		json, ok := fld.Tag.Lookup("json")
		if ok {
			return json
		}

		return lo.SnakeCase(fld.Name)
	})

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return errors.NewValidatorError(validationErrors.Translate(trans))
	}
	return errors.NewGenericError(err, "unexpected error during validation")
}
