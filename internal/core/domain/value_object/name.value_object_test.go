package value_object

import "testing"

func TestNewName(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid name",
			args: args{
				name: "John Lennon",
			},
			wantErr: false,
		},
		{
			name: "invalid name - empty",
			args: args{
				name: "",
			},
			wantErr: true,
		},
		{
			name: "invalid name - too short",
			args: args{
				name: "Jo",
			},
			wantErr: true,
		},
		{
			name: "invalid name - only spaces",
			args: args{
				name: "   ",
			},
			wantErr: true,
		},
		{
			name: "valid name - exactly 3 characters",
			args: args{
				name: "Joe",
			},
			wantErr: false,
		},
		{
			name: "valid name - with spaces trimmed",
			args: args{
				name: "  John Doe  ",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
	t.Run("Should return trimmed name value", func(t *testing.T) {
		name, err := NewName("  John Doe  ")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expected := "John Doe"
		if name.Value() != expected {
			t.Errorf("Expected '%s', got '%s'", expected, name.Value())
		}
	})

	t.Run("Should return exact name when no spaces", func(t *testing.T) {
		name, err := NewName("Jane")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expected := "Jane"
		if name.Value() != expected {
			t.Errorf("Expected '%s', got '%s'", expected, name.Value())
		}
	})
}
