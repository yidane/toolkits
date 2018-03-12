package toolkits

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func defaultOpochTime() *time.Time {
	opochTime := time.Date(2016, 10, 24, 0, 0, 0, 0, time.UTC)
	return &opochTime
}

//CreateOrderNo 生产订单号
func CreateOrderNo(f func() *time.Time) string {
	if f == nil {
		f = defaultOpochTime
	}
	opochTime := f()

	now := time.Now()

	arr := make([]int, 11)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	y := (now.Year() - opochTime.Year()) % 100
	arr[0] = y / 10
	if arr[0] == 0 {
		arr[0] = 9 - rand.Intn(9)
	}
	arr[1] = y % 10

	m := (int(now.Month()) - int(opochTime.Month())) % 12
	arr[2] = m / 10
	if arr[2] == 0 {
		arr[2] = 9 - rand.Intn(8)
	}
	arr[3] = m % 10
	if arr[3] < 0 {
		arr[3] = -arr[3]
	}

	d := now.Day() - opochTime.Day()
	arr[4] = d / 10
	if arr[4] <= 0 {
		arr[4] = 9 - rand.Intn(5)
	}
	arr[5] = d % 10
	if arr[5] < 0 {
		arr[5] = -arr[5]
	}

	h := now.Hour()
	arr[6] = h / 10
	if arr[6] == 0 {
		arr[6] = 2
	}
	arr[7] = h % 10
	if arr[7] < 0 {
		arr[7] = -arr[7]
	}

	min := now.Minute()
	arr[6] += min / 10
	arr[8] = min % 10

	arr[9] = rand.Intn(9)
	arr[10] = rand.Intn(9)

	buf := bytes.Buffer{}
	for i := 0; i < len(arr); i++ {
		buf.WriteString(strconv.Itoa(arr[i]))
	}

	luhn := createLuhn(buf.String())
	buf.WriteString(strconv.Itoa(luhn))

	return buf.String()
}
