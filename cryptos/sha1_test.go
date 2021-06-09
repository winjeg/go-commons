package cryptos

import (
	"github.com/stretchr/testify/assert"
	"github.com/winjeg/go-commons/str"
	"testing"
)

func BenchmarkSha1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sha1(str.ToBytes(str.RandomNumbers(10)))
	}
	b.ReportAllocs()
}

func TestSha1(t *testing.T) {
	data := str.RandomNumbers(10)
	assert.Equal(t, Sha1([]byte(data)), Sha1([]byte(data)))
	assert.NotEqual(t, Sha1([]byte(data)), Sha1([]byte(data+"1")))
	assert.True(t, len(Sha1(nil)) > 0)
}
