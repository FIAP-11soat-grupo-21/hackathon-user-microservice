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
			name: "invalid password",
			args: args{
				password: "password",
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
