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

	t.Run("should return error if ID is invalid", func(t *testing.T) {
		repo := &fakeUserRepository{}
		useCase := NewFindUserByIDUseCase(repo)

		_, err := useCase.Execute(invalidUUID)

		if err == nil || err.Error() != "Invalid user ID" {
			t.Errorf("Expected 'Invalid user ID', got %v", err)
		}
	})

	t.Run("should return error if user not found", func(t *testing.T) {
		repo := &fakeUserRepository{
			findByIDResult: entity.User{},
		}
		useCase := NewFindUserByIDUseCase(repo)

		_, err := useCase.Execute(validUUID)

		if err == nil || err.Error() != "User not found" {
			t.Errorf("Expected 'User not found', got %v", err)
		}
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		repo := &fakeUserRepository{
			findByIDErr: fmt.Errorf("database connection error"),
		}
		useCase := NewFindUserByIDUseCase(repo)

		_, err := useCase.Execute(validUUID)

		if err == nil || err.Error() != "database connection error" {
			t.Errorf("Expected 'database connection error', got %v", err)
		}
	})
}
