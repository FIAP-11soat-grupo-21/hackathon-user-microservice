package use_case

import (
	"fmt"
	"testing"
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/dto"
)

// Helper function to create a standard test user
func getStandardTestUser() entity.User {
	return makeUser(validUUID, "John Doe", "john.doe@fakemail.com")
}

func TestUpdateUserUseCase(t *testing.T) {
	// Success tests
	t.Run("should update user name successfully", func(t *testing.T) {
		existing := getStandardTestUser()
		repo := &fakeUserRepository{findByIDResult: existing}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		user, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "Jane Doe",
			Email: "john.doe@fakemail.com",
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if user.Name.Value() != "Jane Doe" {
			t.Errorf("Expected name 'Jane Doe', got '%s'", user.Name.Value())
		}
		if !repo.updateCalled {
			t.Error("Expected Update to be called on the repository")
		}
	})

	t.Run("should update user email successfully", func(t *testing.T) {
		existing := getStandardTestUser()
		repo := &fakeUserRepository{
			findByIDResult:      existing,
			existsByEmailResult: false,
		}
		auth := &fakeAuthService{}
		useCase := NewUpdateUserUseCase(repo, auth)

		user, err := useCase.Execute(dto.UpdateUserDTO{
			ID:    validUUID,
			Name:  "John Doe",
			Email: "new.email@fakemail.com",
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if user.Email.Value() != "new.email@fakemail.com" {
			t.Errorf("Expected email 'new.email@fakemail.com', got '%s'", user.Email.Value())
		}
		if !auth.updateEmailCalled {
			t.Error("Expected UpdateUserEmail to be called on the auth service")
		}
	})

	// Error tests using table-driven approach
	errorTests := []struct {
		name          string
		input         dto.UpdateUserDTO
		setupRepo     func() *fakeUserRepository
		setupAuth     func() *fakeAuthService
		expectedError string
	}{
		{
			name:          "should return error if ID is invalid",
			input:         dto.UpdateUserDTO{ID: invalidUUID, Name: "John Doe", Email: "john.doe@fakemail.com"},
			setupRepo:     func() *fakeUserRepository { return &fakeUserRepository{} },
			setupAuth:     func() *fakeAuthService { return &fakeAuthService{} },
			expectedError: "Invalid user ID",
		},
		{
			name:  "should return error if user not found",
			input: dto.UpdateUserDTO{ID: validUUID, Name: "John Doe", Email: "john.doe@fakemail.com"},
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{findByIDResult: entity.User{}}
			},
			setupAuth:     func() *fakeAuthService { return &fakeAuthService{} },
			expectedError: "User not found",
		},
		{
			name:  "should return error if repository FindByID fails",
			input: dto.UpdateUserDTO{ID: validUUID, Name: "John Doe", Email: "john.doe@fakemail.com"},
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{findByIDErr: fmt.Errorf("database connection error")}
			},
			setupAuth:     func() *fakeAuthService { return &fakeAuthService{} },
			expectedError: "database connection error",
		},
		{
			name:  "should return error if new email already exists",
			input: dto.UpdateUserDTO{ID: validUUID, Name: "John Doe", Email: "taken@fakemail.com"},
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{
					findByIDResult:      getStandardTestUser(),
					existsByEmailResult: true,
				}
			},
			setupAuth:     func() *fakeAuthService { return &fakeAuthService{} },
			expectedError: "User with this email already exists",
		},
		{
			name:  "should return error if repository update fails",
			input: dto.UpdateUserDTO{ID: validUUID, Name: "Jane Doe", Email: "john.doe@fakemail.com"},
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{
					findByIDResult: getStandardTestUser(),
					updateErr:      fmt.Errorf("failed to update user"),
				}
			},
			setupAuth:     func() *fakeAuthService { return &fakeAuthService{} },
			expectedError: "failed to update user",
		},
		{
			name:  "should return error if new name is invalid",
			input: dto.UpdateUserDTO{ID: validUUID, Name: "", Email: "john.doe@fakemail.com"},
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{findByIDResult: getStandardTestUser()}
			},
			setupAuth:     func() *fakeAuthService { return &fakeAuthService{} },
			expectedError: "",
		},
		{
			name:  "should return error if new email is invalid",
			input: dto.UpdateUserDTO{ID: validUUID, Name: "John Doe", Email: "invalid-email"},
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{findByIDResult: getStandardTestUser()}
			},
			setupAuth:     func() *fakeAuthService { return &fakeAuthService{} },
			expectedError: "",
		},
		{
			name:  "should return error if ExistsByEmail fails during email update",
			input: dto.UpdateUserDTO{ID: validUUID, Name: "John Doe", Email: "new.email@fakemail.com"},
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{
					findByIDResult:   getStandardTestUser(),
					existsByEmailErr: fmt.Errorf("database connection error"),
				}
			},
			setupAuth:     func() *fakeAuthService { return &fakeAuthService{} },
			expectedError: "database connection error",
		},
		{
			name:  "should return error if auth service UpdateUserEmail fails",
			input: dto.UpdateUserDTO{ID: validUUID, Name: "John Doe", Email: "new.email@fakemail.com"},
			setupRepo: func() *fakeUserRepository {
				return &fakeUserRepository{
					findByIDResult:      getStandardTestUser(),
					existsByEmailResult: false,
				}
			},
			setupAuth: func() *fakeAuthService {
				return &fakeAuthService{updateEmailErr: fmt.Errorf("auth service unavailable")}
			},
			expectedError: "auth service unavailable",
		},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.setupRepo()
			auth := tt.setupAuth()
			useCase := NewUpdateUserUseCase(repo, auth)

			_, err := useCase.Execute(tt.input)

			if err == nil {
				t.Error("Expected error, got nil")
				return
			}
			if tt.expectedError != "" && err.Error() != tt.expectedError {
				t.Errorf("Expected error '%s', got '%s'", tt.expectedError, err.Error())
			}
		})
	}
}
