package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

const (
	testKey    = "key1"
	testValue  = "val2"
	testValue2 = "Val2"
	testExp    = 5
)

type X string

func TestGet(t *testing.T) {
	strCache := NewCache[X]()
	strCache.Put("a", "1")
	v, ok := strCache.Get("a")
	assert.True(t, ok)
	assert.Equal(t, v, X("1"))

	intCache := NewCustomCache[int, X](1, defaultSleepCount, func(k int) uint32 {
		return uint32(k)
	})

	intCache.PutExp(1, "1", 2000)
	v2, ok := intCache.Get(1)
	assert.True(t, ok)
	assert.Equal(t, v2, X("1"))
	time.Sleep(time.Second * 3)
	fmt.Println("comparing values...")
	assert.Equal(t, 1, int(intCache.expireCount.Load()))

	_, ok = intCache.Get(1)
	assert.True(t, !ok)
	assert.Equal(t, int(intCache.expireCount.Load()), 1)
}

func BenchmarkCacheManager_Get(b *testing.B) {
	strCache := NewCache[X]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := strconv.Itoa(i)
		strCache.Put(k, X(k))
		strCache.Get(k)
	}
	b.ReportAllocs()
}
