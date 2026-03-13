package value_object

import "testing"

func TestNewPassword(t *testing.T) {
	tests := []TestCase{
		{"valid password", "P@ssw0rd", false},
		{"invalid password - no special chars", "Password123", true},
		{"invalid password - no numbers", "P@ssword", true},
		{"invalid password - no uppercase", "p@ssw0rd", true},
		{"invalid password - no lowercase", "P@SSW0RD", true},
		{"invalid password - too short", "P@s1", true},
		{"empty password", "", true},
	}

	RunValidationTests(t, "NewPassword()", tests, func(input string) error {
		_, err := NewPassword(input)
		return err
	})
}

func TestPassword_Value(t *testing.T) {
	t.Run("Should return encrypted password value", func(t *testing.T) {
		password, err := NewPassword("P@ssw0rd123")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// The value should be encrypted, so it should not be the original password
		if password.Value() == "P@ssw0rd123" {
			t.Error("Expected encrypted password, got original password")
		}

		// Check that the value is not empty
		if password.Value() == "" {
			t.Error("Expected encrypted password value, got empty string")
		}
	})
}
