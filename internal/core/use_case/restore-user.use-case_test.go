package use_case

import (
	"fmt"
	"testing"
)

func TestRestoreUserUseCase(t *testing.T) {

	t.Run("should restore user successfully", func(t *testing.T) {
		repo := &fakeUserRepository{}
		auth := &fakeAuthService{}
		useCase := NewRestoreUserUseCase(repo, auth)

		err := useCase.Execute(validUUID)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !repo.restoreCalled {
			t.Error("Expected Restore to be called on the repository")
		}

		if !auth.restoreUserCalled {
			t.Error("Expected RestoreUser to be called on the auth service")
		}
	})

	t.Run("should return error if ID is invalid", func(t *testing.T) {
		repo := &fakeUserRepository{}
		auth := &fakeAuthService{}
		useCase := NewRestoreUserUseCase(repo, auth)

		err := useCase.Execute(invalidUUID)

		if err == nil || err.Error() != "Invalid user ID" {
			t.Errorf("Expected 'Invalid user ID', got %v", err)
		}
	})

	t.Run("should return error if repository restore fails", func(t *testing.T) {
		repo := &fakeUserRepository{
			restoreErr: fmt.Errorf("failed to restore user"),
		}
		auth := &fakeAuthService{}
		useCase := NewRestoreUserUseCase(repo, auth)

		err := useCase.Execute(validUUID)

		if err == nil || err.Error() != "failed to restore user" {
			t.Errorf("Expected 'failed to restore user', got %v", err)
		}
	})

	t.Run("should return error if auth service restore fails", func(t *testing.T) {
		repo := &fakeUserRepository{}
		auth := &fakeAuthService{
			restoreUserErr: fmt.Errorf("auth service unavailable"),
		}
		useCase := NewRestoreUserUseCase(repo, auth)

		err := useCase.Execute(validUUID)

		if err == nil || err.Error() != "auth service unavailable" {
			t.Errorf("Expected 'auth service unavailable', got %v", err)
		}
	})
}
