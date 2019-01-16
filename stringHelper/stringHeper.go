package stringHelper

import (
	"regexp"
	"strconv"
)

//ParseInt returns the result of ParseInt(s, 10, 0) converted to type int.
func ParseInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//ParseInt64 returns the result of ParseInt(s, 10, 0) converted to type int.
func ParseInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 0)
	return i
}

// ParseBool returns the boolean value represented by the string.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func ParseBool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}

//ParseFloat32 converts the string s to a floating-point number
func ParseFloat32(str string) float32 {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}

	return float32(f)
}

//ParseFloat64 converts the string s to a floating-point number
func ParseFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 32)
	return f
}

//IsNumber check the string match number
func IsNumber(str string) bool {
	reg, _ := regexp.Compile(`^(\-)?\d+(\.)?(\d+)?$`)
	return reg.MatchString(str)
}

//IsEmail check the string match email
func IsEmail(str string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	return reg.MatchString(str)
}
