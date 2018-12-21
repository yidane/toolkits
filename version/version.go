package version

import (
	"fmt"
	"strconv"
	"strings"
)

//Version description version like Major.Minor.Patch
type Version struct {
	Major int
	Minor int
	Patch int
}

func New(major, minor, patch int) Version {
	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

//Parse string like 1.2.3 to Version
//if s = 1.2.3.4, only 1.2.3 can be used
func Parse(s string) (v Version, err error) {
	arr := strings.Split(s, ".")
	v.Major, err = strconv.Atoi(arr[0])
	if err != nil {
		return
	}

	if len(arr) > 1 {
		v.Minor, err = strconv.Atoi(arr[1])
		if err != nil {
			return
		}
	}

	if len(arr) > 2 {
		v.Patch, err = strconv.Atoi(arr[2])
	}

	return
}

//String return string about Version,like 1.2.3
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

//CompareTo for compare tow version and return a int
//if return -1, v < other
//if return 0, v == other
//if return 1, v > other
func (v Version) CompareTo(other Version) int {
	switch {
	case v.Major > other.Major:
		return 1
	case v.Major < other.Major:
		return -1
	}

	switch {
	case v.Minor > other.Minor:
		return 1
	case v.Minor < other.Minor:
		return -1
	}

	switch {
	case v.Patch > other.Patch:
		return 1
	case v.Patch < other.Patch:
		return -1
	}

	return 0
}
