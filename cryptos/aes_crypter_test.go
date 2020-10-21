package cryptos

import (
	"github.com/stretchr/testify/assert"
	"github.com/winjeg/go-commons/str"

	"testing"
)

const (
	key = "321423u9y8d2fwfl12345678"
)

var (
	text = str.RandomStrWithSpecialChars(100)
)

func TestAesDecrypt(t *testing.T) {
	encrypted, err := AesEncrypt(text, key)
	assert.Nil(t, err)
	decrypted, err2 := AesDecrypt(string(encrypted), key)
	assert.Nil(t, err2)
	assert.Equal(t, string(decrypted), text)
}

func BenchmarkAesDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
