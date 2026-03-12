package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/value_object"
)

func TestFindAllUsersUseCase(t *testing.T) {
	// Testes de erro
	errorTests := []ErrorTestCase{
		{
			Name:          "should return error if repository fails",
			Input:         nil,
			ExpectedError: "database connection error",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{listAllErr: fmt.Errorf("database connection error")}, nil
			},
		},
	}

	runErrorTests(t, errorTests, func(repo, auth, input interface{}) error {
		useCase := NewFindAllUsersUseCase(repo.(*fakeUserRepository))
		_, err := useCase.Execute()
		return err
	})

	// Testes de sucesso
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
}
