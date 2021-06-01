package cryptos

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"testing"
)

const (
	bkey  = "321423u9y8d2fwfl12345678"
	btext = "hello, world!"
)

func TestBase64Decrypt(t *testing.T) {
	encrypted, err := Base64Encrypt(btext, bkey)
	assert.Nil(t, err)
	fmt.Print(encrypted)
	dec, decErr := Base64Decrypt(encrypted, bkey)
	assert.Nil(t, decErr)
	assert.Equal(t, dec, btext)
}
