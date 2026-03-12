package exception

import (
	"testing"
)

func TestUserNotFoundException(t *testing.T) {
	t.Run("Should return default message when nil exception", func(t *testing.T) {
		var err *UserNotFoundException
		expectedMsg := "User not found"

		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Should return default message when empty message", func(t *testing.T) {
		err := &UserNotFoundException{}
		expectedMsg := "User not found"

		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Should return custom message when provided", func(t *testing.T) {
		customMsg := "User with ID 123 not found"
		err := &UserNotFoundException{Message: customMsg}

		if err.Error() != customMsg {
			t.Errorf("Expected error message '%s', got '%s'", customMsg, err.Error())
		}
	})
}

func TestUserAlreadyExistsException(t *testing.T) {
	t.Run("Should return default message when empty message", func(t *testing.T) {
		err := &UserAlreadyExistsException{}
		expectedMsg := "User already exists"

		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Should return custom message when provided", func(t *testing.T) {
		customMsg := "User with email john@example.com already exists"
		err := &UserAlreadyExistsException{Message: customMsg}

		if err.Error() != customMsg {
			t.Errorf("Expected error message '%s', got '%s'", customMsg, err.Error())
		}
	})
}

func TestInvalidUserDataException(t *testing.T) {
	t.Run("Should return default message when empty message", func(t *testing.T) {
		err := &InvalidUserDataException{}
		expectedMsg := "Invalid user data"

		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Should return custom message when provided", func(t *testing.T) {
		customMsg := "Invalid email format"
		err := &InvalidUserDataException{Message: customMsg}

		if err.Error() != customMsg {
			t.Errorf("Expected error message '%s', got '%s'", customMsg, err.Error())
		}
	})
}
