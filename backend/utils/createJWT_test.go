package utils

import (
	"strings"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	username := "testuser"
	token, err := CreateJWT(username)
	if err != nil {
		t.Errorf("Error creating JWT: %v", err)
	}
	// Verify that the token is not empty
	if strings.Trim(token, "Bearer ") == "" {
		t.Errorf("JWT token is empty")
	}
}
