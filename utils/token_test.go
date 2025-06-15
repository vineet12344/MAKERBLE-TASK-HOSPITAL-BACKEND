package utils

import (
	"os"
	"testing"
)

// Setup: provide a dummy JWT secret for testing
func init() {
	os.Setenv("JWT_SECRET", "testsecret123")
}

// ✅ Test token generation works
func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(1, "test@example.com", "doctor")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Error("Expected a token, got empty string")
	}
}

// ✅ Test token validation returns correct claims
func TestValidateToken_Valid(t *testing.T) {
	token, _ := GenerateToken(42, "alice@hospital.com", "receptionist")
	claims, err := ValidateToken(token)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if claims.Email != "alice@hospital.com" {
		t.Errorf("Expected email to be alice@hospital.com, got %s", claims.Email)
	}
	if claims.Role != "receptionist" {
		t.Errorf("Expected role to be receptionist, got %s", claims.Role)
	}
}

// ❌ Test invalid token returns error
func TestValidateToken_Invalid(t *testing.T) {
	_, err := ValidateToken("this.is.not.a.real.token")
	if err == nil {
		t.Error("Expected error for invalid token, got nil")
	}
}
