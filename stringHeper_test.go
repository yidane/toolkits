package toolkits

import (
	"testing"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		{arg: "1", want: 1},
		{arg: "-1", want: -1},
		{arg: "-1.1", want: 0},
		{arg: "a", want: 0},
		{arg: "4294967296", want: 4294967296},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := ParseInt(tt.arg); got != tt.want {
				t.Errorf("Atoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{arg: "1", want: true},
		{arg: "-1", want: true},
		{arg: "-100", want: true},
		{arg: "-1.1", want: true},
		{arg: "-1.1", want: true},
		{arg: "-1.q", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := IsNumber(tt.arg); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{arg: "yidane@163.com", want: true},
		{arg: "yidane@outlook.com", want: true},
		{arg: "yidane_12@163.com", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := IsEmail(tt.arg); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
