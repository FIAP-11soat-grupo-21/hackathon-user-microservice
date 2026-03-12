package entity

import (
	"testing"
	"user_microservice/internal/common/pkg/identity"
)

func TestNewUser(t *testing.T) {
	t.Run("Should create user successfully with valid data", func(t *testing.T) {
		id := identity.NewUUIDV7()
		name := "John Doe"
		email := "john.doe@example.com"
		password := "Password123@"

		user, err := NewUser(id, name, email, password)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if user == nil {
			t.Error("Expected user to be created, got nil")
		}

		if user.ID != id {
			t.Errorf("Expected ID %s, got %s", id, user.ID)
		}

		if user.Name.Value() != name {
			t.Errorf("Expected name %s, got %s", name, user.Name.Value())
		}

		if user.Email.Value() != email {
			t.Errorf("Expected email %s, got %s", email, user.Email.Value())
		}
	})

	t.Run("Should return error for invalid name", func(t *testing.T) {
		id := identity.NewUUIDV7()
		name := ""
		email := "john.doe@example.com"
		password := "Password123@"

		user, err := NewUser(id, name, email, password)

		if err == nil {
			t.Error("Expected error for empty name, got nil")
		}

		if user != nil {
			t.Error("Expected user to be nil when error occurs")
		}
	})

	t.Run("Should return error for invalid email", func(t *testing.T) {
		id := identity.NewUUIDV7()
		name := "John Doe"
		email := "invalid-email"
		password := "Password123@"

		user, err := NewUser(id, name, email, password)

		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}

		if user != nil {
			t.Error("Expected user to be nil when error occurs")
		}
	})

	t.Run("Should return error for invalid password", func(t *testing.T) {
		id := identity.NewUUIDV7()
		name := "John Doe"
		email := "john.doe@example.com"
		password := "123"

		user, err := NewUser(id, name, email, password)

		if err == nil {
			t.Error("Expected error for weak password, got nil")
		}

		if user != nil {
			t.Error("Expected user to be nil when error occurs")
		}
	})
}

func TestNewUserWithoutPassword(t *testing.T) {
	t.Run("Should create user successfully without password", func(t *testing.T) {
		id := identity.NewUUIDV7()
		name := "John Doe"
		email := "john.doe@example.com"

		user, err := NewUserWithoutPassword(id, name, email)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if user == nil {
			t.Error("Expected user to be created, got nil")
		}

		if user.ID != id {
			t.Errorf("Expected ID %s, got %s", id, user.ID)
		}

		if user.Name.Value() != name {
			t.Errorf("Expected name %s, got %s", name, user.Name.Value())
		}

		if user.Email.Value() != email {
			t.Errorf("Expected email %s, got %s", email, user.Email.Value())
		}
	})

	t.Run("Should return error for invalid name", func(t *testing.T) {
		id := identity.NewUUIDV7()
		name := ""
		email := "john.doe@example.com"

		user, err := NewUserWithoutPassword(id, name, email)

		if err == nil {
			t.Error("Expected error for empty name, got nil")
		}

		if user != nil {
			t.Error("Expected user to be nil when error occurs")
		}
	})

	t.Run("Should return error for invalid email", func(t *testing.T) {
		id := identity.NewUUIDV7()
		name := "John Doe"
		email := "invalid-email"

		user, err := NewUserWithoutPassword(id, name, email)

		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}

		if user != nil {
			t.Error("Expected user to be nil when error occurs")
		}
	})
}

func TestUser_SetName(t *testing.T) {
	user := &User{
		ID: identity.NewUUIDV7(),
	}

	t.Run("Should set name successfully", func(t *testing.T) {
		newName := "Jane Doe"
		err := user.SetName(newName)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if user.Name.Value() != newName {
			t.Errorf("Expected name %s, got %s", newName, user.Name.Value())
		}
	})

	t.Run("Should return error for invalid name", func(t *testing.T) {
		err := user.SetName("")

		if err == nil {
			t.Error("Expected error for empty name, got nil")
		}
	})
}

func TestUser_SetEmail(t *testing.T) {
	user := &User{
		ID: identity.NewUUIDV7(),
	}

	t.Run("Should set email successfully", func(t *testing.T) {
		newEmail := "jane.doe@example.com"
		err := user.SetEmail(newEmail)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if user.Email.Value() != newEmail {
			t.Errorf("Expected email %s, got %s", newEmail, user.Email.Value())
		}
	})

	t.Run("Should return error for invalid email", func(t *testing.T) {
		err := user.SetEmail("invalid-email")

		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}
	})
}

func TestUser_SetPassword(t *testing.T) {
	user := &User{
		ID: identity.NewUUIDV7(),
	}

	t.Run("Should set password successfully", func(t *testing.T) {
		newPassword := "NewPassword123@"
		err := user.SetPassword(newPassword)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Should return error for invalid password", func(t *testing.T) {
		err := user.SetPassword("123")

		if err == nil {
			t.Error("Expected error for weak password, got nil")
		}
	})
}

func TestValidateID(t *testing.T) {
	t.Run("Should validate valid UUID successfully", func(t *testing.T) {
		validID := identity.NewUUIDV7()
		err := ValidateID(validID)

		if err != nil {
			t.Errorf("Expected no error for valid UUID, got %v", err)
		}
	})

	t.Run("Should return error for invalid UUID", func(t *testing.T) {
		invalidID := "invalid-uuid"
		err := ValidateID(invalidID)

		if err == nil {
			t.Error("Expected error for invalid UUID, got nil")
		}

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
