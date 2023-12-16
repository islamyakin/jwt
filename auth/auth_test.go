package auth

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {

	token, err := CreateToken("kanaya rainbowdrinker")

	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}

	if token == "" {
		t.Error("Token is empty")
	}
}
