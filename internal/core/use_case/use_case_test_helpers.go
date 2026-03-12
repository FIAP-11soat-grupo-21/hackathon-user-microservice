package use_case

import (
	"testing"
)

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
