package value_object

import "testing"

func TestNewName(t *testing.T) {
	tests := []TestCase{
		{"valid name", "John Lennon", false},
		{"invalid name - empty", "", true},
		{"invalid name - too short", "Jo", true},
		{"invalid name - only spaces", "   ", true},
		{"valid name - exactly 3 characters", "Joe", false},
		{"valid name - with spaces trimmed", "  John Doe  ", false},
	}

	RunValidationTests(t, "NewName()", tests, func(input string) error {
		_, err := NewName(input)
		return err
	})
}

func TestNewName_TooLong(t *testing.T) {
	t.Run("Should return error for name too long", func(t *testing.T) {
		longName := ""
		for i := 0; i < 101; i++ {
			longName += "a"
		}

		_, err := NewName(longName)
		if err == nil {
			t.Error("Expected error for name too long, got nil")
		}

		expectedMsg := "name must have at most 100 characters"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})
}

func TestName_Value(t *testing.T) {
	tests := []ValueTestCase{
		{"Should return trimmed name value", "  John Doe  ", "John Doe"},
		{"Should return exact name when no spaces", "Jane", "Jane"},
	}

	RunValueTests(t, tests, func(input string) (string, error) {
		name, err := NewName(input)
		if err != nil {
			return "", err
		}
		return name.Value(), nil
	})
}
