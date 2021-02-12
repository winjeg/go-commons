package cryptos

/// md5 package is good and convenient enough
// but we use the md5 sum as strings more often

import (
	"crypto/md5"
	"fmt"
)

func MD5(data []byte) string {
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func StrMD5(data string) string {
	return MD5([]byte(data))
}
