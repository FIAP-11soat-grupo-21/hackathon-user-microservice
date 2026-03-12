package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/value_object"
)

const validUUID = "018f4e4a-1c9b-7f3a-8b2d-1a2b3c4d5e6f"
const invalidUUID = "not-a-valid-uuid"

func makeUser(id, name, email string) entity.User {
	n, _ := value_object.NewName(name)
	e, _ := value_object.NewEmail(email)
	return entity.User{ID: id, Name: n, Email: e}
}

func makeUserWithPassword(id, name, email, password string) entity.User {
	n, _ := value_object.NewName(name)
	e, _ := value_object.NewEmail(email)
	p, _ := value_object.NewPassword(password)
	return entity.User{ID: id, Name: n, Email: e, Password: p}
}

func TestFindUserByIDUseCase(t *testing.T) {
	// Testes de erro
	errorTests := []ErrorTestCase{
		{
			Name:          "should return error if ID is invalid",
			Input:         IDTestInput{ID: invalidUUID},
			ExpectedError: "Invalid user ID",
			SetupFunc:     func() (interface{}, interface{}) { return &fakeUserRepository{}, nil },
		},
		{
			Name:          "should return error if user not found",
			Input:         IDTestInput{ID: validUUID},
			ExpectedError: "User not found",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{findByIDResult: entity.User{}}, nil
			},
		},
		{
			Name:          "should return error if repository fails",
			Input:         IDTestInput{ID: validUUID},
			ExpectedError: "database connection error",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{findByIDErr: fmt.Errorf("database connection error")}, nil
			},
		},
	}

	runErrorTests(t, errorTests, func(repo, auth, input interface{}) error {
		useCase := NewFindUserByIDUseCase(repo.(*fakeUserRepository))
		testInput := input.(IDTestInput)
		_, err := useCase.Execute(testInput.ID)
		return err
	})

	// Teste de sucesso
	t.Run("should return user successfully", func(t *testing.T) {
		expected := makeUser(validUUID, "John Doe", "john.doe@fakemail.com")
		repo := &fakeUserRepository{
			findByIDResult: expected,
		}
		useCase := NewFindUserByIDUseCase(repo)

		user, err := useCase.Execute(validUUID)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if user.ID != expected.ID {
			t.Errorf("Expected user ID '%s', got '%s'", expected.ID, user.ID)
		}
	})
}
