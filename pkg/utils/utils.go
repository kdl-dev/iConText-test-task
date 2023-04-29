package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type errorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (e *errorResponse) String() string {
	return fmt.Sprintf("[%+v %+v %+v] ", e.FailedField, e.Tag, e.Value)
}

func ValidateStruct(s interface{}) []*errorResponse {
	var errors []*errorResponse
	v := validator.New()

	err := v.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
		return errors
	}

	return nil
}

func BindJSON(reader io.Reader, obj interface{}) error {
	if err := json.NewDecoder(reader).Decode(obj); err != nil {
		return fmt.Errorf("json decode error: %w", err)
	}

	var fullErrMsg string
	if err := ValidateStruct(obj); err != nil {
		for _, e := range err {
			fullErrMsg += e.String()
		}
		return fmt.Errorf("struct validate error: %w", errors.New(fullErrMsg))
	}

	return nil
}

func SendJSON(w http.ResponseWriter, obj map[string]interface{}) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}

	w.Header().Add("Content-Type", mime.TypeByExtension(".json"))
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("json write error: %w", err)
	}

	return nil
}
