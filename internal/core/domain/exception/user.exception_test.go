package exception

import (
	"testing"
)

func TestUserNotFoundException(t *testing.T) {
	t.Run("Should return default message when nil exception", func(t *testing.T) {
		var err *UserNotFoundException
		testExceptionMessage(t, err, "User not found")
	})

	t.Run("Should return default message when empty message", func(t *testing.T) {
		err := &UserNotFoundException{}
		testExceptionMessage(t, err, "User not found")
	})

	t.Run("Should return custom message when provided", func(t *testing.T) {
		customMsg := "User with ID 123 not found"
		err := &UserNotFoundException{Message: customMsg}
		testExceptionMessage(t, err, customMsg)
	})
}

func TestUserAlreadyExistsException(t *testing.T) {
	t.Run("Should return default message when empty message", func(t *testing.T) {
		err := &UserAlreadyExistsException{}
		testExceptionMessage(t, err, "User already exists")
	})

	t.Run("Should return custom message when provided", func(t *testing.T) {
		customMsg := "User with email john@example.com already exists"
		err := &UserAlreadyExistsException{Message: customMsg}
		testExceptionMessage(t, err, customMsg)
	})
}

func TestInvalidUserDataException(t *testing.T) {
	t.Run("Should return default message when empty message", func(t *testing.T) {
		err := &InvalidUserDataException{}
		testExceptionMessage(t, err, "Invalid user data")
	})

	t.Run("Should return custom message when provided", func(t *testing.T) {
		customMsg := "Invalid email format"
		err := &InvalidUserDataException{Message: customMsg}
		testExceptionMessage(t, err, customMsg)
	})
}
