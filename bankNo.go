package toolkits

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

//IsBankNo 是否银行卡账号,主要判断长度是否在16-19之间，验证luhn算法
//Luhn
//	检验数字算法（Luhn Check Digit Algorithm），也叫做模数10公式，是一种简单的算法，用于验证银行卡、信用卡号码的有效性的算法。
//	对所有大型信用卡公司发行的信用卡都起作用，这些公司包括美国Express、护照、万事达卡、Discover和用餐者俱乐部等。
//	这种算法最初是在20世纪60年代由一组数学家制定，现在Luhn检验数字算法属于大众，任何人都可以使用它。
//算法：将每个奇数加倍和使它变为单个的数字，如果必要的话通过减去9和在每个偶数上加上这些值。如果此卡要有效，那么，结果必须是10的倍数。
//http://blog.csdn.net/wangzhjj/article/details/52597614
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

	//luhn校验
	luhn := CreateLuhn(no[:len(no)-1])
	lastNum, _ := strconv.Atoi(no[len(no)-1:])
	if luhn != lastNum {
		return false, errors.New(`Luhn校验失败`)
	}

	return true, nil
}

//CreateLuhn 使用luhn算法生产luhn验证码
func CreateLuhn(no string) int {
	reg, _ := regexp.Compile(`^\d*$`)
	if !reg.MatchString(no) {
		return -1
	}

	return createLuhn(no)
}

func createLuhn(no string) int {
	total := 0
	for i := 1; i <= len(no); i++ {
		num, _ := strconv.Atoi(no[i-1 : i])
		switch i % 2 {
		case 0:
			v := num * 2
			total += v
			if v > 9 {
				total += -9
			}
		case 1:
			total += num
		}
	}

	k := total % 10

	if k == 0 {
		return 0
	}
	return 10 - k
}
