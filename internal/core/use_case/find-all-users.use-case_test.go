package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/domain/entity"
)

// Helper function to create a list of test users
func makeTestUserList(count int) []entity.User {
	users := make([]entity.User, count)
	for i := 0; i < count; i++ {
		users[i] = makeUser(fmt.Sprintf("user-id-%d", i+1), "John Doe", "john.doe@fakemail.com")
	}
	return users
}

func TestFindAllUsersUseCase(t *testing.T) {
	// Error tests
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

	// Success tests
	successTests := []struct {
		name          string
		setupRepo     func() *fakeUserRepository
		expectedCount int
	}{
		{
			name: "should return all users successfully",
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{listAllResult: makeTestUserList(2)}
			},
			expectedCount: 2,
		},
		{
			name: "should return empty list when no users exist",
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{listAllResult: []entity.User{}}
			},
			expectedCount: 0,
		},
	}

	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.setupRepo()
			useCase := NewFindAllUsersUseCase(repo)

			users, err := useCase.Execute()

			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if len(users) != tt.expectedCount {
				t.Errorf("Expected %d users, got %d", tt.expectedCount, len(users))
			}
		})
	}
}
