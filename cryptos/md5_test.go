package cryptos

import (
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
)

// simple test for md5 sum
func TestMD5(t *testing.T) {
	assert.True(t, strings.EqualFold(MD5([]byte("123abc")), StrMD5("123abc")))
	assert.False(t, strings.EqualFold(MD5([]byte("123abc")), StrMD5("123ab")))
}
