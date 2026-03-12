package dto

import (
	"testing"
)

func TestCreateUserDTO(t *testing.T) {
	t.Run("Should create CreateUserDTO with all fields", func(t *testing.T) {
		dto := CreateUserDTO{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Password123@",
		}

		expectedFields := map[string]string{
			"Name":     "John Doe",
			"Email":    "john.doe@example.com",
			"Password": "Password123@",
		}
		testDTOWithFields(t, dto, expectedFields)
	})

	t.Run("Should create CreateUserDTO with empty fields", func(t *testing.T) {
		dto := CreateUserDTO{}
		testEmptyDTO(t, dto)
	})
}

func TestUpdateUserDTO(t *testing.T) {
	t.Run("Should create UpdateUserDTO with all fields", func(t *testing.T) {
		dto := UpdateUserDTO{
			ID:    "123e4567-e89b-12d3-a456-426614174000",
			Name:  "Jane Doe",
			Email: "jane.doe@example.com",
		}

		expectedFields := map[string]string{
			"ID":    "123e4567-e89b-12d3-a456-426614174000",
			"Name":  "Jane Doe",
			"Email": "jane.doe@example.com",
		}
		testDTOWithFields(t, dto, expectedFields)
	})

	t.Run("Should create UpdateUserDTO with empty fields", func(t *testing.T) {
		dto := UpdateUserDTO{}
		testEmptyDTO(t, dto)
	})
}

func TestUpdateUserPasswordDTO(t *testing.T) {
	t.Run("Should create UpdateUserPasswordDTO with all fields", func(t *testing.T) {
		dto := UpdateUserPasswordDTO{
			ID:          "123e4567-e89b-12d3-a456-426614174000",
			NewPassword: "NewPassword123@",
		}

		expectedFields := map[string]string{
			"ID":          "123e4567-e89b-12d3-a456-426614174000",
			"NewPassword": "NewPassword123@",
		}
		testDTOWithFields(t, dto, expectedFields)
	})

	t.Run("Should create UpdateUserPasswordDTO with empty fields", func(t *testing.T) {
		dto := UpdateUserPasswordDTO{}
		testEmptyDTO(t, dto)
	})
}
