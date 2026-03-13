package entity

import (
	"testing"
	"user_microservice/internal/common/pkg/identity"
)

// Helper functions for tests
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func assertError(t *testing.T, err error, context string) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error for %s, got nil", context)
	}
}

func assertUserNotNil(t *testing.T, user *User) {
	t.Helper()
	if user == nil {
		t.Error("Expected user to be created, got nil")
	}
}

func assertUserNil(t *testing.T, user *User) {
	t.Helper()
	if user != nil {
		t.Error("Expected user to be nil when error occurs")
	}
}

func assertUserFields(t *testing.T, user *User, id, name, email string) {
	t.Helper()
	if user.ID != id {
		t.Errorf("Expected ID %s, got %s", id, user.ID)
	}
	if user.Name.Value() != name {
		t.Errorf("Expected name %s, got %s", name, user.Name.Value())
	}
	if user.Email.Value() != email {
		t.Errorf("Expected email %s, got %s", email, user.Email.Value())
	}
}

// Test data constants
const (
	validName     = "John Doe"
	validEmail    = "john.doe@example.com"
	validPassword = "Password123@"
	invalidEmail  = "invalid-email"
	weakPassword  = "123"
)

func TestNewUser(t *testing.T) {
	t.Run("Should create user successfully with valid data", func(t *testing.T) {
		id := identity.NewUUIDV7()
		user, err := NewUser(id, validName, validEmail, validPassword)

		assertNoError(t, err)
		assertUserNotNil(t, user)
		assertUserFields(t, user, id, validName, validEmail)
	})

	t.Run("Should return error for invalid name", func(t *testing.T) {
		id := identity.NewUUIDV7()
		user, err := NewUser(id, "", validEmail, validPassword)

		assertError(t, err, "empty name")
		assertUserNil(t, user)
	})

	t.Run("Should return error for invalid email", func(t *testing.T) {
		id := identity.NewUUIDV7()
		user, err := NewUser(id, validName, invalidEmail, validPassword)

		assertError(t, err, "invalid email")
		assertUserNil(t, user)
	})

	t.Run("Should return error for invalid password", func(t *testing.T) {
		id := identity.NewUUIDV7()
		user, err := NewUser(id, validName, validEmail, weakPassword)

		assertError(t, err, "weak password")
		assertUserNil(t, user)
	})
}

func TestNewUserWithoutPassword(t *testing.T) {
	t.Run("Should create user successfully without password", func(t *testing.T) {
		id := identity.NewUUIDV7()
		user, err := NewUserWithoutPassword(id, validName, validEmail)

		assertNoError(t, err)
		assertUserNotNil(t, user)
		assertUserFields(t, user, id, validName, validEmail)
	})

	t.Run("Should return error for invalid name", func(t *testing.T) {
		id := identity.NewUUIDV7()
		user, err := NewUserWithoutPassword(id, "", validEmail)

		assertError(t, err, "empty name")
		assertUserNil(t, user)
	})

	t.Run("Should return error for invalid email", func(t *testing.T) {
		id := identity.NewUUIDV7()
		user, err := NewUserWithoutPassword(id, validName, invalidEmail)

		assertError(t, err, "invalid email")
		assertUserNil(t, user)
	})
}

func TestUser_SetName(t *testing.T) {
	user := &User{ID: identity.NewUUIDV7()}

	t.Run("Should set name successfully", func(t *testing.T) {
		newName := "Jane Doe"
		err := user.SetName(newName)

		assertNoError(t, err)
		if user.Name.Value() != newName {
			t.Errorf("Expected name %s, got %s", newName, user.Name.Value())
		}
	})

	t.Run("Should return error for invalid name", func(t *testing.T) {
		err := user.SetName("")
		assertError(t, err, "empty name")
	})
}

func TestUser_SetEmail(t *testing.T) {
	user := &User{ID: identity.NewUUIDV7()}

	t.Run("Should set email successfully", func(t *testing.T) {
		newEmail := "jane.doe@example.com"
		err := user.SetEmail(newEmail)

		assertNoError(t, err)
		if user.Email.Value() != newEmail {
			t.Errorf("Expected email %s, got %s", newEmail, user.Email.Value())
		}
	})

	t.Run("Should return error for invalid email", func(t *testing.T) {
		err := user.SetEmail(invalidEmail)
		assertError(t, err, "invalid email")
	})
}

func TestUser_SetPassword(t *testing.T) {
	user := &User{ID: identity.NewUUIDV7()}

	t.Run("Should set password successfully", func(t *testing.T) {
		newPassword := "NewPassword123@"
		err := user.SetPassword(newPassword)
		assertNoError(t, err)
	})

	t.Run("Should return error for invalid password", func(t *testing.T) {
		err := user.SetPassword(weakPassword)
		assertError(t, err, "weak password")
	})
}

func TestValidateID(t *testing.T) {
	t.Run("Should validate valid UUID successfully", func(t *testing.T) {
		validID := identity.NewUUIDV7()
		err := ValidateID(validID)
		assertNoError(t, err)
	})

	t.Run("Should return error for invalid UUID", func(t *testing.T) {
		invalidID := "invalid-uuid"
		err := ValidateID(invalidID)

		assertError(t, err, "invalid UUID")
		expectedMsg := "Invalid user ID"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})
}

func TestUser_IsEmpty(t *testing.T) {
	t.Run("Should return true for empty user", func(t *testing.T) {
		user := &User{}

		if !user.IsEmpty() {
			t.Error("Expected user to be empty")
		}
	})

	t.Run("Should return false for user with ID", func(t *testing.T) {
		user := &User{
			ID: identity.NewUUIDV7(),
		}

		if user.IsEmpty() {
			t.Error("Expected user not to be empty")
		}
	})
}
