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
			name: "invalid name",
			args: args{
				name: "",
			},
			wantErr: true,
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
