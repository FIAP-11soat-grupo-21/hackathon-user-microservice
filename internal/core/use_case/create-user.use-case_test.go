package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/dto"
)

func TestCreateUserCase(t *testing.T) {
	// Testes de erro
	errorTests := []ErrorTestCase{
		{
			Name:          "should return error if email already exists",
			Input:         CreateUserTestInput{Name: "John Doe", Email: "jhon.doe@fakemail.com"},
			ExpectedError: "User with this email already exists",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{existsByEmailResult: true}, &fakeAuthService{}
			},
		},
		{
			Name:          "Should return error if password is too weak",
			Input:         CreateUserTestInput{Name: "John Doe", Email: "jhon.doe@fakemail.com", Password: "123"},
			ExpectedError: "Password must be at least 8 characters long and include uppercase, lowercase, numbers, and special characters",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{existsByEmailResult: false}, &fakeAuthService{}
			},
		},
		{
			Name:          "Should return error if repository insert fails",
			Input:         CreateUserTestInput{Name: "John Doe", Email: "jhon.doe@fakemail.com", Password: "P@ssword123"},
			ExpectedError: "Failed to insert user into repository",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{existsByEmailResult: false, insertErr: fmt.Errorf("Failed to insert user into repository")}, &fakeAuthService{}
			},
		},
		{
			Name:          "Should return error if auth service register fails",
			Input:         CreateUserTestInput{Name: "John Doe", Email: "jhon.doe@fakemail.com", Password: "P@ssword123"},
			ExpectedError: "auth service unavailable",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{existsByEmailResult: false}, &fakeAuthService{registerUserErr: fmt.Errorf("auth service unavailable")}
			},
		},
		{
			Name:          "Should return error if ExistsByEmail repository fails",
			Input:         CreateUserTestInput{Name: "John Doe", Email: "jhon.doe@fakemail.com", Password: "P@ssword123"},
			ExpectedError: "database connection error",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{existsByEmailErr: fmt.Errorf("database connection error")}, &fakeAuthService{}
			},
		},
		{
			Name:  "Should return error if name is invalid",
			Input: CreateUserTestInput{Name: "", Email: "jhon.doe@fakemail.com", Password: "P@ssword123"},
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{existsByEmailResult: false}, &fakeAuthService{}
			},
			// Não especificamos ExpectedError pois qualquer erro de validação de nome é válido
		},
	}

	runErrorTests(t, errorTests, func(repo, auth, input interface{}) error {
		useCase := NewCreateUserUseCase(repo.(*fakeUserRepository), auth.(*fakeAuthService))
		testInput := input.(CreateUserTestInput)
		_, err := useCase.Execute(dto.CreateUserDTO{
			Name:     testInput.Name,
			Email:    testInput.Email,
			Password: testInput.Password,
		})
		return err
	})

	// Teste de sucesso
	t.Run("should create user successfully", func(t *testing.T) {
		repo := &fakeUserRepository{existsByEmailResult: false}
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
}
