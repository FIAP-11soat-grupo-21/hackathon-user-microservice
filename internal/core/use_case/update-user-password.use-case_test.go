package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/dto"
)

func TestUpdateUserPasswordUseCase(t *testing.T) {

	t.Run("should update password successfully", func(t *testing.T) {
		existing := makeUserWithPassword(validUUID, "John Doe", "john.doe@fakemail.com", "OldP@ss1")
		repo := &fakeUserRepository{
			findByIDResult: existing,
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserPasswordUseCase(repo, auth)

		err := useCase.Execute(dto.UpdateUserPasswordDTO{
			ID:          validUUID,
			NewPassword: "NewP@ss123",
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !repo.updateCalled {
			t.Error("Expected Update to be called on the repository")
		}

		if !auth.updatePasswordCalled {
			t.Error("Expected UpdateUserPassword to be called on the auth service")
		}
	})

	t.Run("should return error if ID is invalid", func(t *testing.T) {
		repo := &fakeUserRepository{}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserPasswordUseCase(repo, auth)

		err := useCase.Execute(dto.UpdateUserPasswordDTO{
			ID:          invalidUUID,
			NewPassword: "NewP@ss123",
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
		useCase := NewUpdateUserPasswordUseCase(repo, auth)

		err := useCase.Execute(dto.UpdateUserPasswordDTO{
			ID:          validUUID,
			NewPassword: "NewP@ss123",
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
		useCase := NewUpdateUserPasswordUseCase(repo, auth)

		err := useCase.Execute(dto.UpdateUserPasswordDTO{
			ID:          validUUID,
			NewPassword: "NewP@ss123",
		})

		if err == nil || err.Error() != "database connection error" {
			t.Errorf("Expected 'database connection error', got %v", err)
		}
	})

	t.Run("should return error if new password is too weak", func(t *testing.T) {
		existing := makeUserWithPassword(validUUID, "John Doe", "john.doe@fakemail.com", "OldP@ss1")
		repo := &fakeUserRepository{
			findByIDResult: existing,
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserPasswordUseCase(repo, auth)

		err := useCase.Execute(dto.UpdateUserPasswordDTO{
			ID:          validUUID,
			NewPassword: "123",
		})

		if err == nil {
			t.Error("Expected error for weak password, got nil")
		}
	})

	t.Run("should return error if repository update fails", func(t *testing.T) {
		existing := makeUserWithPassword(validUUID, "John Doe", "john.doe@fakemail.com", "OldP@ss1")
		repo := &fakeUserRepository{
			findByIDResult: existing,
			updateErr:      fmt.Errorf("failed to update user"),
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserPasswordUseCase(repo, auth)

		err := useCase.Execute(dto.UpdateUserPasswordDTO{
			ID:          validUUID,
			NewPassword: "NewP@ss123",
		})

		if err == nil || err.Error() != "failed to update user" {
			t.Errorf("Expected 'failed to update user', got %v", err)
		}
	})

	t.Run("should return error if auth service UpdateUserPassword fails", func(t *testing.T) {
		existing := makeUserWithPassword(validUUID, "John Doe", "john.doe@fakemail.com", "OldP@ss1")
		repo := &fakeUserRepository{
			findByIDResult: existing,
		}
		auth := &fakeAuthService{
			updatePasswordErr: fmt.Errorf("auth service unavailable"),
		}
		useCase := NewUpdateUserPasswordUseCase(repo, auth)

		err := useCase.Execute(dto.UpdateUserPasswordDTO{
			ID:          validUUID,
			NewPassword: "NewP@ss123",
		})

		if err == nil || err.Error() != "auth service unavailable" {
			t.Errorf("Expected 'auth service unavailable', got %v", err)
		}
	})
}
