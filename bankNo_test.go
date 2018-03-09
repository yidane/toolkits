package toolkits

import "testing"

func TestIsBankNo(t *testing.T) {
	type args struct {
		no string
	}
	tests := []struct {
		args    string
		want    bool
		wantErr bool
	}{
		{args: "12", want: false, wantErr: true},
		{args: "12345678901234567a", want: false, wantErr: true},
		{args: "12345678901234567890", want: false, wantErr: true},
		{args: "12345678901234567", want: false, wantErr: true},
		{args: "62345678901234567", want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			got, err := IsBankNo(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsBankNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsBankNo() = %v, want %v", got, tt.want)
			}
		})
	}
}
