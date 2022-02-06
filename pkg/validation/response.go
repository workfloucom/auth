package validation

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Errors map[string]string `json:"errors"`
}

func Respond(w http.ResponseWriter, err error) {
	errs := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		errs[err.Field()] = err.Tag()
	}

	w.WriteHeader(http.StatusBadRequest)

	r := Response{
		Errors: errs,
	}

	json.NewEncoder(w).Encode(r)
}
