package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"workflou.com/auth/pkg/loginlink"
	"workflou.com/auth/pkg/validation"
)

func TestEmailIsMissing(t *testing.T) {
	defer Teardown()

	resp, _ := Post("/loginlink", nil)

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 status code, got: %v", resp.StatusCode)
	}

	var body validation.Response
	json.NewDecoder(resp.Body).Decode(&body)

	expected := map[string]string{
		"email": "required",
	}

	if !reflect.DeepEqual(body.Errors, expected) {
		t.Errorf("expected errors: %v, got: %v", expected, body.Errors)
	}
}

func TestEmailFormatIsInvalid(t *testing.T) {
	defer Teardown()

	req, _ := json.Marshal(loginlink.CreateRequest{
		Email: "invalid",
	})

	resp, _ := Post("/loginlink", bytes.NewReader(req))

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 status code, got: %v", resp.StatusCode)
	}

	var body validation.Response
	json.NewDecoder(resp.Body).Decode(&body)

	expected := map[string]string{
		"email": "email",
	}

	if !reflect.DeepEqual(body.Errors, expected) {
		t.Errorf("expected errors: %v, got: %v", expected, body.Errors)
	}
}
