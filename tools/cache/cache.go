package cache

import (
	"github.com/winjeg/go-commons/tools/cmap"
	"log"

	"sync/atomic"
	"time"
)

const (
	defaultCleanInterval = 600  // ten minutes
	sleepInterval        = 10   // 10 micro-seconds
	defaultSleepCount    = 1000 // clean 1000 keys to sleep
)

var (
	beginTime = time.Unix(0, 0)
)

type cacheValue[V any] struct {
	Value  V
	Expire time.Time
}

// Get cache value, nil will be returned if if has expired
func (c *cacheValue[V]) Get() (V, bool) {
	if c.Expire.After(time.Now()) || c.Expire.Equal(beginTime) {
		return c.Value, true
	}
	return *new(V), false
}

func (c *cacheValue[V]) Set(value V, exp time.Time) {
	c.Value = value
	c.Expire = exp
}

type cacheManager[K comparable, V any] struct {
	Cache         cmap.ConcurrentMap[K, cacheValue[V]]
	keyCount      atomic.Uint64
	delCount      atomic.Uint64
	expireCount   atomic.Uint64
	cleanInterval int
	sleepCount    int
}

func NewCustomCache[K comparable, V any](ci, sc int, sharding func(k K) uint32) *cacheManager[K, V] {
	cm := &cacheManager[K, V]{
		Cache:         cmap.NewWithCustomShardingFunction[K, cacheValue[V]](sharding),
		keyCount:      atomic.Uint64{},
		delCount:      atomic.Uint64{},
		expireCount:   atomic.Uint64{},
		sleepCount:    sc,
		cleanInterval: ci,
	}
	go cm.cleanTasks()
	return cm
}

func NewCache[V any]() *cacheManager[string, V] {
	return NewCustomCache[string, V](defaultCleanInterval, defaultSleepCount, cmap.Fnv32)
}

func NewStringCache[K cmap.Stringer, V any]() *cacheManager[K, V] {
	return NewCustomCache[K, V](defaultCleanInterval, defaultSleepCount, cmap.Strfnv32[K])
}

func NewCacheWithConfig[V any](ci, sc int) *cacheManager[string, V] {
	return NewCustomCache[string, V](ci, sc, cmap.Fnv32)
}

func (cm *cacheManager[K, V]) Get(k K) (V, bool) {
	if v, ok := cm.Cache.Get(k); ok {
		if r, ok := v.Get(); ok {
			return r, true
		} else {
			cm.Cache.Remove(k)
			cm.expireCount.Add(1)
			return r, false
		}
	} else {
		return *new(V), false
	}
}

// PutExp expire in milliseconds
// if expire time <= 0, then the key never expires
func (cm *cacheManager[K, V]) PutExp(k K, val V, exp int) {
	var cv cacheValue[V]
	var expire time.Time
	if exp <= 0 {
		expire = time.Unix(0, 0)
	} else {
		expire = time.Now().Add(time.Millisecond * time.Duration(exp))
	}
	cv.Set(val, expire)
	cm.Cache.Set(k, cv)
}

func (cm *cacheManager[K, V]) Put(k K, val V) {
	cm.PutExp(k, val, -1)
}

func (cm *cacheManager[K, V]) Remove(k K) {
	cm.Cache.Remove(k)
	cm.delCount.Add(1)
}

func (cm *cacheManager[K, V]) cleanTasks() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("error running clean up")
		}
	}()

	ticker := time.NewTicker(time.Second * time.Duration(cm.cleanInterval))
	for {
		select {
		case <-ticker.C:
			cm.clean()
		default:
		}
	}
}

func (cm *cacheManager[K, V]) clean() {
	// 清理任务， 用于删除无用的Key，定期执行
	count := uint64(0)
	for item := range cm.Cache.IterBuffered() {
		count++
		k := item.Key
		v := item.Val
		if _, ok := v.Get(); !ok {
			cm.Cache.Remove(k)
			cm.expireCount.Add(1)
			if count%uint64(cm.sleepCount) == 0 {
				time.Sleep(time.Microsecond * sleepInterval)
			}
		}
	}
}
