package toolkits

import (
	"errors"
	"regexp"
	"strings"
)

//IsBankNo 是否银行卡账号
func IsBankNo(no string) (bool, error) {
	no = strings.TrimSpace(no)
	l := len(no)
	if l < 16 || l > 19 {
		return false, errors.New(`银行卡号长度必须在16到19之间`)
	}

	reg, _ := regexp.Compile(`^\d*$`)
	if !reg.MatchString(no) {
		return false, errors.New(`银行卡号必须全为数字`)
	}

	if !strings.EqualFold(no[0:2], `62`) {
		return false, errors.New(`Luhn校验失败`)
	}

	return true, nil
}
