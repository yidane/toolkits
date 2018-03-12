package toolkits

import (
	"testing"
)

func TestIsBankNo(t *testing.T) {
	type args struct {
		no string
	}
	tests := []struct {
		args    string
		want    bool
		wantErr bool
	}{
		//{args: "12", want: false, wantErr: true},
		{args: "6259650871772090", want: true, wantErr: false},
		{args: "12345678901234567a", want: false, wantErr: true},
		{args: "6228480402564890018", want: true, wantErr: false},
		{args: "6228482298797273578", want: true, wantErr: false},
		{args: "6212262201023557228", want: true, wantErr: false},
		{args: "6228481698729890079", want: true, wantErr: false},
		{args: "6227003325370110828", want: true, wantErr: false},
		{args: "6217002710000684874", want: true, wantErr: false},
		{args: "6216616105001489359", want: true, wantErr: false},
		{args: "6216616105001489351", want: false, wantErr: true},
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

func TestCreateLuhn(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		{arg: "622848040256489001", want: 8},
		{arg: "622848229879727357", want: 8},
		{arg: "621226220102355722", want: 8},
		{arg: "622848169872989007", want: 9},
		{arg: "622700332537011082", want: 8},
		{arg: "621700271000068487", want: 4},
		{arg: "621661610500148935", want: 9},
		{arg: "62166161050014893a", want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := CreateLuhn(tt.arg); got != tt.want {
				t.Errorf("CreateLuhn() = %v, want %v", got, tt.want)
			}
		})
	}
}
