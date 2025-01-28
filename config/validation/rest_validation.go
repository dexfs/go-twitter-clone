package validation

import (
	"encoding/json"
	"errors"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"net/http"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func RestError(validation_err error) *rest_errors.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_errors.NewBadRequestError("Invalid field type")
	}
	if errors.As(validation_err, &jsonValidationError) {
		errorsCause := []rest_errors.Cause{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_errors.Cause{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCause = append(errorsCause, cause)
		}
		return rest_errors.NewBadRequestValidationError(http.StatusText(http.StatusBadRequest), errorsCause)
	}

	return rest_errors.NewBadRequestError("Error trying to convert fields")
}
