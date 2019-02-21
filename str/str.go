package str

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

const (
	KindNumber = 0 // numbers only
	KindLower  = 1 // lower alphabets
	KindUpper  = 2 // upper alphabets
	KindAll    = 3 // numbers and alphabets
)

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func RandomNumAlphabets(length int) string {
	return string(Krand(length, KindAll))
}

func RandomNumers(length int) string {
	return string(Krand(length, KindNumber))
}

func RandomAlphabetsLower(length int) string {
	return string(Krand(length, KindLower))
}

func RandomAlphabetsUpper(length int) string {
	return string(Krand(length, KindUpper))
}

// generate uuid
func UUID() string {
	uid := uuid.NewV4()
	return uid.String()
}

// this method will generate a unique id using uuid, but the result is too long
// so we just use the digits from 0 to 8, thus, increasing the possibility to get a
// duplicated id, but It's okay
// not true uuid, not for tons of ids
func UUIDShort() string {
	u2 := uuid.NewV4()
	d := u2.String()
	return d[24:] + d[9:13]
}

// judge if a string is empty
func IsEmpty(str string) bool {
	return len(str) == 0;
}

// judge if a string is blank
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// replace all old str to new str for a string
func ReplaceAll(src, old, new string) string {
	return strings.Replace(src, old, new, -1)
}
