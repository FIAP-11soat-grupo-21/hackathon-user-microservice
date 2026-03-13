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

		expectedFields := map[string]string{
			"Email":    "john.doe@example.com",
			"Password": "Password123@",
		}
		testDTOWithFields(t, dto, expectedFields)
	})

	t.Run("Should create RegisterUserDTO with empty fields", func(t *testing.T) {
		dto := RegisterUserDTO{}
		testEmptyDTO(t, dto)
	})
}
