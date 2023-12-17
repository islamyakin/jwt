package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	payload := map[string]string{
		"username": "test_user",
		"password": "test_pass",
		"roles":    "user",
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/create-user", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUserHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, status)
	}

	expected := "User created successfully"
	if body := rr.Body.String(); body != expected {
		t.Errorf("Expected response body %s but got %s", expected, body)
	}
}
