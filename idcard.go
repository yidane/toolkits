package toolkits

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//IsIDCard 是否身份证号码
func IsIDCard(no string) (bool, error) {
	no = strings.TrimSpace(no)
	l := len(no)
	if l != 15 && l != 18 {
		return false, errors.New("身份证号长度不对，或者号码不符合规定！15位号码应全为数字，18位号码末位可以为数字或X")
	}

	no = strings.ToUpper(no)
	rep, _ := regexp.Compile(`(^\d{15}$)|(^\d{17}([0-9]|X)$)`)
	if !rep.MatchString(no) {
		return false, errors.New("身份证号长度不对，或者号码不符合规定！15位号码应全为数字，18位号码末位可以为数字或X")
	}

	citys := map[int]string{11: `北京`, 12: `天津`, 13: `河北`, 14: `山西`, 15: `内蒙古`, 21: `辽宁`, 22: `吉林`, 23: `黑龙江`, 31: `上海`, 32: `江苏`, 33: `浙江`, 34: `安徽`, 35: `福建`, 36: `江西`, 37: `山东`, 41: `河南`, 42: `湖北`, 43: `湖南`, 44: `广东`, 45: `广西`, 46: `海南`, 50: `重庆`, 51: `四川`, 52: `贵州`, 53: `云南`, 54: `西藏 `, 61: `陕西`, 62: `甘肃`, 63: `青海`, 64: `宁夏`, 65: `新疆`, 71: `台湾`, 81: `香港`, 82: `澳门`, 91: `国外 `}

	cChar := no[0:2]
	city, _ := strconv.Atoi(cChar)

	if _, ok := citys[city]; !ok {
		return false, errors.New("非法地区")
	}

	year, month, day := 0, 0, 0
	switch l {
	case 15:
		year, _ = strconv.Atoi(no[6:8])
		year += 1900 //15位身份证号码年份是简写的
		month, _ = strconv.Atoi(no[8:10])
		day, _ = strconv.Atoi(no[10:12])
	case 18:
		year, _ = strconv.Atoi(no[6:10])
		month, _ = strconv.Atoi(no[10:12])
		day, _ = strconv.Atoi(no[12:14])
	}

	birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if birthday.Year() != year || birthday.Month() != time.Month(month) || birthday.Day() != day {
		return false, errors.New(`身份证号里出生日期不对`)
	}

	if l == 15 {
		return true, nil
	}

	//18位身份证末尾验证码校验算法
	arr1 := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	arr2 := []string{`1`, `0`, `X`, `9`, `8`, `7`, `6`, `5`, `4`, `3`, `2`}
	num := 0
	for i := 0; i < 17; i++ {
		ni, _ := strconv.Atoi(no[i : i+1])
		num += ni * arr1[i]
	}
	code := arr2[num%11]
	if code != no[17:18] {
		return false, errors.New("身份证号里末位输入错误")
	}

	return true, nil
}
