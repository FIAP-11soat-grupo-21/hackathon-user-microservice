package use_case

import (
	"testing"
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/value_object"
)

// Test constants
const (
	validUUID   = "018f4e4a-1c9b-7f3a-8b2d-1a2b3c4d5e6f"
	invalidUUID = "not-a-valid-uuid"
)

// Helper functions to create test users
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

// ErrorTestCase representa um teste que deve resultar em erro
type ErrorTestCase struct {
	Name          string
	Input         interface{}
	ExpectedError string
	SetupFunc     func() (interface{}, interface{}) // Returns (repo, auth)
}

// runErrorTests executa testes que devem resultar em erro
func runErrorTests(t *testing.T, testCases []ErrorTestCase, executor func(interface{}, interface{}, interface{}) error) {
	t.Helper()

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo, auth := tc.SetupFunc()
			err := executor(repo, auth, tc.Input)

			if err == nil {
				t.Errorf("Expected error '%s', got nil", tc.ExpectedError)
				return
			}

			if tc.ExpectedError != "" && err.Error() != tc.ExpectedError {
				t.Errorf("Expected error '%s', got '%s'", tc.ExpectedError, err.Error())
			}
		})
	}
}

// CreateUserTestInput representa os inputs para teste de criação de usuário
type CreateUserTestInput struct {
	Name     string
	Email    string
	Password string
}

// IDTestInput representa entrada apenas com ID
type IDTestInput struct {
	ID string
}
