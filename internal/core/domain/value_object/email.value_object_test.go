package value_object

import "testing"

func TestNewEmail(t *testing.T) {
	tests := []TestCase{
		{"valid email", "john.lennon@testmail.com", false},
		{"invalid email", "john.lennon", true},
		{"empty email", "", true},
		{"email without @", "john.lennon.com", true},
		{"email without domain", "john@", true},
	}

	RunValidationTests(t, "NewEmail()", tests, func(input string) error {
		_, err := NewEmail(input)
		return err
	})
}

func TestEmail_Value(t *testing.T) {
	tests := []ValueTestCase{
		{"Should return email value in lowercase", "JOHN.DOE@EXAMPLE.COM", "john.doe@example.com"},
		{"Should return original lowercase email", "jane.doe@example.com", "jane.doe@example.com"},
	}

	RunValueTests(t, tests, func(input string) (string, error) {
		email, err := NewEmail(input)
		if err != nil {
			return "", err
		}
		return email.Value(), nil
	})
}
