package cryptos

import (
	"crypto/sha1"
	"fmt"
)

// Sha1  to calc the sha1 value of data
// empty value will be returned when error occurred
func Sha1(data []byte) string {
	shaInst := sha1.New()
	_, err := shaInst.Write(data)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", shaInst.Sum(nil))
}
