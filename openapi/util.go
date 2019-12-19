package openapi

import (
	"math/rand"
	"time"
)

// generate random string
const (
	keyLen    = 16
	secretLen = 32
	kindAll   = 3
)

// export RandStr
// random strings
func randomStr(size int, kind int) []byte {
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
		if iKind == 3 {
			result[i] = specialChars[rand.Intn(specialCharLen)]
		} else {
			scope, base := kinds[iKind][0], kinds[iKind][1]
			result[i] = uint8(base + rand.Intn(scope))
		}
	}
	return result
}
