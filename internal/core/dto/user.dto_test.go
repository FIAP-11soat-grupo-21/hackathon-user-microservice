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

		if dto.Name != "John Doe" {
			t.Errorf("Expected Name to be 'John Doe', got '%s'", dto.Name)
		}

		if dto.Email != "john.doe@example.com" {
			t.Errorf("Expected Email to be 'john.doe@example.com', got '%s'", dto.Email)
		}

		if dto.Password != "Password123@" {
			t.Errorf("Expected Password to be 'Password123@', got '%s'", dto.Password)
		}
	})

	t.Run("Should create CreateUserDTO with empty fields", func(t *testing.T) {
		dto := CreateUserDTO{}

		if dto.Name != "" {
			t.Errorf("Expected Name to be empty, got '%s'", dto.Name)
		}

		if dto.Email != "" {
			t.Errorf("Expected Email to be empty, got '%s'", dto.Email)
		}

		if dto.Password != "" {
			t.Errorf("Expected Password to be empty, got '%s'", dto.Password)
		}
	})
}

func TestUpdateUserDTO(t *testing.T) {
	t.Run("Should create UpdateUserDTO with all fields", func(t *testing.T) {
		dto := UpdateUserDTO{
			ID:    "123e4567-e89b-12d3-a456-426614174000",
			Name:  "Jane Doe",
			Email: "jane.doe@example.com",
		}

		if dto.ID != "123e4567-e89b-12d3-a456-426614174000" {
			t.Errorf("Expected ID to be '123e4567-e89b-12d3-a456-426614174000', got '%s'", dto.ID)
		}

		if dto.Name != "Jane Doe" {
			t.Errorf("Expected Name to be 'Jane Doe', got '%s'", dto.Name)
		}

		if dto.Email != "jane.doe@example.com" {
			t.Errorf("Expected Email to be 'jane.doe@example.com', got '%s'", dto.Email)
		}
	})

	t.Run("Should create UpdateUserDTO with empty fields", func(t *testing.T) {
		dto := UpdateUserDTO{}

		if dto.ID != "" {
			t.Errorf("Expected ID to be empty, got '%s'", dto.ID)
		}

		if dto.Name != "" {
			t.Errorf("Expected Name to be empty, got '%s'", dto.Name)
		}

		if dto.Email != "" {
			t.Errorf("Expected Email to be empty, got '%s'", dto.Email)
		}
	})
}

func TestUpdateUserPasswordDTO(t *testing.T) {
	t.Run("Should create UpdateUserPasswordDTO with all fields", func(t *testing.T) {
		dto := UpdateUserPasswordDTO{
			ID:          "123e4567-e89b-12d3-a456-426614174000",
			NewPassword: "NewPassword123@",
		}

		if dto.ID != "123e4567-e89b-12d3-a456-426614174000" {
			t.Errorf("Expected ID to be '123e4567-e89b-12d3-a456-426614174000', got '%s'", dto.ID)
		}

		if dto.NewPassword != "NewPassword123@" {
			t.Errorf("Expected NewPassword to be 'NewPassword123@', got '%s'", dto.NewPassword)
		}
	})

	t.Run("Should create UpdateUserPasswordDTO with empty fields", func(t *testing.T) {
		dto := UpdateUserPasswordDTO{}

		if dto.ID != "" {
			t.Errorf("Expected ID to be empty, got '%s'", dto.ID)
		}

		if dto.NewPassword != "" {
			t.Errorf("Expected NewPassword to be empty, got '%s'", dto.NewPassword)
		}
	})
}
