package cache

import (
	"sync"
	"time"
)

const (
	defaultCacheSize = 128
)

var (
	beginTime = time.Unix(0, 0)
)

type cacheValue struct {
	Value  interface{}
	Expire time.Time
}

// Get cache value, nil will be returned if if has expired
func (c *cacheValue) Get() interface{} {
	if c.Expire.After(time.Now()) || c.Expire.Equal(beginTime) {
		return c.Value
	}
	return nil
}

func (c *cacheValue) Set(value interface{}, exp time.Time) {
	c.Value = value
	c.Expire = exp
}

type cacheManager struct {
	Cache map[string]cacheValue
	Lock  sync.Mutex
}

func (cm *cacheManager) mayInit() {
	if cm.Cache == nil {
		cm.Lock.Lock()
		cm.Cache = make(map[string]cacheValue, defaultCacheSize)
		cm.Lock.Unlock()
	}
}

func (cm *cacheManager) Get(k string) interface{} {
	cm.mayInit()
	if v, ok := cm.Cache[k]; ok {
		r := v.Get()
		if r == nil {
			cm.Lock.Lock()
			delete(cm.Cache, k)
			cm.Lock.Unlock()
			return nil
		}
		return r
	} else {
		return nil
	}
}

// Put expire in milliseconds
// if expire time <= 0, then the key never expires
func (cm *cacheManager) Put(k string, val interface{}, exp int) {
	cm.mayInit()
	cm.Lock.Lock()
	var cacheValue cacheValue
	var expire time.Time
	if exp <= 0 {
		expire = time.Unix(0, 0)
	} else {
		expire = time.Now().Add(time.Millisecond * time.Duration(exp))
	}
	cacheValue.Set(val, expire)
	cm.Cache[k] = cacheValue
	cm.Lock.Unlock()
}

var (
	cache = cacheManager{}
)

func Get(k string) interface{} {
	return cache.Get(k)
}

func Set(k string, val interface{}, exp int) {
	cache.Put(k, val, exp)
}
