package use_case

import (
	"fmt"
	"testing"
)

func TestRestoreUserUseCase(t *testing.T) {
	// Testes de erro
	errorTests := []ErrorTestCase{
		{
			Name:          "should return error if ID is invalid",
			Input:         IDTestInput{ID: invalidUUID},
			ExpectedError: "Invalid user ID",
			SetupFunc:     func() (interface{}, interface{}) { return &fakeUserRepository{}, &fakeAuthService{} },
		},
		{
			Name:          "should return error if repository restore fails",
			Input:         IDTestInput{ID: validUUID},
			ExpectedError: "failed to restore user",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{restoreErr: fmt.Errorf("failed to restore user")}, &fakeAuthService{}
			},
		},
		{
			Name:          "should return error if auth service restore fails",
			Input:         IDTestInput{ID: validUUID},
			ExpectedError: "auth service unavailable",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{}, &fakeAuthService{restoreUserErr: fmt.Errorf("auth service unavailable")}
			},
		},
	}

	runErrorTests(t, errorTests, func(repo, auth, input interface{}) error {
		useCase := NewRestoreUserUseCase(repo.(*fakeUserRepository), auth.(*fakeAuthService))
		testInput := input.(IDTestInput)
		return useCase.Execute(testInput.ID)
	})

	// Teste de sucesso
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
}
