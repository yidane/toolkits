package version

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

//Version description GNC Version like Major_Version_Number.Minor_Version_Number[.Revision_Number[.Build_Number]] as follow
//sample1：1.2
//sample2：1.2.0
//sample3：1.2.0.1234
type Version struct {
	Major    int
	Minor    int
	Revision int
	Build    int
}

func New(major, minor, revision, build int) Version {
	return Version{
		Major:    major,
		Minor:    minor,
		Revision: revision,
		Build:    build,
	}
}

//Parse string like 1.2.3 to Version
//if s = 1.2.3.4, only 1.2.3 can be used
func Parse(s string) (v Version, err error) {
	arr := strings.Split(s, ".")
	m := make(map[int]int, 4)

	var n int
	for i := 0; i < len(arr); i++ {
		n, err = strconv.Atoi(arr[i])
		if err != nil {
			return
		}

		m[i] = n
	}

	//if key not exists in map, it will return default value 0
	v.Major = m[0]
	v.Minor = m[1]
	v.Revision = m[2]
	v.Build = m[3]

	return
}

//String return string about Version,like 1.2.3
func (v Version) String() string {
	buf := bytes.Buffer{}
	//at least has Major and Minor
	buf.WriteString(fmt.Sprintf("%d.%d", v.Major, v.Minor))

	s := make([]int, 0)
	if v.Build != 0 {
		s = append(s, v.Build)
	}

	if v.Revision != 0 || len(s) > 0 {
		s = append(s, v.Revision)
	}

	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteString(fmt.Sprintf(".%d", s[i]))
	}

	return buf.String()
}

//CompareTo for compare tow version and return a int
//if return -1, v < other
//if return 0, v == other
//if return 1, v > other
func (v Version) CompareTo(other Version) int {
	//compare from big to small
	arr := [][]int{
		{v.Major, other.Major},
		{v.Minor, other.Minor},
		{v.Revision, other.Revision},
		{v.Build, other.Build},
	}

	num1, num2, result := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		num1 = arr[i][0]
		num2 = arr[i][1]

		switch {
		case num1 > num2:
			result = 1
		case num1 < num2:
			result = -1
		default:
			result = 0
		}

		if result != 0 {
			return result
		}
	}

	return 0
}
