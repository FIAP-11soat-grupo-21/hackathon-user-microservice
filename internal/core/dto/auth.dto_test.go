package dto

import (
	"testing"
)

func TestRegisterUserDTO(t *testing.T) {
	t.Run("Should create RegisterUserDTO with all fields", func(t *testing.T) {
		dto := RegisterUserDTO{
			Email:    "john.doe@example.com",
			Password: "Password123@",
		}

		if dto.Email != "john.doe@example.com" {
			t.Errorf("Expected Email to be 'john.doe@example.com', got '%s'", dto.Email)
		}

		if dto.Password != "Password123@" {
			t.Errorf("Expected Password to be 'Password123@', got '%s'", dto.Password)
		}
	})

	t.Run("Should create RegisterUserDTO with empty fields", func(t *testing.T) {
		dto := RegisterUserDTO{}

		if dto.Email != "" {
			t.Errorf("Expected Email to be empty, got '%s'", dto.Email)
		}

		if dto.Password != "" {
			t.Errorf("Expected Password to be empty, got '%s'", dto.Password)
		}
	})
}
