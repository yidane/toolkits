package version

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	major, minor, patch := 2018, 4, 20
	v := New(major, minor, patch)

	if v.Major != major || v.Minor != minor || v.Patch != patch {
		t.Fatal()
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "1", wantErr: false},
		{name: "1.2", wantErr: false},
		{name: "1.2.3", wantErr: false},
		{name: "1.2.3.4", wantErr: false},
		{name: "-1", wantErr: false},
		{name: "a", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type testCase struct {
		Major int
		Minor int
		Patch int
		want  string
	}

	tests := []testCase{
		{Major: 1, Minor: 2, Patch: 3, want: "1.2.3"},
		{Major: 2, Minor: 3, Patch: 4, want: "2.3.4"},
		{Major: 2, Minor: 32223, Patch: 4, want: "2.32223.4"},
		{Major: 2, Minor: 3, Patch: 422221, want: "2.3.422221"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d%d%d", tt.Major, tt.Minor, tt.Patch), func(t *testing.T) {
			v := New(tt.Major, tt.Minor, tt.Patch)
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_CompareTo(t *testing.T) {

	v := New(1, 3, 4)

	type testCase struct {
		Major int
		Minor int
		Patch int
		want  int
	}
	tests := []testCase{
		{1, 3, 3, 1},
		{1, 2, 4, 1},
		{0, 3, 4, 1},
		{1, 3, 4, 0},
		{1, 3, 5, -1},
		{1, 4, 4, -1},
		{2, 3, 4, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d%d%d", tt.Major, tt.Minor, tt.Patch), func(t *testing.T) {
			newV := New(tt.Major, tt.Minor, tt.Patch)
			if got := v.CompareTo(newV); got != tt.want {
				t.Errorf("Version.CompareTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
