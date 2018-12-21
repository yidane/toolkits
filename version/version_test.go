package version

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	major, minor, revision, build := 2018, 4, 20, 12
	v := New(major, minor, revision, build)

	if v.Major != major || v.Minor != minor || v.Revision != revision || v.Build != build {
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
		Major    int
		Minor    int
		Revision int
		Build    int
		want     string
	}

	tests := []testCase{
		{1, 2, 0, 0, "1.2"},
		{1, 2, 0, 4, "1.2.0.4"},
		{1, 2, 3, 0, "1.2.3"},
		{2, 3, 4, 0, "2.3.4"},
		{2, 32223, 4, 0, "2.32223.4"},
		{2, 3, 422221, 2, "2.3.422221.2"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d%d%d%d", tt.Major, tt.Minor, tt.Revision, tt.Build), func(t *testing.T) {
			v := New(tt.Major, tt.Minor, tt.Revision, tt.Build)
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_CompareTo(t *testing.T) {

	v := New(1, 3, 4, 5)

	type testCase struct {
		Major    int
		Minor    int
		Revision int
		Build    int
		want     int
	}
	tests := []testCase{
		{1, 3, 3, 0, 1},
		{1, 2, 4, 0, 1},
		{0, 3, 4, 0, 1},
		{1, 3, 4, 5, 0},
		{1, 3, 5, 0, -1},
		{1, 4, 4, 0, -1},
		{2, 3, 4, 0, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d%d%d%d", tt.Major, tt.Minor, tt.Revision, tt.Build), func(t *testing.T) {
			newV := New(tt.Major, tt.Minor, tt.Revision, tt.Build)
			if got := v.CompareTo(newV); got != tt.want {
				t.Errorf("Version.CompareTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
