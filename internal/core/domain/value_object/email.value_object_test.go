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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
