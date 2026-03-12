package value_object

import "testing"

func TestNewPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid password",
			args: args{
				password: "P@ssw0rd",
			},
			wantErr: false,
		},
		{
			name: "invalid password - no special chars",
			args: args{
				password: "Password123",
			},
			wantErr: true,
		},
		{
			name: "invalid password - no numbers",
			args: args{
				password: "P@ssword",
			},
			wantErr: true,
		},
		{
			name: "invalid password - no uppercase",
			args: args{
				password: "p@ssw0rd",
			},
			wantErr: true,
		},
		{
			name: "invalid password - no lowercase",
			args: args{
				password: "P@SSW0RD",
			},
			wantErr: true,
		},
		{
			name: "invalid password - too short",
			args: args{
				password: "P@s1",
			},
			wantErr: true,
		},
		{
			name:    "empty password",
			args:    args{password: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewPassword(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("NewPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
