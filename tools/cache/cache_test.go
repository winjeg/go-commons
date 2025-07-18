package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
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

	f := func(k uint) bool { return k*k > 100 }
	uintCache := NewIntCache[uint, bool](math.MaxInt32, math.MaxInt32)
	uintCache.PutIfAbsent(13, f)
	uintCache.PutIfAbsent(9, f)
	uv1, _ := uintCache.Get(13)
	uv2, _ := uintCache.Get(9)
	assert.Equal(t, uv1, true)
	assert.Equal(t, uv2, false)
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

func TestConvert(t *testing.T) {
	x := 310293.2131
	y := 87387128378173
	fmt.Println(uint32(x))
	fmt.Println(uint32(y))
}
