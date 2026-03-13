package use_case

import (
	"fmt"
	"testing"
)

func TestDeleteUserUseCase(t *testing.T) {
	// Testes de erro
	errorTests := []ErrorTestCase{
		{
			Name:          "should return error if ID is invalid",
			Input:         IDTestInput{ID: invalidUUID},
			ExpectedError: "Invalid user ID",
			SetupFunc:     func() (interface{}, interface{}) { return &fakeUserRepository{}, &fakeAuthService{} },
		},
		{
			Name:          "should return error if repository delete fails",
			Input:         IDTestInput{ID: validUUID},
			ExpectedError: "failed to delete user",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{deleteErr: fmt.Errorf("failed to delete user")}, &fakeAuthService{}
			},
		},
		{
			Name:          "should return error if auth service delete fails",
			Input:         IDTestInput{ID: validUUID},
			ExpectedError: "auth service unavailable",
			SetupFunc: func() (interface{}, interface{}) {
				return &fakeUserRepository{}, &fakeAuthService{deleteUserErr: fmt.Errorf("auth service unavailable")}
			},
		},
	}

	runErrorTests(t, errorTests, func(repo, auth, input interface{}) error {
		useCase := NewDeleteUserUseCase(repo.(*fakeUserRepository), auth.(*fakeAuthService))
		testInput := input.(IDTestInput)
		return useCase.Execute(testInput.ID)
	})

	// Teste de sucesso
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
}
