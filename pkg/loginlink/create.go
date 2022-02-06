package loginlink

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"workflou.com/auth/pkg/validation"
)

type Create struct {
	Validate validator.Validate
}

type CreateRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (h Create) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body CreateRequest
	json.NewDecoder(r.Body).Decode(&body)

	if err := h.Validate.Struct(body); err != nil {
		validation.Respond(w, err)
		return
	}
}
