package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginUser(t *testing.T) {
	payload := map[string]string{
		"username": "test_user",
		"password": "test_pass",
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, status)
	}

	if body := rr.Body.String(); body == "" {
		t.Errorf("Expected non-empty response body but got empty")
	}
}
