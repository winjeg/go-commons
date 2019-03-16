package str

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

const (
	// numbers only
	KindNumber = 0
	// lower alphabets
	KindLower = 1
	// upper alphabets
	KindUpper = 2
	// numbers and alphabets
	KindAll = 3
	// all kinds with special chars
	KindAllWithSpecial = 4
)

// RandStr random strings
func RandStr(size int, kind int) []byte {
	kinds := [][]int{{10, 48}, {26, 97}, {26, 65}}
	specialChars := []byte{95, 45, 46, 35, 36, 37, 38}
	specialCharLen := len(specialChars)
	iKind, result := kind, make([]byte, size)
	isAll := kind == 3
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {

		// random iKind
		if isAll {
			iKind = rand.Intn(3)
		}
		if kind == KindAllWithSpecial {
			iKind = rand.Intn(4)
		}
		if iKind == 3 {
			result[i] = specialChars[rand.Intn(specialCharLen)]
		} else {
			scope, base := kinds[iKind][0], kinds[iKind][1]
			result[i] = uint8(base + rand.Intn(scope))
		}
	}
	return result
}

// RandomNumAlphabets
// random number and alphabets
func RandomNumAlphabets(length int) string {
	return string(RandStr(length, KindAll))
}

// RandomNumers
// random numbers
func RandomNumers(length int) string {
	return string(RandStr(length, KindNumber))
}

// RandomAlphabetsLower
// random alphabets in lower case
func RandomAlphabetsLower(length int) string {
	return string(RandStr(length, KindLower))
}

// RandomAlphabetsUpper
// random alphabets in upper case
func RandomAlphabetsUpper(length int) string {
	return string(RandStr(length, KindUpper))
}

// RandomStrWithSpecialChars
// random string with special chars including _-.#$%&
func RandomStrWithSpecialChars(length int) string {
	return string(RandStr(length, KindAllWithSpecial))
}

// UUID is to generate unique ids
func UUID() string {
	uid := uuid.NewV4()
	return uid.String()
}

// UUIDShort this method will generate a unique id using uuid, but the result is too long
// so we just use the digits from 0 to 8, thus, increasing the possibility to get a
// duplicated id, but It's okay
// not true uuid, not for tons of ids
func UUIDShort() string {
	u2 := uuid.NewV4()
	d := u2.String()
	return d[24:] + d[9:13]
}

// IsEmpty
// judge if a string is empty
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsBlank
// judge if a string is blank
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// ReplaceAll
// replace all old str to new str for a string
func ReplaceAll(src, old, new string) string {
	return strings.Replace(src, old, new, -1)
}
