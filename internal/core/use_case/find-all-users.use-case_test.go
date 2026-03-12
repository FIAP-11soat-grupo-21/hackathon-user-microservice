package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/value_object"
)

func TestFindAllUsersUseCase(t *testing.T) {

	t.Run("should return all users successfully", func(t *testing.T) {
		name, _ := value_object.NewName("John Doe")
		email, _ := value_object.NewEmail("john.doe@fakemail.com")

		repo := &fakeUserRepository{
			listAllResult: []entity.User{
				{ID: "some-id-1", Name: name, Email: email},
				{ID: "some-id-2", Name: name, Email: email},
			},
		}
		useCase := NewFindAllUsersUseCase(repo)

		users, err := useCase.Execute()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(users) != 2 {
			t.Errorf("Expected 2 users, got %d", len(users))
		}
	})

	t.Run("should return empty list when no users exist", func(t *testing.T) {
		repo := &fakeUserRepository{
			listAllResult: []entity.User{},
		}
		useCase := NewFindAllUsersUseCase(repo)

		users, err := useCase.Execute()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(users) != 0 {
			t.Errorf("Expected 0 users, got %d", len(users))
		}
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		repo := &fakeUserRepository{
			listAllErr: fmt.Errorf("database connection error"),
		}
		useCase := NewFindAllUsersUseCase(repo)

		_, err := useCase.Execute()

		if err == nil || err.Error() != "database connection error" {
			t.Errorf("Expected error 'database connection error', got %v", err)
		}
	})
}
