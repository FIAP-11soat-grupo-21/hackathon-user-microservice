package value_object

import "testing"

func TestNewEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid email",
			args: args{
				email: "john.lennon@testmail.com",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			args: args{
				email: "john.lennon",
			},
			wantErr: true,
		},
		{
			name: "empty email",
			args: args{
				email: "",
			},
			wantErr: true,
		},
		{
			name: "email without @",
			args: args{
				email: "john.lennon.com",
			},
			wantErr: true,
		},
		{
			name: "email without domain",
			args: args{
				email: "john@",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmail_Value(t *testing.T) {
	t.Run("Should return email value in lowercase", func(t *testing.T) {
		email, err := NewEmail("JOHN.DOE@EXAMPLE.COM")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expected := "john.doe@example.com"
		if email.Value() != expected {
			t.Errorf("Expected %s, got %s", expected, email.Value())
		}
	})

	t.Run("Should return original lowercase email", func(t *testing.T) {
		email, err := NewEmail("jane.doe@example.com")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expected := "jane.doe@example.com"
		if email.Value() != expected {
			t.Errorf("Expected %s, got %s", expected, email.Value())
		}
	})
}
