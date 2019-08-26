package cache

import (
	"sync"
	"time"
)

const (
	defaultCacheSize = 128
)

type cacheValue struct {
	Value  interface{}
	Expire time.Time
}

// get cache value, nil will be returned if if has expired
func (c *cacheValue) Get() interface{} {
	currentNano := time.Now()
	if c.Expire.After(currentNano) {
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
		r :=  v.Get()
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

// expire in milliseconds
func (cm *cacheManager) Put(k string, val interface{}, exp int) {
	cm.mayInit()
	if exp <= 0 {
		return
	}
	cm.Lock.Lock()
	var cacheValue cacheValue
	expire := time.Now().Add(time.Millisecond * time.Duration(exp))
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
