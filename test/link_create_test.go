package test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"workflou.com/auth/pkg/link"
	"workflou.com/auth/pkg/validation"
)

func TestRequestValidation(t *testing.T) {
	defer Teardown()

	tt := []struct {
		name     string
		request  *link.CreateRequest
		expected map[string]string
	}{
		{
			name:    "Empty request",
			request: nil,
			expected: map[string]string{
				"email": "required",
			},
		},
		{
			name: "Invalid email",
			request: &link.CreateRequest{
				Email: "invalid",
			},
			expected: map[string]string{
				"email": "email",
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			resp, _ := Post("/link", tc.request)

			if resp.StatusCode != http.StatusBadRequest {
				t.Errorf("expected 400 status code, got: %v", resp.StatusCode)
			}

			var body validation.Response
			json.NewDecoder(resp.Body).Decode(&body)

			if !reflect.DeepEqual(body.Errors, tc.expected) {
				t.Errorf("expected errors: %v, got: %v", tc.expected, body.Errors)
			}
		})
	}
}

func TestUserNotFound(t *testing.T) {
	defer Teardown()

	resp, _ := Post("/link", link.CreateRequest{
		Email: "test@example.com",
	})

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404 status code, got: %v", resp.StatusCode)
	}
}
