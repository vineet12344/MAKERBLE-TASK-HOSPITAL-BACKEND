package utils

import "testing"

// !  Test if password is hashed or not

func TestHashPassword(t *testing.T) {
	password := "admin123"
	hashed, err := HashPassword(password)

	if err != nil {
		t.Errorf("Expected no error but got %v \n", err)
	}

	if hashed == password {
		t.Errorf("Hashed password should not be the same as original")
	}
}

// !  Test that correct password passes validation

func TestCheckPassword_Valid(t *testing.T) {
	password := "doctor1234"
	hashed, _ := HashPassword(password)

	if !CheckPassword(password, hashed) {
		t.Errorf("Expected password check to pass, but it failed")
	}
}

// !  Test that wrong password fails validation

func TestCheckPassword_Invalid(t *testing.T) {
	password := "mypassword"
	hashed, _ := HashPassword(password)

	if CheckPassword(hashed, "wrongpassword") {
		t.Errorf("Expected password check to fail, but it passed")
	}
}
