package use_case

import (
	"fmt"
	"testing"
)

func TestDeleteUserUseCase(t *testing.T) {

	t.Run("should delete user successfully", func(t *testing.T) {
		repo := &fakeUserRepository{}
		auth := &fakeAuthService{}
		useCase := NewDeleteUserUseCase(repo, auth)

		err := useCase.Execute(validUUID)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !repo.deleteCalled {
			t.Error("Expected Delete to be called on the repository")
		}

		if !auth.deleteUserCalled {
			t.Error("Expected DeleteUser to be called on the auth service")
		}
	})

	t.Run("should return error if ID is invalid", func(t *testing.T) {
		repo := &fakeUserRepository{}
		auth := &fakeAuthService{}
		useCase := NewDeleteUserUseCase(repo, auth)

		err := useCase.Execute(invalidUUID)

		if err == nil || err.Error() != "Invalid user ID" {
			t.Errorf("Expected 'Invalid user ID', got %v", err)
		}
	})

	t.Run("should return error if repository delete fails", func(t *testing.T) {
		repo := &fakeUserRepository{
			deleteErr: fmt.Errorf("failed to delete user"),
		}
		auth := &fakeAuthService{}
		useCase := NewDeleteUserUseCase(repo, auth)

		err := useCase.Execute(validUUID)

		if err == nil || err.Error() != "failed to delete user" {
			t.Errorf("Expected 'failed to delete user', got %v", err)
		}
	})

	t.Run("should return error if auth service delete fails", func(t *testing.T) {
		repo := &fakeUserRepository{}
		auth := &fakeAuthService{
			deleteUserErr: fmt.Errorf("auth service unavailable"),
		}
		useCase := NewDeleteUserUseCase(repo, auth)

		err := useCase.Execute(validUUID)

		if err == nil || err.Error() != "auth service unavailable" {
			t.Errorf("Expected 'auth service unavailable', got %v", err)
		}
	})
}
