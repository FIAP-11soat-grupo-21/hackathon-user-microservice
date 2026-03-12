package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/dto"
)

func TestUpdateUserUseCase(t *testing.T) {

	t.Run("should update user name successfully", func(t *testing.T) {
		existing := makeUser(validUUID, "John Doe", "john.doe@fakemail.com")
		repo := &fakeUserRepository{
			findByIDResult: existing,
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		user, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "Jane Doe",
			Email: "john.doe@fakemail.com",
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if user.Name.Value() != "Jane Doe" {
			t.Errorf("Expected name 'Jane Doe', got '%s'", user.Name.Value())
		}

		if !repo.updateCalled {
			t.Error("Expected Update to be called on the repository")
		}
	})

	t.Run("should update user email successfully", func(t *testing.T) {
		existing := makeUser(validUUID, "John Doe", "john.doe@fakemail.com")
		repo := &fakeUserRepository{
			findByIDResult:      existing,
			existsByEmailResult: false,
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		user, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "John Doe",
			Email: "new.email@fakemail.com",
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if user.Email.Value() != "new.email@fakemail.com" {
			t.Errorf("Expected email 'new.email@fakemail.com', got '%s'", user.Email.Value())
		}

		if !auth.updateEmailCalled {
			t.Error("Expected UpdateUserEmail to be called on the auth service")
		}
	})

	t.Run("should return error if ID is invalid", func(t *testing.T) {
		repo := &fakeUserRepository{}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    invalidUUID,
			Name:  "John Doe",
			Email: "john.doe@fakemail.com",
		})

		if err == nil || err.Error() != "Invalid user ID" {
			t.Errorf("Expected 'Invalid user ID', got %v", err)
		}
	})

	t.Run("should return error if user not found", func(t *testing.T) {
		repo := &fakeUserRepository{
			findByIDResult: entity.User{},
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "John Doe",
			Email: "john.doe@fakemail.com",
		})

		if err == nil || err.Error() != "User not found" {
			t.Errorf("Expected 'User not found', got %v", err)
		}
	})

	t.Run("should return error if repository FindByID fails", func(t *testing.T) {
		repo := &fakeUserRepository{
			findByIDErr: fmt.Errorf("database connection error"),
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "John Doe",
			Email: "john.doe@fakemail.com",
		})

		if err == nil || err.Error() != "database connection error" {
			t.Errorf("Expected 'database connection error', got %v", err)
		}
	})

	t.Run("should return error if new email already exists", func(t *testing.T) {
		existing := makeUser(validUUID, "John Doe", "john.doe@fakemail.com")
		repo := &fakeUserRepository{
			findByIDResult:      existing,
			existsByEmailResult: true,
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "John Doe",
			Email: "taken@fakemail.com",
		})

		if err == nil || err.Error() != "User with this email already exists" {
			t.Errorf("Expected 'User with this email already exists', got %v", err)
		}
	})

	t.Run("should return error if repository update fails", func(t *testing.T) {
		existing := makeUser(validUUID, "John Doe", "john.doe@fakemail.com")
		repo := &fakeUserRepository{
			findByIDResult: existing,
			updateErr:      fmt.Errorf("failed to update user"),
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "Jane Doe",
			Email: "john.doe@fakemail.com",
		})

		if err == nil || err.Error() != "failed to update user" {
			t.Errorf("Expected 'failed to update user', got %v", err)
		}
	})

	t.Run("should return error if new name is invalid", func(t *testing.T) {
		existing := makeUser(validUUID, "John Doe", "john.doe@fakemail.com")
		repo := &fakeUserRepository{
			findByIDResult: existing,
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "",
			Email: "john.doe@fakemail.com",
		})

		if err == nil {
			t.Error("Expected error for empty name, got nil")
		}
	})

	t.Run("should return error if auth service UpdateUserEmail fails", func(t *testing.T) {
		existing := makeUser(validUUID, "John Doe", "john.doe@fakemail.com")
		repo := &fakeUserRepository{
			findByIDResult:      existing,
			existsByEmailResult: false,
		}
		auth := &fakeAuthService{
			updateEmailErr: fmt.Errorf("auth service unavailable"),
		}
		useCase := NewUpdateUserUseCase(repo, auth)

		_, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "John Doe",
			Email: "new.email@fakemail.com",
		})

		if err == nil || err.Error() != "auth service unavailable" {
			t.Errorf("Expected 'auth service unavailable', got %v", err)
		}
	})
}
