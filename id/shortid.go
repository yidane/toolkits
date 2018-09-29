package id

import (
	"bytes"
	"math/rand"
	"time"
)

const (
	bigs              = "ABCDEFGHIJKLMNOPQRSTUVWXY"
	smalls            = "abcdefghjlkmnopqrstuvwxyz"
	numbers           = "0123456789"
	specials          = "-_"
	defaultLength     = 7
	defaultUseNumbers = false
	defaultUseSpecial = true
)

//ShortID for generating short id
type ShortID struct {
	useNumbers bool
	useSpecial bool
	length     int
	rand       *rand.Rand
}

//New a instance of ShortID
func New() *ShortID {
	return &ShortID{
		useNumbers: defaultUseNumbers,
		useSpecial: defaultUseSpecial,
		length:     defaultLength,
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

//UseNumbers whether or not numbers are include in the string
func (shortID *ShortID) UseNumbers(f bool) *ShortID {
	shortID.useNumbers = f
	return shortID
}

//UseSpecial whether or not special characters are included
func (shortID *ShortID) UseSpecial(f bool) *ShortID {
	shortID.useSpecial = f
	return shortID
}

//SetLength set the length of random string which returns
func (shortID *ShortID) SetLength(length uint) *ShortID {
	if length == 0 { //length must be grather than 0
		length = defaultLength
	}

	shortID.length = int(length)
	return shortID
}

//Reset shortID properties and change them to default value
func (shortID *ShortID) Reset() {
	shortID.useNumbers = defaultUseNumbers
	shortID.useSpecial = defaultUseSpecial
	shortID.length = defaultLength
	shortID.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

//Generate a random string
func (shortID *ShortID) Generate() string {
	//default: bigs smalls
	//useNumbers: bigs smalls numbers
	//useSpecials: bigs specials smalls specials numbers

	buf := bytes.Buffer{}
	buf.WriteString(bigs)
	buf.WriteString(smalls)
	if shortID.useSpecial {
		buf.WriteString(specials)
	}
	if shortID.useNumbers {
		buf.WriteString(numbers)
	}

	rtnBuf := bytes.Buffer{}
	l := buf.Len()
	for i := 0; i < shortID.length; i++ {
		r := shortID.rand.Intn(l)
		rtnBuf.WriteByte(buf.Bytes()[r])
	}

	return rtnBuf.String()
}
