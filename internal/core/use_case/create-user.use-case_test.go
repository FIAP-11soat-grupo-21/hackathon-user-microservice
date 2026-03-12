package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/dto"
)

func TestCreateUserCase(t *testing.T) {

	t.Run("should return error if email already exists", func(t *testing.T) {
		repo := &fakeUserRepository{
			existsByEmailResult: true,
		}
		auth := &fakeAuthService{}
		useCase := NewCreateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.CreateUserDTO{
			Name:  "John Doe",
			Email: "jhon.doe@fakemail.com",
		})

		if err == nil || err.Error() != "User with this email already exists" {
			t.Errorf("Expected error 'User with this email already exists', got %v", err)
		}
	})

	t.Run("should create user successfully", func(t *testing.T) {
		repo := &fakeUserRepository{
			existsByEmailResult: false,
		}
		auth := &fakeAuthService{}
		useCase := NewCreateUserUseCase(repo, auth)

		user, err := useCase.Execute(dto.CreateUserDTO{
			Name:     "John Doe",
			Email:    "jhon.doe@fakemail.com",
			Password: "P@ssword123",
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if user.Name.Value() != "John Doe" || user.Email.Value() != "jhon.doe@fakemail.com" {
			t.Errorf("Expected user with name 'John Doe' and email 'jhon.doe@fakemail.com', got name '%s' and email '%s'", user.Name, user.Email.Value())
		}

		if !repo.insertCalled {
			t.Error("Expected Insert to be called on the repository")
		}

		if !auth.registerUserCalled {
			t.Error("Expected RegisterUser to be called on the auth service")
		}

	})

	t.Run("Should return error if password is too weak", func(t *testing.T) {
		repo := &fakeUserRepository{
			existsByEmailResult: false,
		}
		auth := &fakeAuthService{}
		useCase := NewCreateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.CreateUserDTO{
			Name:     "John Doe",
			Email:    "jhon.doe@fakemail.com",
			Password: "123",
		})

		if err == nil || err.Error() != "Password must be at least 8 characters long and include uppercase, lowercase, numbers, and special characters" {
			t.Errorf("Expected error 'Password must be at least 8 characters long and contain letters and numbers', got %v", err)
		}
	})

	t.Run("Should return error if repository insert fails", func(t *testing.T) {
		repo := &fakeUserRepository{
			existsByEmailResult: false,
			insertErr:           fmt.Errorf("Failed to insert user into repository"),
		}
		auth := &fakeAuthService{}
		useCase := NewCreateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.CreateUserDTO{
			Name:     "John Doe",
			Email:    "jhon.doe@fakemail.com",
			Password: "P@ssword123",
		})

		if err == nil || err.Error() != "Failed to insert user into repository" {
			t.Errorf("Expected error 'Failed to insert user into repository', got %v", err)
		}
	})

	t.Run("Should return error if auth service register fails", func(t *testing.T) {
		repo := &fakeUserRepository{
			existsByEmailResult: false,
		}
		auth := &fakeAuthService{
			registerUserErr: fmt.Errorf("auth service unavailable"),
		}
		useCase := NewCreateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.CreateUserDTO{
			Name:     "John Doe",
			Email:    "jhon.doe@fakemail.com",
			Password: "P@ssword123",
		})

		if err == nil || err.Error() != "auth service unavailable" {
			t.Errorf("Expected error 'auth service unavailable', got %v", err)
		}
	})

	t.Run("Should return error if ExistsByEmail repository fails", func(t *testing.T) {
		repo := &fakeUserRepository{
			existsByEmailErr: fmt.Errorf("database connection error"),
		}
		auth := &fakeAuthService{}
		useCase := NewCreateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.CreateUserDTO{
			Name:     "John Doe",
			Email:    "jhon.doe@fakemail.com",
			Password: "P@ssword123",
		})

		if err == nil || err.Error() != "database connection error" {
			t.Errorf("Expected error 'database connection error', got %v", err)
		}
	})

	t.Run("Should return error if name is invalid", func(t *testing.T) {
		repo := &fakeUserRepository{
			existsByEmailResult: false,
		}
		auth := &fakeAuthService{}
		useCase := NewCreateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.CreateUserDTO{
			Name:     "",
			Email:    "jhon.doe@fakemail.com",
			Password: "P@ssword123",
		})

		if err == nil {
			t.Error("Expected error for empty name, got nil")
		}
	})
}
